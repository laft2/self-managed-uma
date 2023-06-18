package queuing_server

import "github.com/google/uuid"

type RPTInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       []string
}
type RPTMap map[string]RPTInfo

var rptStore = RPTMap{}

func CreateRpt(resourceId string, scope []string) (*RPTInfo, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	rpt := uuid.String()
	rptInfo := &RPTInfo{
		AccessToken: rpt,
		TokenType:   "Bearer",
		Scope:       scope,
	}
	rptStore[rpt] = *rptInfo
	return rptInfo, nil
}
