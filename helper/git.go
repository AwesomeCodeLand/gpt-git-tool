package helper

import (
	"fmt"
	"ggt/types"
	"os"
	"strings"

	git "github.com/libgit2/git2go/v34"
	"github.com/sirupsen/logrus"
)

func GetChangeFiles() ([]string, error) {
	files, err := getChangeFiles(".")
	if err != nil {
		return nil, err
	}

	var result []string
	ignoreList := types.GitIgnoreList()
	suffixList := types.GitIgnoreSuffix()

	for _, file := range files {
		isIgnored := false

		for key := range ignoreList {
			if strings.Contains(file, key) {
				isIgnored = true
				break
			}
		}

		if isIgnored {
			continue
		}

		for suffix := range suffixList {
			if strings.HasSuffix(file, suffix) {
				isIgnored = true
				break
			}
		}

		if !isIgnored {
			result = append(result, file)
		}
	}

	return result, nil
}

func GetChangeContentWithFile(path string) {
	getChangeContentWithFile(path)
}

// getChangeFiles returns the files that have been changed
// It is used to get all the changed files via git diff
func getChangeFiles(repoPath string) ([]string, error) {
	var result []string
	// get all file names via git diff --name-only
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		return nil, err
	}

	head, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("Failed to get HEAD reference: %v\n", err)
	}

	commit, err := repo.LookupCommit(head.Target())
	if err != nil {
		return nil, fmt.Errorf("Failed to lookup commit: %v\n", err)
	}

	// parents := commit.Parent(0)
	if commit.ParentCount() == 0 {
		return nil, fmt.Errorf("No parent commits found")
	}

	parentTree, _ := commit.Parent(0).Tree()
	currTree, _ := commit.Tree()
	diffOpts, _ := git.DefaultDiffOptions()

	diff, err := repo.DiffTreeToTree(parentTree, currTree, &diffOpts)
	if err != nil {

		return nil, fmt.Errorf("Failed to create diff: %v\n", err)
	}

	numDeltas, err := diff.NumDeltas()
	if err != nil {
		return nil, fmt.Errorf("Failed to get number of deltas: %v\n", err)
	}

	for i := 0; i < numDeltas; i++ {
		delta, err := diff.GetDelta(i)
		if err != nil {
			logrus.Errorf("Failed to get delta: %v\n", err)
			continue
		}

		if delta.Status == git.DeltaAdded || delta.Status == git.DeltaModified || delta.Status == git.DeltaDeleted {
			result = append(result, delta.NewFile.Path)
		}
	}

	return result, nil
}

func getCurrentAndPreviousCommitID(path string) (*git.Oid, *git.Oid, error) {
	repoPath := "."
	filePath := path

	// Open the repository
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Printf("Failed to open repository: %v\n", err)
		os.Exit(1)
	}

	// Get the HEAD reference
	headRef, err := repo.Head()
	if err != nil {
		fmt.Printf("Failed to get HEAD reference: %v\n", err)
		os.Exit(1)
	}

	// Lookup the commit for the HEAD reference
	headCommit, err := repo.LookupCommit(headRef.Target())
	if err != nil {
		fmt.Printf("Failed to lookup HEAD commit: %v\n", err)
		os.Exit(1)
	}

	// Find the previous commit that modified the file
	var prevCommit *git.Commit

	prevCommit = headCommit.Parent(0)
	// Print the previous commit ID if found
	if prevCommit != nil {
		fmt.Println("Previous commit:", prevCommit.Id().String())
	} else {
		fmt.Println("No previous commit found")
	}

	// Find the current commit that modified the file
	var currCommit *git.Commit
	tree, err := headCommit.Tree()
	if err != nil {
		fmt.Printf("Failed to get HEAD tree: %v\n", err)
		os.Exit(1)
	}
	entry, err := tree.EntryByPath(filePath)
	if err == nil && entry != nil {
		currCommit = headCommit
	}

	// Print the current commit ID if found
	if currCommit != nil {
		fmt.Println("Current commit:", currCommit.Id().String())
	} else {
		fmt.Println("No current commit found")
	}

	return currCommit.Id(), prevCommit.Id(), nil
}

