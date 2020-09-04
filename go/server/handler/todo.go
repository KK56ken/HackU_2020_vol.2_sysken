package handler

import (
	"encoding/json"
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

// リクエスト(送られてきた値
type addTodoListRequest struct {
	Token   string `json:"token"`
	Name    string `json:"name"`
	Limit   string `json:"time_limite"`
	Subject string `json:"subject"`
}

// レスポンス(返す値
type addTodoListResponse struct {
	Name    string `json:"name"`
	Limit   string `json:"time_limite"`
	Subject string `json:"subject"`
}

type todoListResponse struct {
	Todos []todoList `json:"todos"`
}

// todo全件取得 リクエスト
type ToDoGetAllRequest struct {
	Token string `json:"token"`
}

// todo OK
// todo全件取得
func HandleToDoGetAll() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody ToDoGetAllRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		todoListAlls, err := model.SelectGettingTodo(requestBody.Token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		var subject string
		todoLists := make([]todoList, len(todoListAlls))
		for i, to := range todoListAlls {

			subject = ""
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
		userId := model.SelectUserId(requestBody.Token)
		subject := requestBody.Subject
		name := requestBody.Name
		limit := requestBody.Limit

		subjectId := 0
		switch requestBody.Subject {
		case "国語":
			subjectId = 1
		case "算数":
			subjectId = 2
		case "英語":
			subjectId = 3
		case "理科":
			subjectId = 4
		case "社会":
			subjectId = 5
		case "その他":
			subjectId = 6
		}

		err := model.InsertTodo(&model.Task{
			UserId:    userId,
			SubjectId: subjectId,
			Name:      requestBody.Name,
			Limit:     requestBody.Limit,
		})

		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error ww")
			return
		}

		response.Success(writer, &addTodoListResponse{
			Subject: subject,
			Name:    name,
			Limit:   limit,
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

		userId := model.SelectUserId(requestBody.Token)
		subject := requestBody.Subject
		name := requestBody.Name
		limit := requestBody.Limit

		subjectId := 0
		switch requestBody.Subject {
		case "国語":
			subjectId = 1
		case "算数":
			subjectId = 2
		case "英語":
			subjectId = 3
		case "理科":
			subjectId = 4
		case "社会":
			subjectId = 5
		case "その他":
			subjectId = 6
		}

		err := model.InsertTodoEnd(&model.Task{
			UserId:    userId,
			SubjectId: subjectId,
			Name:      requestBody.Name,
			Limit:     requestBody.Limit,
		})

		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error ww")
			return
		}

		response.Success(writer, &addTodoListResponse{
			Name:    name,
			Limit:   limit,
			Subject: subject,
		})
	}
}
