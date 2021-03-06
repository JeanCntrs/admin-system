package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
	"github.com/gorilla/mux"
)

func GetRoleTypes(w http.ResponseWriter, r *http.Request) {
	roleTypes := dal.GetRoleTypes()
	roleTypesByte, _ := json.Marshal(roleTypes)
	fmt.Fprint(w, string(roleTypesByte))
}

func CreateRoleType(w http.ResponseWriter, r *http.Request) {
	roleType := models.RoleType{}
	pageRoles := []models.RolePage{}
	data := json.NewDecoder(r.Body)
	err := data.Decode(&roleType)

	if err != nil {
		panic("An error occurred while decoding role type")
	}

	if roleType.PagesId != "" {
		pagesId := strings.Split(roleType.PagesId, "*")

		for _, v := range pagesId {
			id, _ := strconv.Atoi(v)
			pageRoles = append(pageRoles, models.RolePage{PageId: id, RoleTypeId: roleType.RoleTypeId})
		}
	}

	if roleType.RoleTypeId == 0 {
		registerError := dal.RegisterRoleType(roleType, pageRoles)
		if registerError != nil {
			fmt.Fprintf(w, "0")
			return
		}
	} else {
		updateError := dal.UpdateRoletypeItems(roleType, pageRoles)
		if updateError != nil {
			fmt.Fprintf(w, "0")
			return
		}
	}

	fmt.Fprintf(w, "1")
}

func GetRoleTypeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roloeTypeId := vars["roleTypeId"]
	roloeTypeIdConv, _ := strconv.Atoi(roloeTypeId)

	roleTypes := dal.GetRoleTypeById(roloeTypeIdConv)
	roleTypesByte, _ := json.Marshal(roleTypes)
	fmt.Fprint(w, string(roleTypesByte))
}
