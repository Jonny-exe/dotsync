package git

import (
	"github.com/Jonny-exe/dotsync-cli/conf"
	. "github.com/WAY29/icecream-go/icecream"
	"github.com/go-git/go-git/v5"
	"os"
)

func main() {
	Ic("")
}

// Clone clones the github repo
func Clone() error {
	// Info("git clone https://github.com/go-git/go-git")

	_, err := git.PlainClone("/tmp/dotsync", false, &git.CloneOptions{
		URL:      "https://github.com/" + conf.Conf.GhUsername + "/dotsync",
		Progress: os.Stdout,
	})

	return err
}

//Test ..
func Test() {
	Ic("test")
}
