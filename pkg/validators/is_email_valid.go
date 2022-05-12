package validators

import "regexp"

func IsEmailValid(email string) bool {
	var rxEmail = regexp.MustCompile("[A-Za-z0-9!#$%&'()*+,./:;<=>?@[\\^_`{|}~]")

	if len(email) < 3 || len(email) > 254 || !rxEmail.MatchString(email) {
		return false
	}
	return true
}
