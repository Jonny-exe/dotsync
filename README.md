#DOTSYNC

This is a tool designed to sync all your important dotfiles in your $HOME.
It syncs your files by creating a private repo and pushing and pulling each 30 mins and every time you start or turn off your computer

You can choose between the cli or the gui.
Keep in mind that the gui needs the cli.

# Technologies used

This doesn't mean you need any of these technologies installed if you want to use it.

## Cli

 * Golang

## Gui

 * Electron
 * React

# Config files

There are two config files, both are located in $HOME/:
 - .dotsyncignore
 - .dotsync.yaml

## .dotsyncignore

Works exactly like a `.gitignore`. Simply list all the dirs you'd like to ignore. For example.

```
#.dotsyncignore

.ssh
.npm

```

## .dotsync.yaml

This is a yaml config with 2 fields. `gh-username` and `gh-access-token`. An example config would be:


```

- gh-username: Jonny-exe
- gh-access-token: youaccesstoken


```

For the access token you will have to create one. You can create one in https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token


