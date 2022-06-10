package main

import (
	git "gopkg.in/src-d/go-git.v4"
)

func gitClone(remoteUrl string, localPath string) error {
	_, err := git.PlainClone(localPath, false, &git.CloneOptions{
		URL: remoteUrl,
	})
	return err
}

func gitPull(localPath string) error {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(localPath)
	if err != nil {
		return err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}

	return nil
}
