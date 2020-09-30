package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Models"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// googleOauthConfig ...
var googleOauthConfig *oauth2.Config

// InitGoogleAuth ...
func InitGoogleAuth() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("CALLBACK_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

// oauthStateString
var oauthStateString string

// HandleGoogleLogin ...
func HandleGoogleLogin(c *gin.Context) {
	store := ginsession.FromContext(c)
	state := generateState()
	store.Set("state", state)
	err := store.Save()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	url := googleOauthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url+"&prompt=consent&access_type=offline")
}

// HandleGoogleCallback ...
func HandleGoogleCallback(c *gin.Context) {
	store := ginsession.FromContext(c)
	state, ok := store.Get("state")
	if !ok {
		c.AbortWithStatus(404)
		return
	}

	content, token, err := getUserInfo(c.DefaultQuery("state", " "), c.DefaultQuery("code", " "), state.(string))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data map[string]interface{}
	if err = json.Unmarshal(content, &data); err != nil {
		fmt.Println(err.Error())
		return
	}

	var user Models.User
	userToLookFor := Models.User{Email: data["email"].(string), FirstName: data["given_name"].(string), LastName: data["family_name"].(string)}
	if err = Models.GetOrCreateUser(&user, userToLookFor); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	newToken := Models.Token{Token: token, UserID: user.ID}

	err = Models.CreateOrUpdateToken(&newToken)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"data":  data,
		"token": token,
	})
}

// savedToken ...
var savedToken *oauth2.Token

// getUserInfo ...
func getUserInfo(state string, code string, stateFromSession string) ([]byte, *oauth2.Token, error) {
	if state != stateFromSession {
		return nil, nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	/*fmt.Println(token.AccessToken)
	fmt.Println(token)
	savedToken = token*/
	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, token, nil
}

// generateState ...
func generateState() string {
	length := 32
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// GetNewToken ...
func GetNewToken(token *oauth2.Token) (*oauth2.Token, error) {
	source := googleOauthConfig.TokenSource(oauth2.NoContext, token)
	newToken, err := source.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

/*
// HandleGetToken ...
func HandleGetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Authorization"))
	if r.Header.Get("Authorization") == fmt.Sprintf("%s %s", savedToken.TokenType, savedToken.AccessToken) {
		source := googleOauthConfig.TokenSource(oauth2.NoContext, savedToken)
		newToken, err := source.Token()
		if err != nil {
			panic(err.Error)
		}
		fmt.Fprintf(w, "New Token: %s\n", newToken)
	}
}
*/
