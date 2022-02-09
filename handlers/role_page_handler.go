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

func RolePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "role_page", nil)
}

func GetPageByRoleTypeId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleTypeId := vars["roleTypeId"]
	roleTypeIdConv, _ := strconv.Atoi(roleTypeId)

	pages := dal.GetPageByRoleTypeId(roleTypeIdConv)
	pagesByte, _ := json.Marshal(pages)
	fmt.Fprint(w, string(pagesByte))
}

func CreateRolePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "create_role_page", nil)
}

func EditRolePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roleTypeId := vars["roleTypeId"]
	roleTypeIdConv, _ := strconv.Atoi(roleTypeId)
	roleType := models.RoleType{RoleTypeId: roleTypeIdConv}
	utils.RenderTemplate(w, "edit_role_page", roleType)
}
