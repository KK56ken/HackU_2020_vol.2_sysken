package server

import (
	"fmt"
	"log"
	"net/http"

	"hacku_vol2/server/handler"
)

func Serve(addr string) {
	fmt.Println("1")
	//http.HandleFunc("/", handler.kkk)
	// get
	// todo全件取得
	http.HandleFunc("/todo/get", get(handler.HandleToDoGet()))
	//　ユーザーのログイン
	http.HandleFunc("/user/login", get(handler.HandleUserLogin()))

	// post
	//　todoの登録
	http.HandleFunc("/todo/register", post(handler.HandleToDoPost()))
	// todoの終了
	http.HandleFunc("/todo/end", post(handler.HandleToDoEnd()))
	// user登録
	http.HandleFunc("/user/signup", post(handler.HandleUserSignup()))
	// 餌
	http.HandleFunc("/raise/feed", post(handler.HandleFeedChange()))

	http.ListenAndServe(":3001", nil)

	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
