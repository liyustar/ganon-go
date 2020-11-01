package gweb

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"nuts/gweb/data"
)

func Session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return
	}

	sess = data.Session{Uuid: cookie.Value}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("invalid session")
	}
	return
}


func GenerateHTML(w http.ResponseWriter, data interface{}, fn ...string) error {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	return templates.ExecuteTemplate(w, "layout", data)
}