// getChangeContentWithFile
// get the content of the file that has been changed
func getChangeContentWithFile(path string) {
	cmd, pmd, err := getCurrentAndPreviousCommitID(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repoPath := "."
	filePath := path
	oldCommitID := pmd
	newCommitID := cmd

	// Open the repository
	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		fmt.Printf("Failed to open repository: %v\n", err)
		os.Exit(1)
	}

	// Lookup the old commit
	oldCommit, err := repo.LookupCommit(oldCommitID)
	if err != nil {
		fmt.Printf("Failed to lookup old commit: %v\n", err)
		os.Exit(1)
	}

	// Lookup the new commit
	newCommit, err := repo.LookupCommit(newCommitID)
	if err != nil {
		fmt.Printf("Failed to lookup new commit: %v\n", err)
		os.Exit(1)
	}

	// Get the tree for the old commit
	oldTree, err := oldCommit.Tree()
	if err != nil {
		fmt.Printf("Failed to get old tree: %v\n", err)
		os.Exit(1)
	}

	// Get the tree for the new commit
	newTree, err := newCommit.Tree()
	if err != nil {
		fmt.Printf("Failed to get new tree: %v\n", err)
		os.Exit(1)
	}

	// Get the diff between the two trees
	diffOptions, err := git.DefaultDiffOptions()
	if err != nil {
		fmt.Printf("Failed to get default diff options: %v\n", err)
		os.Exit(1)
	}
	diff, err := repo.DiffTreeToTree(oldTree, newTree, &diffOptions)
	if err != nil {
		fmt.Printf("Failed to get diff: %v\n", err)
		os.Exit(1)
	}

	// Find the file in the diff
	var delta *git.DiffDelta
	err = diff.ForEach(func(file git.DiffDelta, progress float64) (git.DiffForEachHunkCallback, error) {
		fmt.Printf("file [%+v]\n", file)
		if file.OldFile.Path == filePath || file.NewFile.Path == filePath {
			delta = &file
			// return nil, git.ErrUser
			return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
				// fmt.Printf("hunk: [%+v]\n", hunk)
				return func(line git.DiffLine) error {
					// fmt.Printf("line: [%+v]\n", line)

					fmt.Printf("+ line: [%+v]\n", line)
					return nil
				}, nil
			}, nil
		}
		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			// fmt.Printf("hunk: [%+v]\n", hunk)
			return func(line git.DiffLine) error {
				// fmt.Printf("- line: [%+v]\n", line.Content)
				return nil
			}, nil
		}, nil
	}, 100)

	if err != nil {
		fmt.Printf("Failed to find file in diff: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("delta [%+v]\n", delta)
	// fmt.Printf("diff [%+v]\n", diff.Patch(int(delta.Status)))
	// Get the patch for the file
	patch, err := diff.Patch(0)
	if err != nil {
		fmt.Printf("Failed to get patch: %v\n", err)
		os.Exit(1)
	}

	// Print the patch
	patchStr, err := patch.String()
	if err != nil {
		fmt.Printf("Failed to convert patch to string: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(patchStr)
}

// func getChangeContentWithFile(path string) {
// 	// 打开仓库
// 	repo, err := git.OpenRepository(".")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// 获取最新提交对象
// 	head, err := repo.Head()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	commit, err := repo.LookupCommit(head.Target())
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// 获取指定文件的 blob 对象
// 	tree, err := commit.Tree()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	entry, err := tree.EntryByPath(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	blob, err := repo.LookupBlob(entry.Id)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// 获取文件的历史版本
// 	history, err := repo.BlameFile(path, &git.BlameOptions{
// 		NewestCommit: commit.Id(),
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// 找到和上个版本相比变动的内容
// 	var startLine uint16
// 	for i := 0; i < history.HunkCount(); i++ {
// 		hunk, err := history.HunkByIndex(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		if hunk.FinalCommitId.Equal(commit.Id()) {
// 			// 找到最新版本中的 hunk
// 			startLine = hunk.FinalStartLineNumber
// 		} else if !hunk.OrigCommitId.Equal(commit.Parent(0).Id()) {
// 			// 忽略非上个版本和最新版本之间的变动
// 			continue
// 		} else {
// 			// 输出上个版本和最新版本之间的变动内容
// 			lines := blob.Contents()[startLine : startLine+hunk.LinesInHunk]
// 			for j := 0; j < len(lines); j++ {

// 				if hunk.FinalCommitId.Equal() == git.BlamerNoCommitID {
// 					fmt.Printf("+%s\n", string(lines[j]))
// 				} else if hunk.Lines[j].OrigLineOrigin == git.BlameNoLine {
// 					fmt.Printf("-%s\n", string(lines[j]))
// 				}
// 			}
// 			startLine += hunk.LinesInHunk
// 		}
// 	}
// }
