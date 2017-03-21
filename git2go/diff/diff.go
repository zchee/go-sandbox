package main

import (
	"log"
	"os"
	"path/filepath"

	git "github.com/libgit2/git2go"
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
*/

func main() {
	// After go get -v github.com/libgit2/git2go
	// path to git2go repository in your $GOPATH
	repoPath := filepath.Join(os.Getenv("GOPATH"), "src/github.com/libgit2/git2go")
	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		log.Fatal(err)
	}
	// commit SHA-1 checksum
	commitID := `80cf533fe4e48ddfab3015d9570f2833951c1dea`
	commitOid, err := git.NewOid(commitID)
	if err != nil {
		log.Fatal(err)
	}
	commit, err := gitRepo.LookupCommit(commitOid)
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

	// Show all file patch diffs in a commit.
	numDeltas, err := gitDiff.NumDeltas()
	if err != nil {
		log.Fatal(err)
	}
	for d := 0; d < numDeltas; d++ {
		patch, err := gitDiff.Patch(d)
		if err != nil {
			log.Fatal(err)
		}
		patchString, err := patch.String()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\n%s", patchString)
		patch.Free()
	}
}
