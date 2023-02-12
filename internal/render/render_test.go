package render

import (
	"net/http"
	"testing"

	"github.com/almendrosl/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not fount in session")
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error("could not get template from template cache", err)
	}

	app.TemplateCache = tc

	var ww myWriter
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = Template(&ww, r, "non-existing.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}

func TestNewRenderer(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()

	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}
