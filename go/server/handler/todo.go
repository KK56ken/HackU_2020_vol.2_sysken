package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"hacku_vol2/http/response"
	"hacku_vol2/server/model"
)

type todoList struct {
	Name         string `json:"name"`
	Limit        string `json:"time_limite"`
	Subject      string `json:"subject"`
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

type ToDoGet struct {
	Token string `json:"token"`
}

func HandleToDoGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody ToDoGet
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		todoListAlls, err := model.SelectGettingTodo(requestBody.Token)
		//todoListAlls, err := model.SelectGettingTodo()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		var subject string
		todoLists := make([]todoList, len(todoListAlls))
		for i, to := range todoListAlls {

			subject = ""

			fmt.Println(to.SubjectId)

			switch to.SubjectId {
			case 1:
				subject = "国語"
			case 2:
				subject = "算数"
			case 3:
				subject = "英語"
			case 4:
				subject = "理科"
			case 5:
				subject = "社会"
			case 6:
				subject = "その他"
			}
			todoLists[i] = todoList{
				Name:         to.Name,         //タスク名
				Limit:        to.Limit,        //期限2020,08,31
				Subject:      subject,         //科目
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

func HandleToDoEnd() http.HandlerFunc {
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

		err := model.InsertTodoEnd(&model.Task{
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
