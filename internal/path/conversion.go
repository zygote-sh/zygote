package path

import (
	"github.com/zygote-sh/zygote/internal/template"
)

func (p Path) String() string {
	return string(p)
}

func (p Path) Template() template.Template {
	return template.Template(p.String())
}

func (p Path) Parse() template.Template {
	return p.Template().Parse()
}
