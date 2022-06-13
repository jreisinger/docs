package main

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
)

func gitClone(localPath, remoteUrl string) error {
	if _, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL: remoteUrl,
	}); err != nil {
		return fmt.Errorf("cloning git repo: %v", err)
	}
	return nil
}

func gitPull(localPath string) error {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(localPath)
	if err != nil {
		return fmt.Errorf("opening git repo: %v", err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("getting git repo worktree: %v", err)
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return fmt.Errorf("pulling git repo: %v", err)
	}

	return nil
}
