package tools

import (
	"ggt/types"
	"strings"
)

// GetIngoreList returns a list of gitignore file
func GetIngoreList() map[types.FilterType][]string {
	return map[types.FilterType][]string{
		types.FileNameFilter: types.GitIgnoreList(),
		types.SuffixFilter:   types.GitIgnoreSuffix(),
		types.DirFilter:      types.GitDirIgnoreList(),
	}
}

// IngoreFile It will check given file whether it is in the ignore list
func IngoreFile(file string, ignoreList map[types.FilterType][]string) bool {
	for filterType, filterValue := range ignoreList {
		switch filterType {
		case types.FileNameFilter:
			for _, value := range filterValue {
				if file == value {
					return true
				}
			}
		case types.SuffixFilter:
			for _, value := range filterValue {
				if strings.HasSuffix(file, value) {
					return true
				}
			}
		case types.DirFilter:
			for _, value := range filterValue {
				if strings.HasPrefix(file, value) {
					return true
				}
			}
		}
	}
	return false
}
