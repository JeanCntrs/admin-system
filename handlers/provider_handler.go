package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Provider(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "provider", nil)
}

func GetProviders(w http.ResponseWriter, r *http.Request) {
	providers := dal.GetProviders()
	providersByte, _ := json.Marshal(providers)
	fmt.Fprint(w, string(providersByte))
}
