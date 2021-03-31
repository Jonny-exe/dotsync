package gh

import (
	// "github.com/google/go-github/github"
	"context"
	"errors"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/google/go-github/v34/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func init() {
}

//Test ..
func Test() {
	Ic("test")
}

// CreateRepo creates a github repo for dotsync
func CreateRepo() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "3529786de91f272d2e56686f0b84a4c080c8b790"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:    github.String("dotsync-files"),
		Private: github.Bool(true),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return errors.New("dotsync-files repository already exists")
	}
	log.Info("Repo has been created successfully")
	return nil
}
