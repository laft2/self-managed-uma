package queuing_server

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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

var pool = []interface{}{} // not thread safe

var testData = []interface{}{
	map[string]interface{}{
		"ticket": "test_ticket",
		"requested_scopes": map[string]interface{}{
			"resources": []map[string]interface{}{
				{
					"resource_id": "test_resource_id",
					"resource_scopes": []string{
						"view",
					},
				},
			},
			"client_key": "test_client_key",
		},
		"client_request": map[string]interface{}{
			"grant_type":  "test_grant_type",
			"ticket":      "test_ticket",
			"client_info": "sample_client",
		},
	},
}

func AddQueueGroup(e *echo.Echo) {
	queueGroup := e.Group("/queue")
	// communicate with smartphone (authorization server)
	queueGroup.GET("/test/requests", func(c echo.Context) error {
		// test plain connection
		return c.JSON(http.StatusOK, testData)
	})
	queueGroup.POST("/test/request", func(c echo.Context) error {
		pool = append(pool, testData...)
		return nil
	})
	queueGroup.GET("/requests", func(c echo.Context) error {
		resp := pool
		if len(resp) == 0 {
			return c.JSON(http.StatusNoContent, nil)
		}
		pool = []interface{}{}
		return c.JSON(http.StatusOK, resp)
	})

	queueGroup.POST("/rpt", func(c echo.Context) error {
		ticket := c.FormValue("ticket")
		fmt.Printf("ticket: %v\n", ticket)
		rpt := c.FormValue("rpt")
		fmt.Printf("rpt: %v\n", rpt)
		clientEp := c.FormValue("clientEp")
		body := bytes.NewBufferString("{rpt:" + rpt + "}")
		http.DefaultClient.Post(
			clientEp,
			"application/json",
			body,
		)
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

}

// send authorization request to authorization server

// func makeAuthorizationRequest() *AuthorizationRequest {
// 	x := &AuthorizationRequest{
// 		Timestamp: 1,
// 	}
// 	return nil
// }
