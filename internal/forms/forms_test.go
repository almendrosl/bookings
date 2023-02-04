package forms

import (
	"net/url"
	"testing"
)

const fieldFirstName = "first_name"
const fieldLastName = "last_name"
const fieldEmail = "email"

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	noErrorsForm := New(postedData)
	errorsForm := New(postedData)
	errorsForm.Errors.Add("field", "error message")
	tests := []struct {
		name string
		f    *Form
		want bool
	}{
		{
			"test no error form",
			noErrorsForm,
			true,
		},
		{
			"test error form",
			errorsForm,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Valid(); got != tt.want {
				t.Errorf("Form.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	postedData.Add(fieldFirstName, "Pete")
	postedData.Add(fieldLastName, "Argent")
	postedData.Add(fieldEmail, "pete@lala.com")
	mockForm := New(postedData)
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		f    *Form
		args args
		want bool
	}{
		{
			"Required data should be in form",
			mockForm,
			args{[]string{fieldFirstName, fieldLastName, fieldEmail}},
			true,
		},
		{
			"Required data isn't in the form",
			mockForm,
			args{[]string{"requiredField"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Required(tt.args.fields...); got != tt.want {
				t.Errorf("Form.Required() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add(fieldFirstName, "Pete")
	postedData.Add(fieldLastName, "Argent")
	postedData.Add(fieldEmail, "pete@lala.com")
	mockForm := New(postedData)
	type args struct {
		field string
	}
	tests := []struct {
		name string
		f    *Form
		args args
		want bool
	}{
		{
			"Required filed should be in form",
			mockForm,
			args{fieldFirstName},
			true,
		},
		{
			"Required data isn't in the form",
			mockForm,
			args{"requiredField"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Has(tt.args.field); got != tt.want {
				t.Errorf("Form.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add(fieldFirstName, "Pete")
	mockForm := New(postedData)
	type args struct {
		field  string
		length int
	}
	tests := []struct {
		name string
		f    *Form
		args args
		want bool
	}{
		{
			"The min length should pass",
			mockForm,
			args{fieldFirstName, 3},
			true,
		},
		{
			"The min length shouldn't pass'",
			mockForm,
			args{fieldFirstName, 6},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.MinLength(tt.args.field, tt.args.length); got != tt.want {
				t.Errorf("Form.MinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add(fieldEmail, "name@name.com")
	mockFormCorrectEmail := New(postedData)

	postedDataFailEmail := url.Values{}
	postedDataFailEmail.Add(fieldEmail, "falseEmail@fail")
	mockFormFailEmail := New(postedDataFailEmail)
	type args struct {
		field string
	}
	tests := []struct {
		name string
		f    *Form
		args args
		want bool
	}{
		{
			"The email format should be correct",
			mockFormCorrectEmail,
			args{fieldEmail},
			true,
		},
		{
			"The email format should be incorrect",
			mockFormFailEmail,
			args{fieldEmail},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsEmail(tt.args.field); got != tt.want {
				t.Errorf("Form.IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
