package main

import "os"

func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func removeTail(slice []int) []int {
	return slice[:len(slice)-1]
}

func ConvertStringSliceToMap(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

func InMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}
