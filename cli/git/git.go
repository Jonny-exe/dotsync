package git

import (
	// "errors"
	//"io/ioutil"
	"github.com/Jonny-exe/dotsync-cli/conf"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var auth = &http.TokenAuth{Token: conf.Conf.AccessToken}
var repo *git.Repository
var worktree *git.Worktree

func init() {
	// You have to assign it first so you don't create a local repo variable
	var err error
	repo, err = Clone()
	if err != nil {
		// Not logging because normally the repo already exists
		repo, err = open()
		Ic("Repo")
		Ic(repo)
		if err != nil {
			log.Error("Error opening:")
			log.Error(err)
			return
		}
	}
	worktree, err = repo.Worktree()
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Git status:")
		log.Info(worktree.Status())
	}
}

func open() (*git.Repository, error) {
	repo, err := git.PlainOpen("/home/a/dotsync-files")
	return repo, err
}

// Clone clones the github repo
func Clone() (*git.Repository, error) {
	// Info("git clone https://github.com/go-git/go-git")

	repo, err := git.PlainClone("/home/a/dotsync-files", false, &git.CloneOptions{
		URL:      "git@github.com:" + conf.Conf.GhUsername + "/dotsync-files.git",
		Progress: os.Stdout,
		Auth:     auth,
	})

	return repo, err
}

// Update updates the repo by pulling, commiting and pushing
func Update() error {
	err := pull()

	err = commit()
	Ic("commit")
	if err != nil {
		return err
	}

	err = push()
	if err != nil {
		return err
	}
	return nil
}

func push() error {

	err := repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	return err
}

func commit() error {
	dt := time.Now()
	worktree.AddGlob("") // "" is everything. Ignored should go in .gitignore
	commitOptions := &git.CommitOptions{All: true}
	_, err := worktree.Commit("Update-"+string(dt.String()), commitOptions)
	if err != nil {
		return err
	}
	return nil
}

func pull() error {
	err := worktree.Pull(&git.PullOptions{})
	return err
}
