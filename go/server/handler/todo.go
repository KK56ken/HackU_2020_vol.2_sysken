package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"hacku_vol2/http/response"
	"hacku_vol2/server/model"
)

type todoList struct {
	Name         string `json:"name"`
	Limit        string `json:"Limit"`
	Subject      string `json:"subject"`
	EndFlag      int    `json:"end_flag"`
	RegisterDate string `json:"register_date"`
	EndDate      string `json:"end_date"`
}
type todoListResponse struct {
	Todos []todoList `json:"todos"`
}
type addTodo struct {
	UserId    int    `json:"user_id"`
	SubjectId int    `json:"subject_id"`
	Name      string `json:"name"`
	Limit     string `json:"Limit"`
	EndFlag   int    `json:"end_flag"`
	EndDate   string `json:"end_date"`
}

func HandleToDoGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		userId := 1 //本番環境では消して
		todoListAlls, err := model.SelectGettingTodo(userId)
		//todoListAlls, err := model.SelectGettingTodo()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		if todoListAlls == nil {
			log.Println(errors.New("usersRanking not found"))
			response.BadRequest(writer, fmt.Sprintf("usersRanking not found"))
			return
		}

		todoLists := make([]todoList, len(todoListAlls))
		for i, to := range todoListAlls {
			todoLists[i] = todoList{
				Name:         to.Name,                 //タスク名
				Limit:        to.Limit,                //期限2020,08,31
				Subject:      "todoListAlls[i].Limit", //科目
				EndFlag:      to.EndFlag,              //終わっているかどうか
				RegisterDate: to.RegisterDate,         //登録日時
				EndDate:      to.EndDate,              //終了日時
			}
		}

		response.Success(writer, &todoListResponse{
			Todos: todoLists,
		})
	}
}
func HandleToDoPost() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody addTodo
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		err := model.InsertTodo(&model.Task{
			UserId:    requestBody.UserId,
			SubjectId: requestBody.SubjectId,
			Name:      requestBody.Name,
			Limit:     requestBody.Limit,
			EndFlag:   requestBody.EndFlag,
			EndDate:   requestBody.EndDate,
		})
		fmt.Println(requestBody.UserId)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error ww")
			return
		}

		response.Success(writer, &addTodo{
			Name: requestBody.Name,
		})
	}
}
