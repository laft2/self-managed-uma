package queuing_server

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TokenForm struct {
	GrantType        string `form:"grant_type"`
	Ticket           string `form:"ticket"`
	ClaimToken       string `form:"claim_token"`
	ClaimTokenFormat string `form:"claim_token_format"`
	Pct              string `form:"pct"`
	Rpt              string `form:"rpt"`
	Scope            string `form:"scope"`
}

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       []string
}
type AccessTokenMap map[string]AccessTokenInfo

var permissionTicketStore = PermissionTicketMap{}
var accessTokenStore = AccessTokenMap{}

func createAccessToken(resourceId string, scope []string) (*AccessTokenInfo, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	accessToken := uuid.String()
	accessTokenInfo := &AccessTokenInfo{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		Scope:       scope,
	}
	accessTokenStore[accessToken] = *accessTokenInfo
	return accessTokenInfo, nil
}

func AddUmaGroup(e *echo.Echo) {
	umagroup := e.Group("/suma")
	umagroup.POST("/token", func(c echo.Context) error {
		tokenForm := &TokenForm{}
		c.Bind(tokenForm)
		fmt.Printf("tokenForm: %+v\n", tokenForm)
		// reject if the request does not have required field
		if tokenForm.GrantType != "urn:ietf:params:oauth:grant-type:uma-ticket" {
			return c.String(http.StatusBadRequest, "bad grant type")
		}
		// reject if the request has only one of claim_token and claim_token_format
		if (tokenForm.ClaimToken == "" && tokenForm.ClaimTokenFormat != "") || (tokenForm.ClaimToken != "" && tokenForm.ClaimTokenFormat == "") {
			return c.String(http.StatusBadRequest, "only one of claim_token and claim_token_format")
		}

		authHeader := c.Request().Header.Get("Authorization")
		fmt.Printf("authHeader: %v\n", authHeader) // TODO: implement authentication of client

		ticketInfo, err := GetTicketInfo(tokenForm.Ticket)
		if err != nil {
			return err
		}

		authzReq := AuthorizationRequest{}
		authzReq.RSDomain := ticketInfo.ID

		// accessTokenInfo, err := CreateRpt(ticketInfo.PermRequest.ResourceID, ticketInfo.PermRequest.ResourceScopes)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("accessTokenInfo: %+v\n", accessTokenInfo)
		return c.JSON(RequestSubmitted.StatusCode, RequestSubmitted)
	})
	umagroup.POST("/perm", PostPerm)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// TODO: manage resource
	e.GET("/user", func(c echo.Context) error {
		return nil
	})
	AddUmaGroup(e)

	deviceGroup := e.Group("/device")
	// get authorization request from smart phone
	deviceGroup.GET("/req", func(c echo.Context) error {

		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start("localhost:9010"))
}
