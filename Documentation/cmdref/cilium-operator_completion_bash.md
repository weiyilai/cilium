<!-- This file was autogenerated via cilium-operator cmdref, do not edit manually-->

## cilium-operator completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cilium-operator completion bash)

To load completions for every new session, execute once:

#### Linux:

	cilium-operator completion bash > /etc/bash_completion.d/cilium-operator

#### macOS:

	cilium-operator completion bash > $(brew --prefix)/etc/bash_completion.d/cilium-operator

You will need to start a new shell for this setup to take effect.


```
cilium-operator completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [cilium-operator completion](cilium-operator_completion.md)	 - Generate the autocompletion script for the specified shell

