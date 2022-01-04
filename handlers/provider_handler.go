package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Provider(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "provider", nil)
}
