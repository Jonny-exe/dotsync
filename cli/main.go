package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/Jonny-exe/dotsync-cli/conf"
	"github.com/Jonny-exe/dotsync-cli/gh"
	"github.com/Jonny-exe/dotsync-cli/git"
	. "github.com/WAY29/icecream-go/icecream"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(command string, argument string) (string, error) {
	cmd := exec.Command(command, argument)
	stdout, err := cmd.Output()
	stdoutString := string(stdout)
	if err != nil && err.Error() != "exit status 1" {
		return "", errors.New("An errror has ocurred while trying: " + command + " " + argument)
	}
	log.Println("Stdout: ", stdoutString)
	return stdoutString, nil
}

func initialize() {
	err := gh.CreateRepo()
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Repo created")
	}

	_, err = git.Clone()
	if err != nil {
		log.Error("Error cloning repo: ")
		log.Error(err)
	} else {
		log.Info("Repo cloned")
	}
	err = setUpCronjob()
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Cronjob successfully created")
	}

}

func syncDotfiles() {
	home := os.Getenv("HOME")
	lines, err := linebyLineScan(home + "/.dotsyncignore")
	if err != nil {
		log.Error(err)
		return
	}

	fileExceptions := ""
	for _, element := range lines {
		fileExceptions += " --exclude " + element
	}

	err = syncCommand(fileExceptions)
	if err != nil {
		log.Error(err)
	}

	err = git.Update()
	if err != nil {
		log.Error(err)
	}
}

func setUpCronjob() error {
	cronjobs, err := executeCommand("crontab", "-l")
	if err != nil {
		return err
	}

	if strings.Contains(cronjobs, "0 * * * * dotsync -sync\n") {
		return errors.New("Cronjob already exists")
	}

	file, err := os.Create("/tmp/dotsync-crontab")
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(cronjobs + "# This next line was generated by dotsync. Do not edit it\n0 * * * * dotsync -sync\n")
	if err != nil {
		return err
	}

	writer.Flush()

	_, err = executeCommand("crontab", "/tmp/dotsync-crontab")
	if err != nil {
		return err
	}
	return nil
}

func syncCommand(fileExceptions string) error {
	command := "rsync -rvaL $HOME/.??* " + fileExceptions + " $HOME/dotsync-files/"
	log.Info(command)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()

	if err != nil {
		error := errors.New("An error has ocurred while trying to sync files")
		return error
	}
	log.Info("Files have been successfully copied from $HOME to dotsync-files")

	command = "rsync -rvaL $HOME/dotsync-files/" + fileExceptions + " $HOME/"
	log.Info(command)
	cmd = exec.Command("bash", "-c", command)
	_, err = cmd.Output()

	if err != nil {
		error := errors.New("An error has ocurred while trying to sync files")
		return error
	}
	log.Info("Files have been successfully copied from dotsync-files to $HOME")
	return nil
}

func removeSyncedFiles() error {
	cmd := exec.Command("bash", "-c", "rm -rf $HOME/dotfiles-copy")
	_, err := cmd.Output()

	if err != nil {
		error := errors.New("An error has ocurred while trying to remove synced files")
		return error
	}
	log.Info("Synced files have been successfully removed")
	return nil
}

func linebyLineScan(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return lines, errors.New("An errror has ocurred while trying scanning file")
	}
	return lines, nil
}

func main() {
	// Arguments
	var init bool
	var sync bool
	var config bool

	// flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
	flag.BoolVar(&init, "init", false, "Init dotsync")
	flag.BoolVar(&sync, "sync", false, "Sync dotfiles")
	flag.BoolVar(&config, "config", false, "See your config")

	flag.Parse()
	if init == true {
		initialize()
	} else {
		git.Initialize()
	}

	if sync == true {
		syncDotfiles()
	}

	if config == true {
		Ic(conf.Conf)
	}
}
