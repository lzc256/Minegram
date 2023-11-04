package utils

import "maunium.net/go/mautrix/id"

var MatrixInstance *Client

func InitMatrix(homeserverURL string, userID id.UserID, accessToken string) {
	MatrixInstance = NewClient(homeserverURL, userID, accessToken)

}
