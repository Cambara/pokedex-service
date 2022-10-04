package databaseseed

import "strings"

func ConvertStringInArray(strArray string) []string {

	str := strings.ReplaceAll(strArray, "[", "")
	str2 := strings.ReplaceAll(str, "]", "")
	str3 := strings.ReplaceAll(str2, "'", "")
	result := strings.Split(str3, ",")
	return result
}
