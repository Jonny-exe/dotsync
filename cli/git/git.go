package git

import (
	// "errors"
	//"io/ioutil"
	"github.com/Jonny-exe/dotsync-cli/conf"
	// . "github.com/WAY29/icecream-go/icecream"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var auth = &http.TokenAuth{Token: conf.Conf.AccessToken}

// Open opens a local git repo
func Open() (*git.Repository, error) {
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

func pullAndPush(worktree *git.Worktree) error {
	err := worktree.Pull(&git.PullOptions{})
	return err
}

func push(repo *git.Repository) error {

	err := repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
	})
	return err
}

func commit(worktree *git.Worktree) error {
	dt := time.Now()

	log.Info(worktree.Status())

	_, err := worktree.Commit("Update-"+string(dt.String()), &git.CommitOptions{All: true})
	if err != nil {
		return err
	}
	return nil
}

//Test ..
func Test() {
	repo, err := Clone()
	if err != nil {
		log.Error("Error cloning:")
		log.Error(err)
		repo, err = Open()
		if err != nil {
			log.Error("Error opening:")
			log.Error(err)
			return
		}
	}
	worktree, err := repo.Worktree()
	if err != nil {
		log.Error(err)
	}
	_ = pullAndPush(worktree)
	err = commit(worktree)
	if err != nil {
		log.Error(err)
	}

	err = push(repo)
	if err != nil {
		log.Error(err)
	}

}

//func publicKeyFile(file string) ssh.AuthMethod {
//buffer, err := ioutil.ReadFile(file)
//if err != nil {
//return nil
//}

//key, err := ssh.ParsePrivateKey(buffer)
//if err != nil {
//return nil
//}
//return ssh.PublicKeys(key)
//}
