package main

import (
	"log"

	git "gopkg.in/src-d/go-git.v4"
)

func gitClone(repoURL string, repoPath string) {
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func gitPull(repoPath string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("gitPull: %v\n", err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		log.Fatalf("gitPull: %v\n", err)
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Printf("gitPull: %v\n", err)
	}
}
