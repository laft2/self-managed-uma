package resource_server

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func AddPatHandler(e *echo.Echo, conf *oauth2.Config) {
	g := e.Group("/pat")
	patDB := map[string]*oauth2.Token{}
	g.GET("/callback", func(c echo.Context) error {
		pat, err := conf.Exchange(context.TODO(), c.QueryParam("code"))
		if err != nil {
			return err
		}
		sess, err := store.Get(c.Request(), "suma_rs")
		if err != nil {
			return err
		}
		id, ok := sess.Values["id"].(string)
		if !ok {
			return fmt.Errorf("fail to cast `id` to string")
		}
		patDB[id] = pat
		return nil
	})
}

func GetPat(conf *oauth2.Config) (*oauth2.Token, error) {
	return conf.Exchange(context.TODO(), "")
}

var conf *oauth2.Config = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	Scopes:       []string{"SCOPE1", "SCOPE2"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/o/oauth2/auth",
		TokenURL: "https://provider.com/o/oauth2/token",
	},
}

func Oauth_example() {
	ctx := context.Background()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	client.Get("...")
}
