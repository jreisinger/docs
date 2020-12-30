package util

import (
	"log"
	"os"
	"time"

	git "gopkg.in/src-d/go-git.v4"
)

// GitPuller regularly git pulls a repoURL to repoPath. Can be run as a goroutine.
func GitPuller(repoURL string, repoPath string) {
	for {
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			gitClone(repoURL, repoPath)
		}

		gitPull(repoPath)
		time.Sleep(time.Second * 2)
	}
}

func gitClone(repoURL string, repoPath string) {
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	Check(err)
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
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Printf("gitPull: %v\n", err)
	}
}
