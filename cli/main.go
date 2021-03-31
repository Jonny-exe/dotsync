package main

import (
	"bufio"
	"errors"
	"flag"
	"github.com/Jonny-exe/dotsync-cli/conf"
	"github.com/Jonny-exe/dotsync-cli/gh"
	"github.com/Jonny-exe/dotsync-cli/git"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func executeCommand(command string, argument string) (string, error) {
	cmd := exec.Command(command, argument)
	stdout, err := cmd.Output()
	stdoutString := string(stdout)
	if err != nil && err.Error() != "exit status 1" {
		// log.Error("Error running command: " + command)
		return "", errors.New("An errror has ocurred while trying: " + command + " " + argument)
	}
	log.Println("Stdout: ", stdoutString)
	return stdoutString, nil
}

func initialize() {
	_, err := executeCommand("mkdir", "/home/a/dotfiles")
	if err != nil {
		log.Warn("Warn: ", err)
	} else {
		log.Info("dotfile folder has been created")
	}

	err = gh.CreateRepo()
	if err != nil {
		log.Error(err)
	}

	err = git.Clone()
	if err != nil {
		log.Error("Error cloning repo: ")
		log.Error(err)
	}

}

func syncDotfiles() {
	lines, err := linebyLineScan("/home/a/test.txt")
	if err != nil {
		log.Error(err)
		return
	}

	fileExceptions := ""
	for _, element := range lines {
		fileExceptions += " --exclude " + element
	}

	err = removeSyncedFiles()
	if err != nil {
		log.Error(err)
	}

	err = syncCommand(fileExceptions)
	if err != nil {
		log.Error(err)
	}

}

func syncCommand(fileExceptions string) error {
	command := "cd /home/a; rsync -rva /home/a/.??* " + fileExceptions + " $HOME/dotfiles-copy/"
	log.Info(command)
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()

	if err != nil {
		error := errors.New("An error has ocurred while trying to sync files")
		return error
	}
	log.Info("Files have been successfully synced")
	return nil
}

func removeSyncedFiles() error {
	cmd := exec.Command("bash", "-c", "rm -rf /home/a/dotfiles-copy")
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

	// flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
	flag.BoolVar(&init, "init", false, "Init dotsync")
	flag.BoolVar(&sync, "sync", false, "Sync dotfiles")

	flag.Parse()
	if init == true {
		initialize()
	}

	if sync == true {
		syncDotfiles()
	}

	// gh.Test()
	git.Test()
	log.Info("Hello, this is dotsync")
	log.Info("Your configuration is:")
	log.Info(conf.Conf)
}
