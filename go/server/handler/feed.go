package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"hacku_vol2/http/response"
	"hacku_vol2/server/model"
)

type feedInformationRequest struct {
	Token   string `json:"token"`
	Flag    int    `json:"flag"`
	FeedNum int    `json : "feednum"`
}

type feedInformationResponse struct {
	Token   string `json:"token"`
	FeedNum int    `json:"feednum"`
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

		model.UpdateFeed(requestBody.Token, requestBody.Flag, requestBody.FeedNum)

		afterFeedNum := 0
		if requestBody.Flag == 0 {
			afterFeedNum = requestBody.FeedNum - 1
		} else {
			afterFeedNum = requestBody.FeedNum + 1
		}

		response.Success(writer, &feedInformationResponse{
			Token:   requestBody.Token,
			FeedNum: afterFeedNum,
		})
	}
}
