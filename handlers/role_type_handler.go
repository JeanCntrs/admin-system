package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/models"
)

func GetRoleTypes(w http.ResponseWriter, r *http.Request) {
	roleTypes := dal.GetRoleTypes()
	roleTypesByte, _ := json.Marshal(roleTypes)
	fmt.Fprint(w, string(roleTypesByte))
}

func CreateUserType(w http.ResponseWriter, r *http.Request) {
	roleType := models.RoleType{RoleTypeId: 0, Name: "General supervisor", Description: "Monitor everything"}
	listRolePage := []models.RolePage{
		{PageId: 1},
		{PageId: 2},
	}

	if roleType.RoleTypeId == 0 {
		registerError := dal.RegisterRoleType(roleType, listRolePage)
		if registerError != nil {
			fmt.Fprintf(w, "0")
			return
		}
	}

	fmt.Fprintf(w, "1")
}
