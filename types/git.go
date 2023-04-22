package types

// GitIgnoreList returns a list of gitignore file
func GitIgnoreList() map[string]bool {
	return map[string]bool{
		".gitignore":        true,
		"go.mod":            true,
		"go.sum":            true,
		"vendor":            true,
		"node_modules":      true,
		"package-lock.json": true,
		"yarn.lock":         true,
		"package.json":      true,
	}
}

// GitIgnoreSuffix returns a list of gitignore file suffixes
func GitIgnoreSuffix() map[string]bool {

	return map[string]bool{
		".png": true,
		".jpg": true,
	}
}
