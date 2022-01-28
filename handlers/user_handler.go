package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
)

func User(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := dal.GetUsers()
	usersByte, _ := json.Marshal(users)
	fmt.Fprint(w, string(usersByte))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&user)

	if err != nil {
		panic("An error occurred while decoding country")
	}

	if user.UserId == 0 {
		err := dal.RegisterUserTx(user)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}
