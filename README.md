# JSDock

## Setup & Install
To install the application in your linux system, do the following.

```
curl https://raw.githubusercontent.com/tellaw/jsdock/master/bin/jsdock -o jsdock
chmod +x jsdock
mv jsdock /usr/bin/jsdock
```

## What is JSDock
JSDock is a tool that makes simple to use docker for developers environments.

Instead of creating a docker compose configuration for each of your projects, build one profile for your dev server, and start it from your sources to inject them inside.

## Command line usage
JSDock takes 3 parameters :

* Profile Name (optionnal)
* Source Path (optionnal)
* Action (optionnal)

### Actions

Default action is start / stop in the path context.

#### Start
Start action is used to run the development server.

#### Stop / Down
Stop action is used to stop & remove the dev server.
Dev server is always removed, in order to be restart in a fully clean context next time.

#### Attach
Attach, set the default profile for this directory. It creates a .jsdock file containing the profile name.
To detach a directory, you can simply remove this file.

#### Connect
Up to now, the connect action dumps the command line to use tu connect to the container.

### Profile Name
Profile name is the filename of the profile located in your <home>/jsdock/ directory.
the application will try to understand the name of the profile your are trying to use.

By default, the appication will look for :
* A command line parameter.
* An attached profile to the directory
* Ask you for a profile.

### Source Path
This is the source directory to inject inside the container. If not specified, the current directory will be used.

## Configuration
