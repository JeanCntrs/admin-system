package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login", nil)
}
