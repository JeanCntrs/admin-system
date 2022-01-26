package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
)

func GetRoleTypes(w http.ResponseWriter, r *http.Request) {
	roleTypes := dal.GetRoleTypes()
	roleTypesByte, _ := json.Marshal(roleTypes)
	fmt.Fprint(w, string(roleTypesByte))
}
