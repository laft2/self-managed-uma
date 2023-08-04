package queuing_server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/laft2/self-managed-uma/queuing_server/qs_db"
)

type AuthorizationResponse struct {
	RespJwt string
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

var InvalidRequest AuthzError = AuthzError{
	StatusCode: http.StatusUnauthorized,
	Error:      "invalid_request",
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

func AddQueueGroup(e *echo.Echo) {
	queueGroup := e.Group("/queue")
	queueGroup.GET("/requests/:user_id", func(c echo.Context) error {
		// communicate with smartphone (authorization server)
		// TODO: authenticate user
		// curl example: curl -XGET 'http://localhost:9010/queue/requests/1'
		user_id, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return err
		}
		requests, err := qs_db.SelectWaitingRequests(user_id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, requests)
	})
	queueGroup.POST("/requests", func(c echo.Context) error {
		ticket := c.FormValue("ticket")
		clientReq := c.FormValue("request")
		err := qs_db.AddClientRequest(ticket, clientReq)
		if err != nil {
			return c.JSON(InvalidRequest.StatusCode, InvalidRequest)
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"status": "waiting",
		})
	})
	queueGroup.POST("/rpt", func(c echo.Context) error {
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
