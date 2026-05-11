package phonenumber

import "strconv"

func IsValid(pn string) bool {
	if len(pn) != 11 {

		return false
	}

	if pn[0:2] != "09" {

		return false
	}

	if _, err := strconv.Atoi(pn[2:]); err != nil {
		return false
	}

	return true
}
