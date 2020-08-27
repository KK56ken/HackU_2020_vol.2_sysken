package handler

import (
	"fmt"
	"github.com/HackU_2020_vol.2_sysken/go/model"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var values []string

	values = model.Db_access()

	for i := 0; i < 1; i++ {
		fmt.Fprintln(w, values[i])
	}
}