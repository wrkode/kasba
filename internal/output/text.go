package output

import (
	"github.com/wrkode/kasba/internal/templates"
	"os"
)
import "text/template"

func AsText(data TemplateData) error {
	tmpl, err := template.New("as_text").Parse(templates.Text)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "as_text", data)
	if err != nil {
		return err
	}
	return nil
}
