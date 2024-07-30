## agchat completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(agchat completion zsh)

To load completions for every new session, execute once:

#### Linux:

	agchat completion zsh > "${fpath[1]}/_agchat"

#### macOS:

	agchat completion zsh > $(brew --prefix)/share/zsh/site-functions/_agchat

You will need to start a new shell for this setup to take effect.


```
agchat completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [agchat completion](agchat_completion.md)	 - Generate the autocompletion script for the specified shell

