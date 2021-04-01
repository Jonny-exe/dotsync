package gh

import (
	// "github.com/google/go-github/github"
	"context"
	"errors"
	"github.com/Jonny-exe/dotsync-cli/conf"
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
		&oauth2.Token{AccessToken: conf.Conf.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:     github.String("dotsync-files"),
		Private:  github.Bool(true),
		AutoInit: github.Bool(true),
	}
	_, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return errors.New("Error creating repo: dotsync-files repository already exists")
	}
	log.Info("Repo has been created successfully")
	return nil
}
