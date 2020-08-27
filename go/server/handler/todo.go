package handler

import (
	"errors"
	"fmt"
	"hacku_vol2/http/response"
	"log"
	"net/http"
	"hacku_vol2/server/model"
)

type todoList struct {
	Name          string `json:"name"`
	Limit         string `json:"Limit"`
	Subject       string `json:"subject"`
	End_flag      int    `json:"end_flag"`
	Register_date string `json:"register_date"`
	End_date      string `json:"end_date"`
}

type todoListResponse struct {
	Todos []todoList `json:"todos"`
}

func HandleToDoGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("2")

		user_id := 1 //本番環境では消して
		todoListAll, err := model.SelectGettingTodo(user_id)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		if todoListAll == nil {
			log.Println(errors.New("usersRanking not found"))
			response.BadRequest(writer, fmt.Sprintf("usersRanking not found"))
			return
		}

		fmt.Println("todoListALL=",todoListAll)
		fmt.Println("3")

		//todoLists := make([]todoList, len(todoListAll))
		//for i, user := range todoListAll {
		//	todoLists[i] = todoList{
		//		Name:   user.ID,
		//		Limit: user.Name,
		//		Subject:     int32(start + i), //スタート位置
		//		End_flag:    user.HighScore,
		//		Register_date:
		//		End_date:
		//	}
		//}

	}
}
