package form

import (
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errs errors
}

func (f *Form) CheckNotnull(fields ...string) {
	for _, field := range fields {
		s := f.Get(field)
		//it should trim all space ,otherwise when user pass a value of " "(space),it still pass the check
		if strings.TrimSpace(s) == "" {
			f.Errs.Add(field, "This field can't be blank")
		}
	}
}

func New(v url.Values) *Form {
	return &Form{
		Values: v,
		Errs:   newErr(),
	}
}

func (f *Form) Valid() bool {
	return len(f.Errs) == 0
}

func (f *Form) CheckEmail(field string) {
	email := f.Get(field)
	if !govalidator.IsEmail(email) {
		f.Errs.Add(field, "this field must be email format")
	}
}
