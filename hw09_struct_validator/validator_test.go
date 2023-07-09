package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	AppBadValidateTag struct {
		Version string `validate:"len:1:1"`
	}

	AppBadLenTag struct {
		Version string `validate:"len:five"`
	}

	AppBadRegexpTag struct {
		Version string `validate:"regexp:+"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	AppArray struct {
		Versions []string `validate:"len:5"`
	}

	ResponseBadInTag struct {
		Code int `validate:"in:200,a"`
	}

	ResponseBadMinTag struct {
		Code int `validate:"min:six"`
	}

	Response struct {
		Code int `validate:"in:200,404,500"`
	}

	ResponseArray struct {
		Code []int `validate:"in:200,404,500"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "Not a struct",
			expectedErr: NewIllegalArgumentError("argument must be slice"),
		},
		{
			in: struct{}{},
		},
		{
			in:          AppBadValidateTag{Version: "ios X"},
			expectedErr: NewIllegalArgumentError("validation tag must be in form key:value"),
		},
		{
			in:          AppBadLenTag{Version: "ios X"},
			expectedErr: NewIllegalArgumentError("illegal value in tag for filed 'Version', must be number, got 'five'"),
		},
		{
			in:          AppBadRegexpTag{Version: "ios X"},
			expectedErr: NewIllegalArgumentError("illegal regexp value for filed 'Version', must be valid regexp, got '+'"),
		},
		{
			in: App{Version: "ios X"},
		},
		{
			in: App{Version: "ios"},
			expectedErr: NewValidationErrors(
				ValidationError{
					Field: "Version",
					Err:   NewFieldValidationError("len must be 5, got 3 for 'ios'"),
				},
			),
		},
		{
			in: AppArray{
				Versions: []string{"ios X", "ios Y"},
			},
		},
		{
			in: AppArray{
				Versions: []string{"ios X", "ios Lion", "Android", "ios Y"},
			},
			expectedErr: NewValidationErrors(
				ValidationError{
					Field: "Versions",
					Err:   NewFieldValidationError("len must be 5, got 8 for 'ios Lion'"),
				},
				ValidationError{
					Field: "Versions",
					Err:   NewFieldValidationError("len must be 5, got 7 for 'Android'"),
				},
			),
		},
		{
			in:          ResponseBadInTag{Code: 200},
			expectedErr: NewIllegalArgumentError("illegal value in tag for filed 'Code', must be number, got 'a'"),
		},
		{
			in:          ResponseBadMinTag{Code: 200},
			expectedErr: NewIllegalArgumentError("illegal value in tag for filed 'Code', must be number, got 'six'"),
		},
		{
			in: Response{Code: 200},
		},
		{
			in: Response{Code: 777},
			expectedErr: NewValidationErrors(
				ValidationError{
					Field: "Code",
					Err:   NewFieldValidationError("fieldValue must be one of [200 404 500] values, given 777"),
				},
			),
		},
		{
			in: ResponseArray{
				Code: []int{200, 404},
			},
		},
		{
			in: ResponseArray{
				Code: []int{200, 300, 404, 900},
			},
			expectedErr: NewValidationErrors(
				ValidationError{
					Field: "Code",
					Err:   NewFieldValidationError("fieldValue must be one of [200 404 500] values, given 300"),
				},
				ValidationError{
					Field: "Code",
					Err:   NewFieldValidationError("fieldValue must be one of [200 404 500] values, given 900"),
				},
			),
		},
		{
			in: User{
				ID:     "asdb-vsdrds-asdff-12dfss-asdfe-asdfa",
				Age:    19,
				Email:  "abc@asb.dd",
				Role:   "admin",
				Phones: []string{"+7910455112"},
				meta:   json.RawMessage{},
			},
		},
		{
			in: User{
				ID:     "12",
				Age:    9,
				Email:  "abc@",
				Role:   "user",
				Phones: []string{"1", "2"},
			},
			expectedErr: NewValidationErrors(
				ValidationError{
					Field: "ID",
					Err:   NewFieldValidationError("len must be 36, got 2 for '12'"),
				},
				ValidationError{
					Field: "Age",
					Err:   NewFieldValidationError("min value 18, got 9"),
				},
				ValidationError{
					Field: "Email",
					Err:   NewFieldValidationError("fieldValue must match regexp '^\\w+@\\w+\\.\\w+$', actual value 'abc@'"),
				},
				ValidationError{
					Field: "Role",
					Err:   NewFieldValidationError("fieldValue must be one of [admin stuff] values, given 'user'"),
				},
				ValidationError{
					Field: "Phones",
					Err:   NewFieldValidationError("len must be 11, got 1 for '1'"),
				},
				ValidationError{
					Field: "Phones",
					Err:   NewFieldValidationError("len must be 11, got 1 for '2'"),
				},
			),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			require.Equal(t, tt.expectedErr, Validate(tt.in))
		})
	}
}
