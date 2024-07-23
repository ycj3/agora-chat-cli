# Agora Chat CLI

`agora-chat-cli` is a command-line interface (CLI) tool for managing Agora chat applications.

## Table of Contents

- [Installation](#installation)
- [Commands](#commands)
  - [Apps Management](#apps-management)
  - [Token Management](#token-managemen)
  - [Push Management](#push-managemen)

## Installation

* To install the CLI, clone the repository and build the project:

```sh
$ git clone https://github.com/CarlsonYuan/agora-chat-cli.git
$ cd agora-chat-cli
$ go build -o agchat main.go
```

* Run the command:
```
$ ./agchat --version
agchat version 0.0.1
```

## Commands
### Apps Management

| Command                             | Description                                                           |
|-------------------------------------|-----------------------------------------------------------------------|
| `agchat apps --list`                | List all configured apps                                              |
| `agchat apps --create`              | Create a new chat application                                         |
| `agchat apps --remove <app-id>`     | Remove one or more application                                        |
| `agchat apps --use <app-id>`        | Set an active application for your working directory                  |

### Token Management
| Command                             | Description                                                           |
|-------------------------------------|-----------------------------------------------------------------------|
| `agchat token --user <user-id>`     | Generate a new user token for use in SDK APIs                         |
| `agchat token --parse <token>`      | Parse an Agora token                                                  |

### Push Management
| Command                             | Description                                                           |
|-------------------------------------|-----------------------------------------------------------------------|
| `agchat push test --user <user-id>` | Test push notification |


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
