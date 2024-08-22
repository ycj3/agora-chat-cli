## agchat completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(agchat completion bash)

To load completions for every new session, execute once:

#### Linux:

	agchat completion bash > /etc/bash_completion.d/agchat

#### macOS:

	agchat completion bash > $(brew --prefix)/etc/bash_completion.d/agchat

You will need to start a new shell for this setup to take effect.


```
agchat completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat completion](agchat_completion.md)	 - Generate the autocompletion script for the specified shell

