package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode())

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, template string, data jet.VarMap) error {
	view, err := views.GetTemplate(template)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
