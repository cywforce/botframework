package config

import "regexp"

func uuidValidate(value string) bool {
	match,_ := regexp.MatchString("^[0-9a-f]{8}-?[0-9a-f]{4}-?[1-5][0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$",value)
	return match;
}
