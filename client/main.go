package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type PermissionTicket struct {
	ID string `json:"ticket"`
}
type PermissionTicketResponse struct {
	PermTicket PermissionTicket `json:"permission_ticket"`
	QsUri      string           `json:"qs_endpoint"`
}

func parsePermTicket() (*PermissionTicketResponse, error) {
	f, err := os.Open("./permticket.txt")
	if err != nil {
		return nil, err
	}
	permTicketResp := &PermissionTicketResponse{}
	if err := json.NewDecoder(f).Decode(permTicketResp); err != nil {
		return nil, err
	}
	return permTicketResp, nil
}

type TokenRequestForm struct {
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
}

func requestAccessToken(permTicketResp *PermissionTicketResponse) (*AccessTokenInfo, error) {
	tokenReq := &TokenRequestForm{
		GrantType: "urn:ietf:params:oauth:grant-type:uma-ticket", // TODO: may change urn
		Ticket:    permTicketResp.PermTicket.ID,
	}
	urlValue := url.Values{}
	urlValue.Set("grant_type", tokenReq.GrantType)
	urlValue.Set("ticket", tokenReq.Ticket)
	req, err := http.NewRequest("POST", permTicketResp.QsUri+"/suma/token", strings.NewReader(urlValue.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic iamclient")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	if statusCode >= 400 {
		return nil, fmt.Errorf("qs returns %v", statusCode)
	}

	accessTokenResp := &AccessTokenInfo{}
	err = json.NewDecoder(resp.Body).Decode(accessTokenResp)
	if err != nil {
		panic(err)
		return nil, err
	}
	return accessTokenResp, nil
}

func main() {
	// e := echo.New()
	// e.Use(middleware.Logger())

	// e.Logger.Fatal(e.Start("localhost:10102"))
	permTicket, err := parsePermTicket()
	if err != nil {
		panic(err)
	}
	fmt.Printf("permTicket: %+v\n", permTicket)
	accessToken, err := requestAccessToken(permTicket)
	if err != nil {
		panic(err)
	}
	fmt.Printf("accessToken: %+v\n", accessToken)
}
