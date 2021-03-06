# Dotsync

This is a tool designed to sync all your important dotfiles in your $HOME.
It syncs your files by creating a private github repo and pushing and pulling every 30 mins 
and each time you start or turn off your computer

You can choose between the cli or the gui.
Keep in mind that the gui needs the cli.

# Installation

## Cli

Download the binary from the release and add it to your path. After that you can start using it.

## Gui

Download the binary for the cli and the gui from the releases and make sure you add the cli to your path.

# How to use

## Cli

The program is called dotsync-cli. Once you add it to your path you can call it by typing `dotsync-cli` in your terminal. You can call it with certain options.
Before you start using it make sure to create the config files.
 - `-init` to initialize it. This will create a repo and clone it to syncronize your files.
It will also copy all the dotfiles from your $HOME to that repo and it will create a cronjob to run `-sync` every 30 mins.
 - `-sync` to sync your files with the ones that are in the repo, and to update the repo.
 - `-config` to see your current config.

## Gui

Download it from the releases and use it. It is made to be intuitive so I don't think it needs more explanation.


# Config files

There are three config files, both are located in $HOME/:
 - .dotsyncignore
 - .dotsyncinclude
 - .dotsync.yaml

## .dotsyncignore

Works exactly like a `.gitignore`. Simply list all the dirs you'd like to ignore. For example.

```
#.dotsyncignore

.ssh
.npm
```

## .dotsyncinclude

Works exactly like a `.gitignore`. Simply list all the dirs you'd like include which are in excluded directories. For example.

```
#.dotsyncinclude

.ssh/key.pub
.npm/something/hello.js
```

## .dotsync.yaml

This is a yaml config with 2 fields. `gh-username` and `gh-access-token`. An example config would be:


```
#.dotsync.yaml

gh-username: Jonny-exe
gh-access-token: youaccesstoken
```

For the access token you will have to create one. You can create one in https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token
 
# Technologies used

This doesn't mean you need any of these technologies installed if you want to use it.

## Cli

 * Golang

## Gui

 * Electron
 * React

