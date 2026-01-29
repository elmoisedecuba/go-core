package validators

import (
	"github.com/asaskevich/govalidator"
)

func Email(text string) bool {
	d := govalidator.IsEmail(text)
	return d
}

func Password(text string) bool {
	d := govalidator.MinStringLength(text, "8")
	return d
}

func Names(text string) bool {
	d := govalidator.MinStringLength(text, "3")
	return d
}

func Username(text string) bool {
	u := govalidator.IsLowerCase(text)
	c := govalidator.IsAlphanumeric(text)
	l := govalidator.MinStringLength(text, "6")
	n := govalidator.IsNull(text)
	if u && c && l && n {
		return true
	}
	return false
}

func Token(text string) bool {
	d := govalidator.IsSHA256(text)
	return d
}
