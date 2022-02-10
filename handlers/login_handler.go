package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	utils.DeleteCookie(w, "userId", "")
	utils.RenderTemplate(w, "login", nil)
}
