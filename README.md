# `cmd`

Simple command line interface template for creating internal CLI's 💻

## Structure

```text
├── cmd
│   └── {cmd_name}
│       └── main.go
├── {cmd_name}.go
├── {cmd_sub_name}.go
├── ...
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md
```

The layout consists on the `pkg` directory which has all the dependencies that I
need for creating custom CLI's and the `cmd` directory which just executes the
`cmd` package which is all the _{cmd_name}.go_, _{cmd_sub_name}.go_ files...

The _config_ pkg works as a _getter_ and _setter_ of json, where you can
`config.Query()` and `config.Set()` on a specific config path.

```go
conf := c.Conf{
  Id:   "{cmd_name}",
  Dir:  "{configs_path}",
  File: "config.json",
}
```

Check <https://github.com/thedevsaddam/gojsonq>

### How do I use this?

1. Clone, fork it as a template
1. Do `make init` and add your command name
1. go mod download

### How do I have shell completion?

Put this snippet on your `.bashrc` or any file that is integrated to the
`.bashrc` (sourced in some way).

Note: change `your_cmd` to the name of your executable.

```bash
#!/usr/bin/env bash

if [[ -x "$(command -v your_cmd)" ]]; then
  _cli_bash_autocomplete() {
    if [[ "${COMP_WORDS[0]}" != "source" ]]; then
      local cur opts
      COMPREPLY=()
      cur="${COMP_WORDS[COMP_CWORD]}"
      if [[ "$cur" == "-"* ]]; then
        opts=$("${COMP_WORDS[@]:0:$COMP_CWORD}" "${cur}" --generate-bash-completion)
      else
        opts=$("${COMP_WORDS[@]:0:$COMP_CWORD}" --generate-bash-completion)
      fi
      COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
      return 0
    fi
  }

  complete -o nospace -F _cli_bash_autocomplete your_cmd
fi
```
