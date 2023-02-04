package forms

import "testing"

func Test_errors_Get(t *testing.T) {
	errorsValue := make(errors)
	errorsValue.Add("error_field", "error message")
	type args struct {
		field string
	}
	tests := []struct {
		name string
		e    errors
		args args
		want string
	}{
		{
			"should have error",
			errorsValue,
			args{"error_field"},
			"error message",
		},
		{
			"shouldn't have error",
			errorsValue,
			args{"not_error_field"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Get(tt.args.field); got != tt.want {
				t.Errorf("errors.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
