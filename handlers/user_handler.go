package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/JeanCntrs/admin-system/utils"
	"github.com/gorilla/mux"
)

func User(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := dal.GetUsers()
	usersByte, _ := json.Marshal(users)
	fmt.Fprint(w, string(usersByte))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	userIdConv, _ := strconv.Atoi(userId)

	users := dal.GetUserById(userIdConv)
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
	} else {
		_, err := dal.UpdateUser(user)
		if err != nil {
			fmt.Fprintf(w, "0")
		}
	}

	fmt.Fprintf(w, "1")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idConv, err := strconv.Atoi(id)

	if err != nil {
		fmt.Fprintf(w, "0")
		return
	}

	_, errorFound := dal.DeleteUser(idConv)

	if errorFound != nil {
		fmt.Fprintf(w, "0")
		return
	}

	fmt.Fprintf(w, "1")
}

func EnterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]

	encryptedPassword := utils.Encrypt(password)
	foundUser := dal.ValidateExistingUser(username, encryptedPassword)
	if foundUser == "1" {
		userId := dal.GetUserId(username, encryptedPassword)
		utils.CreateCookie(w, "userId", userId)
	}

	fmt.Fprint(w, foundUser)
}
