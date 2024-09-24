# Agora Chat CLI

`agchat` is a command-line interface (CLI) tool for managing Agora chat applications.

- [Installation](#installation)
- [Usage](#Usage)

## Requirements
- Supported operating system (Windows, Linux or macOS).
- App ID, App Certificate, and Base URL for Chat (you can find them in the Agora Console)

## Installation
### Manual
You can grab the latest version of `agchat` from [GitHub releases](https://github.com/ycj3/agora-chat-cli/releases).

### Package Manager (macOS)
You can install `agchat` using [Homebrew](https://brew.sh/).

```
brew tap ycj3/agchat
brew install agchat
```

To upgrade, run:

```
brew upgrade agchat
```

## Usage
* Application Management
```
Manage chat applications

Usage:
  agchat app [command]

Available Commands:
  create      Create an chat application using details from the Agora console
  list        list all chat applications
  remove      Remove one or more applications
  use         Mark an application as the active one

Flags:
  -h, --help   help for app

Use "agchat app [command] --help" for more information about a command.
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

Flags:
  -a, --app            generate a new app token for use in RESTful APIs
  -h, --help           help for token
  -p, --parse string   parse an agora token
  -u, --user string    generate a new user token for use in SDK APIs
```

* Push Notification testing
```
Test whether the push notification credentials and notification services work properly

Usage:
  agchat push test [flags]

Examples:
# Send a test push notification to a specific user
$ agchat push test --user <user-id>

Flags:
  -h, --help             help for test
  -m, --message string   JSON string for the push message (default "{\"title\": \"Admin sent you a message\", \"content\": \"For push notification testing\", \"sub_title\": \"Test message is sent\"}")
  -u, --user string      the user ID of the target user
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
