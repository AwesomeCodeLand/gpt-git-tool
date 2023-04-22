package helper

import (
	"fmt"
	"ggt/tools"
	"ggt/types"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/format/diff"
)

func GetChangeFiles() (map[string]string, error) {
	return getChangeFileContent(".", tools.GetIngoreList())
}

// getChangeFileContent get the content of the file that has been changed
// return a map of file name and content
// filter is a map, stores which files we should ignore.
// key is the file name
// value is the content of the file
func getChangeFileContent(repoPath string, filter map[types.FilterType][]string) (diffContent map[string]string, err error) {

	diffContent = make(map[string]string)

	// Open the git repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	// Get the commit object for HEAD
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	parent, err := commit.Parent(0)
	if err != nil {
		return nil, err
	}

	patch, err := commit.Patch(parent)
	if err != nil {
		return nil, err
	}

	for _, f := range patch.FilePatches() {
		from, to := f.Files()
		name := ""
		showDiff := false
		if from != nil {
			showDiff = true
			name = from.Path()
		}

		if !showDiff {
			if to != nil {
				showDiff = true
				name = to.Path()
			}
		}

		if showDiff &&
			!tools.IngoreFile(name, filter) {
			theDiffContent := ""
			for _, c := range f.Chunks() {
				data := strings.Split(c.Content(), "\n")
				for _, d := range data {
					switch c.Type() {
					case diff.Add:
						theDiffContent = fmt.Sprintf("%s+%s\n", theDiffContent, d)
					case diff.Delete:
						theDiffContent = fmt.Sprintf("%s-%s\n", theDiffContent, d)
					case diff.Equal:
						theDiffContent = fmt.Sprintf("%s%s\n", theDiffContent, d)
					}
				}
			}
			diffContent[name] = theDiffContent
		}

	}
	return diffContent, nil

}
