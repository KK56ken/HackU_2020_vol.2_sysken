package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"hacku_vol2/http/response"
	"hacku_vol2/server/model"
)

type userInformationResponse struct {
	Token    string `json:"token"`
	Feed_num int    `json:"feed_num"`
}

type userInformationRequest struct {
	Name     string `json:"name"`
	Password string `json : "password"`
}

func HandleUserSignup() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody userInformationRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		md5 := md5.Sum([]byte(requestBody.Password))
		hashingPass := hex.EncodeToString(md5[:])
		fmt.Println(hashingPass)

		err := model.InsertUserr(&model.User{
			Name:     requestBody.Name,
			Password: requestBody.Password,
			Token:    hashingPass,
			Feed_num: 0,
		})
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error ww")
			return
		}

		response.Success(writer, &userInformationResponse{
			Token: hashingPass,
		})
	}
}

func HandleUserLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody userInformationRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		userData, err := model.SelectUser(requestBody.Name, requestBody.Password)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		if userData == nil {
			log.Println(errors.New("usersRanking not found"))
			response.BadRequest(writer, fmt.Sprintf("usersRanking not found"))
			return
		}

		response.Success(writer, &userInformationResponse{
			Token:    userData[0].Token,
			Feed_num: userData[0].Feed_num,
		})

	}
}
