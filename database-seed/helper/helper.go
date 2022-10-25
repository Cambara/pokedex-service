package helper

import "strings"

func ConvertStringInArray(strArray string) []string {

	str := strings.ReplaceAll(strArray, "[", "")
	str2 := strings.ReplaceAll(str, "]", "")
	str3 := strings.ReplaceAll(str2, "'", "")
	data := strings.Split(str3, ",")

	result := make([]string, 0)

	for _, strValue := range data {
		result = append(result, strings.TrimSpace(strValue))
	}

	return result
}
