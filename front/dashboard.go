package front

import (
	"net/http"
	"text/template"

	"github.com/etcinit/deis-dashboard/util"
)

// GetDashboard handle main dashboard route
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	title := "Top Requests"
	t, _ := template.ParseFiles("../resources/views/dashboard.html")

	p := &util.Page{Title: title}
	t.Execute(w, p)
}
