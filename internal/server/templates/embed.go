package templates

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed form.tmpl
var formRawTemplate string

type FormTemplateParams struct {
}

func (p FormTemplateParams) Render(w io.Writer) error {
	t, err := template.New("form").Parse(formRawTemplate)
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}

//go:embed confirm.tmpl
var confirmRawTemplate string

type ConfirmTemplateParams struct {
}

func (p ConfirmTemplateParams) Render(w io.Writer) error {
	t, err := template.New("confirm").Parse(confirmRawTemplate)
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}

//go:embed view.tmpl
var viewRawTemplate string

type ViewTemplateParams struct {
	ImageFileName string
}

func (p ViewTemplateParams) Render(w io.Writer) error {
	t, err := template.New("view").Parse(viewRawTemplate)
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
