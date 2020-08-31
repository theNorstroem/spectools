package template

import (
	"github.com/Masterminds/sprig"
	"html/template"
	"log"
)

var TypeGen *template.Template
var ServiceGen *template.Template

func init() {
	fn := sprig.FuncMap()
	fn["noescape"] = noescape
	var err error
	TypeGen, err = template.New("typespec").Funcs(fn).Parse(TypeTPL)
	ServiceGen, err = template.New("servicespec").Funcs(fn).Parse(ServiceTPL)
	if err != nil {
		log.Fatal(err)
	}
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}
