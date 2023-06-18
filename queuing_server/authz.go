package queuing_server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthorizationResponse struct {
	RespJwt string
}

type Rpt struct {
	AccessToken string `json:"access_token"` // jwt
	TokenType   string `json:"token_type"`
	Pct         string `json:"pct,omitempty"`
}

type AuthzError struct {
	StatusCode     int    // no encode
	Error          string `json:"error"`
	Ticket         string `json:"ticket,omitempty"`
	RequiredClaims []struct {
		ClaimTokenFormat []string `json:"claim_token_format,omitempty"`
		ClaimType        string   `json:"claim_type,omitempty"`
		FriendlyName     string   `json:"friendly_name,omitempty"`
		Issuer           []string `json:"issuer,omitempty"`
		Name             string   `json:"name,omitempty"`
	} `json:"required_claims,omitempty"`
	RedirectUser string `json:"redirect_user,omitempty"`
}

var InvalidGrant AuthzError = AuthzError{
	StatusCode: http.StatusBadRequest,
	Error:      "invalid_grant",
}
var InvalidScope AuthzError = AuthzError{
	StatusCode: http.StatusBadRequest,
	Error:      "invalid_scope",
}
var NeedInfo AuthzError = AuthzError{
	StatusCode: http.StatusForbidden,
	Error:      "need_info",
}
var RequestDenied AuthzError = AuthzError{
	StatusCode: http.StatusForbidden,
	Error:      "request_denied",
}
var RequestSubmitted AuthzError = AuthzError{
	StatusCode: http.StatusForbidden,
	Error:      "request_submitted",
}

// communicate with smartphone (authorization server)
func AddAuthzGroup(e *echo.Echo) {
	authzGroup := e.Group("/authz")
	authzGroup.GET("/requests", func(c echo.Context) error {
		requests, err := GetWaitingRequests()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, requests)
	})
	authzGroup.POST("/rpt", func(c echo.Context) error {
		return nil
	})
}

// send authorization request to authorization server

// func makeAuthorizationRequest() *AuthorizationRequest {
// 	x := &AuthorizationRequest{
// 		Timestamp: 1,
// 	}
// 	return nil
// }
