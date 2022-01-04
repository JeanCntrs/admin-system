package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Country(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "country", nil)
}
