package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JeanCntrs/admin-system/dal"
	"github.com/JeanCntrs/admin-system/utils"
)

func Page(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "page", nil)
}

func GetPages(w http.ResponseWriter, r *http.Request) {
	pages := dal.GetPages()
	pagesByte, _ := json.Marshal(pages)
	fmt.Fprint(w, string(pagesByte))
}
