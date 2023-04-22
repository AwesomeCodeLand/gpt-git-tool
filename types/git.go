package types

// GitIgnoreList returns a list of gitignore file
func GitIgnoreList() []string {
	return []string{
		".gitignore",
		"go.mod",
		"go.sum",
		"vendor",
		"node_modules",
		"package-lock.json",
		"yarn.lock",
		"package.json",
	}
}

// GitIgnoreSuffix returns a list of gitignore file suffixes
func GitIgnoreSuffix() []string {
	return []string{
		".png",
		".jpg",
	}
}

type FilterType int

const FileNameFilter FilterType = 1
const SuffixFilter FilterType = 2
