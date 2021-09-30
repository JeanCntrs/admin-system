package handlers

import (
	"net/http"

	"github.com/JeanCntrs/admin-system/utils"
)

func Person(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "person", nil)
}
