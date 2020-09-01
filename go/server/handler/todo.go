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
	Limit        string `json:"time_limite"`
	SubjectId    int    `json:"subject_id"`
	EndFlag      int    `json:"end_flag"`
	RegisterDate string `json:"register_date"`
	EndDate      string `json:"end_date"`
}
type todoListResponse struct {
	Todos []todoList `json:"todos"`
}
type addTodoListRequest struct {
	Token     string `json:"token"`
	Name      string `json:"name"`
	Limit     string `json:"time_limite"`
	SubjectId int    `json:"subject_id"`
}
type addTodoListResponse struct {
	Name      string `json:"name"`
	Limit     string `json:"time_limite"`
	SubjectId int    `json:"subject_id"`
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
				Name:         to.Name,         //タスク名
				Limit:        to.Limit,        //期限2020,08,31
				SubjectId:    1,               //科目
				EndFlag:      to.EndFlag,      //終わっているかどうか
				RegisterDate: to.RegisterDate, //登録日時
				EndDate:      to.EndDate,      //終了日時
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
		var requestBody addTodoListRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		user_id := model.SelectUserId(requestBody.Token)
		subject_id := requestBody.SubjectId
		name := requestBody.Name
		limit := requestBody.Limit

		err := model.InsertTodo(&model.Task{
			UserId:    user_id,
			SubjectId: requestBody.SubjectId,
			Name:      requestBody.Name,
			Limit:     requestBody.Limit,
		})

		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error ww")
			return
		}

		response.Success(writer, &addTodoListResponse{
			SubjectId: subject_id,
			Name:      name,
			Limit:     limit,
		})
	}
}
