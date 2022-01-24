package handlers

import (
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		UserId:     0,
		Username:   "username 001",
		Password:   "pwd",
		PersonId:   5,
		RoleTypeId: 1,
	}
	// data := json.NewDecoder(r.Body)
	// err := data.Decode(&user)

	// if err != nil {
	// 	panic("An error occurred while decoding country")
	// }

	if user.UserId == 0 {
		err := dal.RegisterUserTx(user)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}
