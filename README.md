# JSDock

## Setup & Install
To install the application in your linux system, do the following.

```
curl https://raw.githubusercontent.com/tellaw/jsdock/master/bin/jsdock -o jsdock
chmod +x jsdock
mv jsdock /usr/bin/jsdock
mkdir ~/jsdock
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
profile parameter is optionnal too, you can use the attach option to remember a value

#### Start
Start action is used to run the development server.

Start command is the default command. It is optionnal.

```
> jsdock start <profile>
```

#### Stop / Down
Stop action is used to stop & remove the dev server.

Dev server is always removed, in order to be restart in a fully clean context next time.
```
> jsdock stop <profile>
```

#### Attach
Attach, set the default profile for this directory. It creates a .jsdock file containing the profile name.
To detach a directory, you can simply remove this file.
```
> jsdock attach
```

#### Connect
Up to now, the connect action dumps the command line to use tu connect to the container.
```
> jsdock connect <profile>
```

### Profile Name
Profile name is the filename of the profile located in your <home>/jsdock/ directory.
the application will try to understand the name of the profile your are trying to use.

By default, the appication will look for :
* A command line parameter.
* An attached profile to the directory
* Ask you for a profile.

### Source Path
This is the source directory to inject inside the container. If not specified, the current directory will be used.

In the source directory you can attach a profile. This then makes possible to remove it from any command.

Sources path is injected using the following option in JSON Profile
```
"sources" : "/var/www/html"
```
Take a look at the configuration section

## Configuration

### Profile
Profiles must be stored in <home>/jsdock directory
```
~/jsdock/
```
	
Sample configuration file
```
{
	"alias" : "phpdev",
	"image" : "tellaw-php-74-apache",
	"sources" : "/var/www/html",
	"ports" : [
        {
            "host" : "80",
            "container" : "80"
        },
		{
            "host" : "82",
            "container" : "82"
        }
    ],
	"env" : {
		"mykey":"myvalue",
		"mykey2":"myvalue2"
	},
	"volumes" : [
        {
            "host" : "/tmp",
            "container" : "/tmp"
        }
    ]
	
}
```

## Commands

### Starting Server using sources
Imagine that the sources of your project are located in '/home/me/sources'.

You can attach the profile of the server before starting the server :
```
> cd /home/me/sources
> jsdock attach
```
A prompt will ask you to choose the correct profile.

Now, to start and stop, it is very easy
```
> cd /home/me/sources
> jsdock
```

All parameters are optionnal. The application will try to resolve them.
The default action is 'start'. Starting an already running server will stop it, and restart it.

So, the following commands will do the same :
```
> jsdock start
> jsdock
> jsdock /home/me/sources
> jsdock /home/me/sources start
> jsdock start myprofile
```

The sources of the application will be mounted inside the container.

### Starting a server without sources

Well, really easy. Remove the section 'sources' from the profile.

Then just start it from anywhere using the format :

```
> jsdock <name_of_profile>
```
