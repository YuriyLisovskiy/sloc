package parser

import "strings"

func GetExt(fileName string) string {
	arr := strings.Split(fileName, ".")
	res := "other"
	if len(arr) > 0 {
		res = arr[len(arr)-1]
	}
	return res
}

func Parse(files []string) ([]Lang, error) {
	return nil, nil
}
