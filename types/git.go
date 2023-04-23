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
		"package.json",
		"yarn.lock",
	}
}

// GitIgnoreSuffix returns a list of gitignore file suffixes
func GitIgnoreSuffix() []string {
	return []string{
		".png",
		".jpg",
	}
}

// GitDirIgnoreList returns a list of gitignore dir
func GitDirIgnoreList() []string {
	return []string{
		"vendor",
		"node_modules",
		"bin",
		".DS_Store",
		".idea",
		".git",
	}
}

type FilterType int

const FileNameFilter FilterType = 1
const SuffixFilter FilterType = 2
const DirFilter FilterType = 3
