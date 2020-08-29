package handler
import (
	"errors"
	"fmt"
	"hacku_vol2/http/response"
	"hacku_vol2/server/model"
	"log"
	"net/http"
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
		todoListAlls, err := model.SelectGettingTodo(user_id)
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
		fmt.Println("todoListALL=", todoListAlls)
		fmt.Println("3")
		todoLists := make([]todoList, 100)
		fmt.Println(todoLists)
		// for i := range todoListAlls {
		// 	todoLists[i] = todoList{
		// 		//Name:          "1",                         //タスク名
		// 		//Limit:         todoListAll.Limit,         //期限2020,08,31
		// 		//Subject:       "sa",                        //科目
		// 		//End_flag:      1,                           //終わっているかどうか
		// 		//Register_date: "todoListAll.Register_date", //登録日時
		// 		//End_date:      "todoListAll.End_date",      //終了日時
		// 	}
		// 	fmt.Println(todoLists[i])
		// }
		//response.Success(writer, &todoListResponse{
		//	Todos: todoLists,
		//})
	}
}
