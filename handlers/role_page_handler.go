package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func RolePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "role_page", nil)
}

func CreateRolePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "create_role_page", nil)
}

func EditRolePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "edit_role_page", nil)
}
