package iteration

import "strings"

func Repeat(character string, repeat int) string {
	var return_string strings.Builder

	for range repeat {
		return_string.WriteString(character)
	}
	return return_string.String()
}
