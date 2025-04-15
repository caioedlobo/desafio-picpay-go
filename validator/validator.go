package validator

import "regexp"

var EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var CpfRegex = regexp.MustCompile(`^\d{11}$`)
var CnpjRegex = regexp.MustCompile(`^\d{14}$`)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
