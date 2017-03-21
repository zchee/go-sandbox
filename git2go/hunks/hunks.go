package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/libgit2/git2go"
)

/*
github.com/libgit2/git2go
commit 80cf533fe4e48ddfab3015d9570f2833951c1dea
Author: David Pierce <david.pierce@moz.com>
Date:   Sat Sep 26 15:37:48 2015 -0700

    Config#LookupString uses git_buf to load value

 config.go      |  8 +++++---
 config_test.go | 58 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 2 files changed, 63 insertions(+), 3 deletions(-)
 commit SHA-1 checksum
 commitID := `80cf533fe4e48ddfab3015d9570f2833951c1dea`
 commitOid, err := git.NewOid(commitID)
*/

func main() {
	repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/libgit2/git2go")
	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Fatal(err)
	}
	commitOid, err := gitRepo.Head()
	if err != nil {
		log.Fatal(err)
	}
	blob, _ := gitRepo.LookupBlob(commitOid.Target())
	log.Println(blob)
	// commit, err := gitRepo.LookupCommit(commitOid)
	commit, err := gitRepo.LookupCommit(commitOid.Target())
	if err != nil {
		log.Fatal(err)
	}
	commitTree, err := commit.Tree()
	if err != nil {
		log.Fatal(err)
	}
	options, err := git.DefaultDiffOptions()
	if err != nil {
		log.Fatal(err)
	}
	options.IdAbbrev = 40
	options.InterhunkLines = 0
	options.Flags = git.DiffIncludeUntracked
	var parentTree *git.Tree
	if commit.ParentCount() > 0 {
		parentTree, err = commit.Parent(0).Tree()
		if err != nil {
			log.Fatal(err)
		}
	}
	gitDiff, err := gitRepo.DiffTreeToTree(parentTree, commitTree, &options)
	if err != nil {
		log.Fatal(err)
	}

	findOpts, err := git.DefaultDiffFindOptions()
	findOpts.Flags = git.DiffFindBreakRewrites
	err = gitDiff.FindSimilar(&findOpts)

	// Show all file patch diffs in a commit.
	files := make([]string, 0)
	hunks := make([]git.DiffHunk, 0)
	lines := make([]git.DiffLine, 0)
	patches := make([]string, 0)
	err = gitDiff.ForEach(func(file git.DiffDelta, progress float64) (git.DiffForEachHunkCallback, error) {
		patch, err := gitDiff.Patch(len(patches))
		if err != nil {
			return nil, err
		}
		defer patch.Free()
		patchStr, err := patch.String()
		if err != nil {
			return nil, err
		}
		patches = append(patches, patchStr)

		files = append(files, file.OldFile.Path)
		return func(hunk git.DiffHunk) (git.DiffForEachLineCallback, error) {
			hunks = append(hunks, hunk)
			return func(line git.DiffLine) error {
				lines = append(lines, line)
				return nil
			}, nil
		}, nil
	}, git.DiffDetailLines)

	log.Println("files: ", files, "\n")
	log.Println("hunks: ", hunks, "\n")
	log.Println("lines: ", lines, "\n")
	log.Println("patches: ", patches, "\n")
}
