# Push-Me
This is an automatic commit tool set to create a fresh commit message as your user
on an hourly basis.

The idea behind this tool is that the user will have "tracked" git repos that
are already configured with remote repositories. These repos are then given an
automatic commit at the :00 of every hour. The goal is to reduce the likelihood
that someone forgets to commit their work.

## How it works
Repos are tracked in a configuration file located in the user's home directory.
This file should be named `.push-me-config.yml` and needs to exist prior to the
tool being installed.

### Example configuration file
An example `.push-me-config.yml` may look like this:

```yml
repos:
    - "my/relative/repo/path"
    - "my/other/repo/path"
```

### Under the hood
Using the home dir we can create a list of absolute paths to the git repos
then execute git commands against them.

```
"my/relative/repo/path" -> "/home/<my username>/my/relative/repo/path"
```

### Install script
The installation script handles the creation of the cron task to be picked up
by the program scheduler.

### Git commands
Git commands are run in the following order:
1. `git add .`
1. `git commit -m "auto commit <timestamp>"`
1. `git push`

## Usage
The tool automatically handles execution so long as the installation script
has been handled.

## Installation
To install please clone the project and run the following command:

```bash
git clone https://github.com/mattdood/push-me
make install
```

## Requirements
This tool requires Go to run.

