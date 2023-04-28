package helper

import (
	"fmt"
	"ggt/tools"
	"ggt/types"
	"io/ioutil"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/kylelemons/godebug/diff"
	"github.com/sirupsen/logrus"
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

	fmt.Println("正在获取变更文件内容...")
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

	fmt.Println("正在获取最新的文件树...")
	latestTree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	treeStatus, err := workTree.Status()
	if err != nil {
		return nil, err
	}

	for filename, status := range treeStatus {
		if status.Worktree == git.Unmodified ||
			status.Worktree == git.Renamed ||
			status.Worktree == git.Copied ||
			tools.IngoreFile(filename, filter) {
			continue
		}
		fmt.Printf("正在获取文件 %s 的内容...\n", filename)
		workTreeFile, err := workTree.Filesystem.Open(filename)
		if err != nil {
			// check if the file was deleted or was not existed
			if status.Worktree == git.Deleted || status.Worktree == git.Renamed {
				logrus.Warnf("File %s was deleted or renamed. Skip it!", filename)
				continue
			}
			return nil, err
		}

		defer workTreeFile.Close()

		objectFile, err := latestTree.File(filename)
		if err != nil {
			if err == object.ErrFileNotFound {
				// file was added
				dat, err := ioutil.ReadAll(workTreeFile)
				if err != nil {
					return nil, err
				}

				diffContent[filename] = printDiff("", string(dat))
				continue
			}
			return nil, err
		}

		objectContent, err := objectFile.Contents()
		if err != nil {
			return nil, err
		}

		workTreeContent, err := ioutil.ReadAll(workTreeFile)
		if err != nil {
			return nil, err
		}

		if string(workTreeContent) != objectContent {
			// file was modified
			diffContent[filename] = printDiff(string(objectContent), string(workTreeContent))
		}

	}
	// parent, err := commit.Parent(0)
	// if err != nil {
	// 	return nil, err
	// }

	// parentTree, err := parent.Tree()
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Printf("%s --> %s", ref.Hash().String(), parent.Hash.String())
	// changes, err := object.DiffTree(parentTree, headTree)
	// if err != nil {
	// 	return nil, err
	// }

	// for _, change := range changes {
	// 	name := ""
	// 	showDiff := false
	// 	if change.From.Name != "" {
	// 		name = change.From.Name
	// 		showDiff = true
	// 	} else if change.To.Name != "" {
	// 		name = change.To.Name
	// 		showDiff = true
	// 	}

	// 	if showDiff &&
	// 		!tools.IngoreFile(name, filter) {
	// 		patch, err := change.Patch()
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		// fmt.Println(patch.String())
	// 		diffContent[name] = patch.String()
	// 	}

	// }

	return diffContent, nil

}

func printDiff(old, new string) string {
	return diff.Diff(old, new)
}
