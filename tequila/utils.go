package tequila

import "strings"

func SubString(str string, subString string) string {
	index := strings.Index(str, subString)
	if index > -1 {
		return str[index+len(subString):]
	}
	return ""
}
