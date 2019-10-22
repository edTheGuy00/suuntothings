package actions

import (
	"fmt"
	"log"

	"github.com/edTheGuy00/suuntothings/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// auth page:  https://cloudapi-oauth.suunto.com/oauth/authorize?response_type=code&client_id=<your client id from profile/oauth settings>

// AuthCallback default implementation.
func AuthCallback(c buffalo.Context) error {
	param := c.Param("code")

	log.Println(param)

	callbackURL := fmt.Sprintf("%s/%s", envy.Get("ENDPOINT", "http://127.0.0.1:3000"), "auth-callback")

	conf := &oauth2.Config{
		ClientID:     envy.Get("CLIENT_ID", ""),
		ClientSecret: envy.Get("CLIENT_SECRET", ""),
		Scopes:       []string{"activity", "location"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  callbackURL,
			TokenURL: "https://cloudapi-oauth.suunto.com/oauth/token",
		},
	}

	tok, err := conf.Exchange(c, param)

	if err != nil {
		log.Fatalln(err)
	}

	u := &models.User{}

	u.AccessToken = tok.AccessToken
	u.RefreshToken = tok.RefreshToken
	u.UserName = getUserName(tok.AccessToken)

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	verrs, err := tx.ValidateAndCreate(u)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		return errors.WithStack(errors.New("sumting wen rong"))
	}

	return c.Render(200, r.JSON(map[string]string{"message": "success"}))
}

func getUserName(tokenStr string) string {
	var claims map[string]interface{}
	token, _ := jwt.ParseSigned(tokenStr)
	_ = token.UnsafeClaimsWithoutVerification(&claims)

	return claims["user_name"].(string)
}
