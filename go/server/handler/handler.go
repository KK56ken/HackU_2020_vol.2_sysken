package handler

import (
	"fmt"
	"net/http"

	"hacku_vol2/server/model"
)

func kkk(w http.ResponseWriter, r *http.Request) {
	var values []string

	values = model.Db_access()

	for i := 0; i < 1; i++ {
		fmt.Fprintln(w, values[i])
	}
}
