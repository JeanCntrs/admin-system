package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "index", nil)
}
