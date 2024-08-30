# Agora Chat CLI

`agchat` is a command-line interface (CLI) tool for managing Agora chat applications.

- [Installation](#installation)
- [Usage](#Usage)

## Installation
### Manual
You can grab the latest version of ipatool from [GitHub releases](https://github.com/ycj3/agora-chat-cli/releases).

### Package Manager (macOS)
You can install `agchat` using [Homebrew](https://brew.sh/).

```
brew tap ycj3/agora-chat-cli https://github.com/ycj3/agora-chat-cli
brew install agchat
```

## Usage
* Apps Management
```
Manage all chat apps

Usage:
  agchat apps [flags]

Flags:
  -c, --create   register an existing chat application using details from the Agora console
  -h, --help     help for apps
  -l, --list     list all chat applications
  -r, --remove   remove one or more application
      --use      set an active application for your working directory
```

* Token Management
```
Generate and parse agora tokens

Usage:
  agchat token [flags]

Examples:
# Generate token for a specific user
$ agchat token --user <user-id>

# Generate application token
$ agchat token --app

# Parse an agora token
$ agchat token --parse <user-token>

	#Service type
		ServiceTypeRtc       = 1
		ServiceTypeRtm       = 2
		ServiceTypeFpa       = 4
		ServiceTypeChat      = 5
		ServiceTypeEducation = 7

	#Rtc
		PrivilegeJoinChannel        = 1
		PrivilegePublishAudioStream = 2
		PrivilegePublishVideoStream = 3
		PrivilegePublishDataStream  = 4

	#Rtm
	#Fpa
		PrivilegeLogin = 1

	#Chat
		PrivilegeChatUser = 1
		PrivilegeChatApp  = 2

	#Education
		PrivilegeEducationRoomUser = 1
		PrivilegeEducationUser     = 2
		PrivilegeEducationApp      = 3


Flags:
  -a, --app            generate a new app token for use in RESTful APIs
  -h, --help           help for token
  -p, --parse string   parse an agora token
  -u, --user string    generate a new user token for use in SDK APIs

Global Flags:
  -v, --verbose   enable verbose output
```

* Push Device Management
```
Manage device information bound to a user

Usage:
  agchat device [command]

Available Commands:
  add         Add a new device
  list        List all devices
  remove      Remove an existing device

Flags:
  -h, --help          help for device
  -u, --user string   the user ID of the target user

Use "agchat device [command] --help" for more information about a command.
```

* Push Management
```
Commands to manage push notifications.

Usage:
  agchat push [command]

Available Commands:
  test        Test push notification

Flags:
  -h, --help   help for push

Use "agchat push [command] --help" for more information about a command.
```


**For more detailed documentation, please refer to [here](https://github.com/ycj3/agora-chat-cli/blob/main/docs/agchat.md).**

## Compiling
The tool can be compiled using the Go toolchain.
```
go build -o agchat
```
Unit tests can be executed with the following commands.
```
go generate github.com/ycj3/agora-chat-cli/...
go test -v github.com/ycj3/agora-chat-cli/...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
