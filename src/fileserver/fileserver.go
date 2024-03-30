package fileserver 

import (
	"html/template"
	"net/http"
)

var T *template.Template
var FS http.Handler

func RunFileServer()  {
  FS = http.FileServer(http.Dir("./static"))
  http.Handle("/static/", http.StripPrefix("/static/", FS))
}
