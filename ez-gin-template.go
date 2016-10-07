package eztemplate

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin/render"
)

type Render struct {
	Templates       map[string]*template.Template
	TemplatesDir    string
	Layout          string
	Ext             string
	TemplateFuncMap map[string]interface{}
	Debug           bool
}

func New() Render {
	r := Render{

		Templates: map[string]*template.Template{},
		// TemplatesDir holds the location of the templates
		TemplatesDir: "app/views/",
		// Layout is the file name of the layout file
		Layout: "layouts/base",
		// Ext is the file extension of the rendered templates
		Ext: ".html",
		// Template's function map
		TemplateFuncMap: nil,
		// Debug enables debug mode
		Debug: false,
	}

	return r
}

func (r Render) Init() Render {
	layout := r.TemplatesDir + r.Layout + r.Ext

	viewDirs, _ := filepath.Glob(r.TemplatesDir + "**/*" + r.Ext)

	for _, view := range viewDirs {
		rendername := getRenderName(view)
		r.AddFromFiles(rendername, layout, view)
	}

	return r
}

func getRenderName(tpl string) string {
	dir, file := filepath.Split(tpl)
	dir = strings.Replace(dir, "app/views/", "", 1)
	file = strings.TrimSuffix(file, ".html")
	return dir + file
}

func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	r.Templates[name] = tmpl
}

func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.New(filepath.Base(r.Layout + r.Ext)).Funcs(r.TemplateFuncMap).ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.Templates[name],
		Data:     data,
	}
}
