# JSDock

JSDock is a shortname for *JSon Docker*. The goal of this application is make your developement easier. JSDock manage profiles of developement servers and make them very easy to start using any source code directory.
Forget the dedicated docker compose per application or any unreadable command lines.

JSDock wrap development using docker and WSL2 into something very easy.

- [JSDock](#jsdock)
  * [How does it works ?](#how-does-it-works--)
  * [Setup & Install](#setup---install)
  * [Command line usage](#command-line-usage)
    + [Actions](#actions)
      - [Start](#start)
      - [Stop / Down](#stop---down)
      - [Attach](#attach)
      - [Connect](#connect)
      - [List](#list)
      - [Help or ?](#help-or--)
    + [Profile Name](#profile-name)
    + [Source Path](#source-path)
  * [Configuration](#configuration)
    + [Profile](#profile)
      - [Sample JSON configuration file](#sample-json-configuration-file)
      - [Sample YAML configuration file](#sample-yaml-configuration-file)
  * [Commands](#commands)
    + [Starting Server using sources from a directory](#starting-server-using-sources-from-a-directory)
    + [Starting a server without sources](#starting-a-server-without-sources)
  * [Profile Configuration](#profile-configuration)
    + [Example of configuration](#example-of-configuration)
      - [Sample JSON configuration file](#sample-json-configuration-file-1)
      - [Sample YAML configuration file](#sample-yaml-configuration-file-1)
    + [Alias](#alias)
    + [Image](#image)
    + [Sources](#sources)
    + [Ports](#ports)
    + [ENV Variables](#env-variables)
    + [Volumes](#volumes)
    + [Conditionnal volumes](#conditionnal-volumes)
  * [Network](#network)
  * [Some samples for the configuration.](#some-samples-for-the-configuration)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>


## How does it works ?
JSDock keeps the configuration of every server in a profile repository described in JSON or YAML (~/jsdock/). Each profile is linked automatically to a dev network, making communication beetween dockers very easy.
Profiles can be configured using dynamic sources directory. The sources directory can be injected automatically inside the server using you path context or any path given by you, or set as default.

*So, want to start ?*

1. Create a profile in ~/jsdock/ or copy a sample profile -> https://github.com/tellaw/jsdock/tree/master/doc/sample_conf
2. Go to your sources directory and attach it the good profile -> '> jsdock attach'
3. Start server => '> jsdock' 

![Introduction](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-1.gif)

Or just read this documentation...

## Setup & Install
To install the application in your linux system, do the following.

```
curl https://raw.githubusercontent.com/tellaw/jsdock/master/bin/jsdock -o jsdock
chmod +x jsdock
mv jsdock /usr/bin/jsdock
mkdir ~/jsdock
```

## Command line usage
JSDock can takes up to 3 arguments :

* Action (optionnal)
* Profile Name (optionnal)
* Source Path (optionnal)


The order of the arguments is not important, JSDock will try to resolve each of them.

### Actions

Default action is start / stop in the path context.
profile parameter is optionnal too, you can use the attach option to remember a value

#### Start
Start action is used to run the development server.
Start command is the default command. It is optionnal.

```
> jsdock start <profile>
```

![JSDock start already running](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-2.gif)

When JSDock is starting a server, it'll check for an already running server and shut it down, it'll also check conflict to resolve them.

#### Stop / Down
Stop action is used to stop & remove the dev server.

Dev server is always removed, in order to be restart in a fully clean context next time.
```
> jsdock stop <profile>
```

![JSDock stop](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-stop.gif)


#### Attach
Attach, set the default profile for this directory. It creates a .jsdock file containing the profile name.
To detach a directory, you can simply remove this file.
```
> jsdock attach
```

![Introduction](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-1.gif)


#### Connect
Up to now, the connect action dumps the command line to use tu connect to the container.
```
> jsdock connect <profile>
```

![Connect Action](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-connect.gif)

#### List
List the available profiles on the setup. This command will dump every profile, its alias and its image.
It makes easy for you to have an overview of your setup, and possible conflicts. 

```
> jsdock list
```

![List profiles](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-list.gif)

#### Help or ?

Display the help

```
> jsdock help
> jsdock ?
```

![Help](https://github.com/tellaw/jsdock/blob/master/doc/images/jsdock-help.gif)

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
see the details of configuration later in this documentation.

#### Sample JSON configuration file
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

#### Sample YAML configuration file

```
alias: imageAlias
image: tellaw-php-74-apache
sources: "/var/www/html"
ports:
- host: '80'
  container: '80'
- host: '82'
  container: '82'
env:
  mykey: myvalue
  mykey2: myvalue2
volumes:
- host: "/home/tellaw/jsdock/conf/000-sf4-default.conf"
  container: "/etc/apache2/sites-available/000-default.conf"
```

## Commands

### Starting Server using sources from a directory
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

## Profile Configuration

### Example of configuration
#### Sample JSON configuration file
```
File name : ~/jsdock/myprofile.json (lowercase)
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

#### Sample YAML configuration file
```
File name : ~/jsdock/myprofile.yaml (lowercase and yaml as extension)

alias: imageAlias
image: tellaw-php-74-apache
sources: "/var/www/html"
ports:
- host: '80'
  container: '80'
- host: '82'
  container: '82'
env:
  mykey: myvalue
  mykey2: myvalue2
volumes:
- host: "/home/tellaw/jsdock/conf/000-sf4-default.conf"
  container: "/etc/apache2/sites-available/000-default.conf"

```

### Alias

Alias is the name you gave to the profile. This name must be uniq as it'll make JSDock able to understand if the application is already running.

The alias will also be used to communicate beetween containers.

```
JSON : 
{
	"alias" : "phpdev",
...
}

YAML :
alias: imageAlias
```
### Image

This is the name of the image that you want to use.
For my purpose, I do use custom images contaiing specific configurations for local dev. I do share thoses samples in one of my github projects (jsdock-samples).

```
JSON : 
{
	"image" : "php:7.4-apache",
...
}

YAML : 
image: tellaw-php-74-apache
```
Here, we are requesting the server to use an image PHP, version 7.4 with Apache.

### Sources

This is the path where you ant to inject the source directory. The sources will be mounted as a volume to this path.

```
JSON : 
{
	"sources" : "/var/www/html",
...
}

YAML : 
sources: "/var/www/html"
```
This means to JSDock that the sources have to be mounted in this directory iside the container.

### Ports
In this section, you describe the ports required by your server. at startup, JSDock will check any conflict of port with already running container. If any conflict is going to happen, JSDock will try to stop the other containers.

```
JSON : 
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

YAML :

ports:
- host: '80'
  container: '80'
- host: '82'
  container: '82'

```
This means the JSDock should expose ports 80 and 82 to the same ports of the host.

### ENV Variables
ENV Variables makes possible to set env variables inside the container.

```
JSON : 

{
	"env" : {
		"mykey":"myvalue",
		"mykey2":"myvalue2"
	},
...
}

YAML :

env:
  mykey: myvalue
  mykey2: myvalue2

```
Here, we set two env variables, mykey and mykey2.

### Volumes
Here you can set volumes to mount inside the container. Sources are a dynamic volume, which is not configure in this section.

```
JSON : 
{
	"volumes" : [
        {
            "host" : "/home/tellaw/jsdock/conf/000-sf4-default.conf",
            "container" : "/etc/apache2/sites-available/000-default.conf"
        }
    ]
...
}


YAML :

volumes:
- host: "/home/tellaw/jsdock/conf/000-sf4-default.conf"
  container: "/etc/apache2/sites-available/000-default.conf"

```
Here, we are injecting an apache configuration inside the container.

### Conditionnal volumes
Conditionnal volumes will help you to inject inside your container specific configuration depending of the source code. Classic use case is an Apache webserver shared beetween projects using Synfony or Wordpress PHP Application.
In such case, you will require that JSDOck inject the correct configuration of the VHost depending of the sources.

Conditionnal volumes can be triggered by :

1. A file available in a specific location
2. A directory available in a specific location
3. A specific text available inside a text file

```
JSON : 
{
...
	"volumes" : [
        {
            "host" : "/home/tellaw/jsdock/conf/000-sf4-default.conf",
            "container" : "/etc/apache2/sites-available/000-default.conf",
			      "conditions" : {

                "fileExists" : [
                    "subdir/webmyfile.conf",
                    "/home/tellaw/myfile.conf"
                ],
                
                "dirExists" : [
                    "public",
                    "/home/tellaw/mydir"
                ],
                
                "fileContains" : [
                    {   
                        "file":"subdir/myfile.conf" ,
                        "value": "mycontent"
                    },
                    {   
                        "file":"/home/tellaw/subdir/myfile.conf" ,
                        "value": "mycontent"
                    }
                ]

            }
        }
    ]
...
}


YAML :

volumes:
- host: "/home/tellaw/jsdock/conf/000-sf4-default.conf"
  container: "/etc/apache2/sites-available/000-default.conf"
  conditions:
    fileExists:
    - subdir/webmyfile.conf
    - "/home/tellaw/myfile.conf"
    dirExists:
    - public
    - "/home/tellaw/mydir"
    fileContains:
    - file: subdir/myfile.conf
      value: mycontent
    - file: "/home/tellaw/subdir/myfile.conf"
      value: mycontent


```

When specifying multiple value inside one criteria, they are always considered as a logic OR.


## Network

JSDock automatically add a network called jsdock_net to all of the container it managed. This makes possible for you to communicate beetween too servers using the alias you declared in the JSON Profile.

## Some samples for the configuration.
I wanted to share with you the samples of my configurations. This may help you to start quickly.
All files are stored in the following github project :
https://github.com/tellaw/jsdock-samples
