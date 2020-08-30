package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"hacku_vol2/http/response"
)

type feedInformationRequest struct {
	Token   string `json:"name"`
	Flag    int    `json:flag`
	feednum string `json : "feednum"`
}

type feedInformationResponse struct {
	Token   string `json:token`
	Feednum string `json:feednum`
}

func HandleFeedChange() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody feedInformationRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		//response.Success(writer, &userInformationResponse{
		//	Token: hashingPass,
		//})

	}
}
