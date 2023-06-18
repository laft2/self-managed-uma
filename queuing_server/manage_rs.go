package queuing_server

import (
	"html/template"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/labstack/echo/v4"
)

type Resource struct {
	ID               string
	AcceptableScopes []string
}
type RS struct {
	RSDomain    string
	Resources   []Resource
	ID          string
	Secret      string
	RedirectUri []string
}
type PAT struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	RS          *RS
}

func renderHTML(htmlpath string, w http.ResponseWriter, data interface{}) error {
	t, err := template.ParseFiles(htmlpath)
	if err != nil {
		return err
	}
	if err := t.Execute(w, data); err != nil {
		return err
	}
	return nil
}

func AddPatGroup(e *echo.Echo) {
	store.NewClientStore()
	g := e.Group("/pat")
	g.GET("/auth", func(c echo.Context) error {
		qps := c.QueryParams()
		// response_type must be code
		if qps.Get("response_type") != "code" {
			return c.JSON(http.StatusBadRequest, ErrorJSON{
				Error:            "invalid_request",
				ErrorDescription: "invalid response_type",
			})
		}
		// client_id := qps.Get("client_id")
		qps.Get("redirect_uri")
		qps.Get("scope")
		qps.Get("state")

		return renderHTML("front/authz.html", c.Response(), struct{}{})
	})
	g.POST("/approve", func(c echo.Context) error {
		return nil
	})
	g.POST("/token", func(c echo.Context) error {
		return nil
	})
}
