# Go Modules clarified

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/Helcaraxan/gomod/workflows/Premerge/badge.svg)](https://github.com/Helcaraxan/gomod/actions?query=workflow%3APremerge)
[![Maintainability](https://api.codeclimate.com/v1/badges/42f5920cf5c46650945b/maintainability)](https://codeclimate.com/github/Helcaraxan/gomod/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/42f5920cf5c46650945b/test_coverage)](https://codeclimate.com/github/Helcaraxan/gomod/test_coverage)

`gomod` is a tool that helps Go project maintainers to understand their project's dependencies and
it can provide useful information to developers "modularising" non-module projects. It helps you by
visualising your dependency graph and, even more, analyse it for your profit. It will help you
answer typical questions such as:

- How can I visualise the network of my dependencies?
- How old are the versions of my dependencies that I depend on?
- Are different dependencies of my project using potentially conflicting forks of the same module?
- What dependency chains lead to `github.com/foo/bar` and what constraints do they put on versions?
- Why is dependency `github.com/foo/bar` used at version `1.12.0` and not at version `1.5.0` as I
  want it to be?

_Release-notes for each version can be found [here](./RELEASE_NOTES.md)_

## Table of Contents

- [Detailed features](#detailed-features)
  - [Dependency analysis commands](#dependency-analysis-commands)
    - [`gomod graph`](#gomod-graph)
    - [`gomod reveal`](#gomod-reveal)
    - [`gomod analyse`](#gomod-analyse)
  - [Command-line use](#command-line-use)
    - [Bash](#bash)
    - [Powershell](#powershell)
    - [ZSH](#zsh)
- [Example output](#example-output)
  - [Shared dependencies](#shared-dependencies)
  - [Dependency chains](#dependency-chains)
  - [Hidden `replace`'s](#hidden-replaces)
  - [Dependency statistics](#dependency-statistics)

## Detailed features

### Dependency analysis commands

#### `gomod graph`

Create a graphical representations of your dependency graph with the possibility to filter out
noise, add annotations and focus on the pieces of the graph that are of interest to you. You can for
example:

- Only show dependencies that are required by more than one module.
- Only show the dependency chains that lead to one or more specified modules.
- Annotate dependencies with the versions in which they are used and the versions constraint
  imposed by each edge of the graph.

This functionality requires the [`dot`](https://www.graphviz.org/) tool which you will need to
install separately. You can produce images in GIF, JPG, PDF, PNG and PS format.

#### `gomod reveal`

Show all the places at which your (indirect) module dependencies use `replace` statements which you
might need to account for in your own `go.mod` in order to build your project.

#### `gomod analyse`

Produce a short statistical report of what is going on with your dependencies. The report includes
things like (in)direct dependency counts, mean and max dependency ages, dependency age distribution,
and more.

**NB**: This command can also be invoked as `gomod analyze` for those who intuitively use American
spelling.

### Command-line use

The sub-commands of `gomod completion` provide you with the benefit of shell auto-completion making
it easier to cycle through the available sub-commands and flags.

#### Bash

Besides completion for static elements such as commands and flags, the auto-complete functionality
for Bash also provides some context-specific completion such as the dependency names of your current
project.

If you only want completion to be added to your current shell you can simply run

```bash
source <(gomod completion bash)
```

If you want to have it added by default to each shell instance that you start you can add the
following to your `.bashrc`, `.profile` (Linux) or `.bash_profile` (MacOS / OSX)

```bash
# Provide Bash auto-completion for the 'gomod' tool if it's in the PATH.
if [[ -n "$(which gomod)" ]]; then
   source <(gomod completion bash)
fi
```

#### Powershell

In order to auto-completion by default to each shell instance that you start you can add the
following to your [PowerShell profile](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-6)

```powershell
if (Get-Command "gomod" -errorAction SilentlyContinue) {
   $tmp = New-TemporaryFile
   gomod completion ps | Out-File $tmp.FullName
   . $tmp.FullName
   Remove-Item $tmp.FullName
}
```

#### ZSH

If you only want completion to be added to your current shell you can simply run

```zsh
source <(gomod completion zsh)
```

If you want to have it added by default to each shell instance that you start you can add the
following to your `.bashrc`, `.profile` (Linux) or `.bash_profile` (MacOS / OSX)

```zsh
# Provide ZSH auto-completion for the 'gomod' tool if it's in the PATH.
if [[ -n "$(which gomod)" ]]; then
   source <(gomod completion zsh)
fi
```

## Example output

### Full dependency graph

The entire dependency graph for the `gomod` codebase, using clustering to reduce the image size.
![Full dependency graph](./images/full.jpg)

### Shared dependencies

Graph with only the shared dependencies for the `gomod` project.
![Shared dependencies graph](./images/shared-dependencies.jpg)

### Dependency chains

Specific zoom on the dependency chains leading to the `github.com/stretchr/testify` and
`golang.org/x/sys` modules with version annotations.
![Annotated dependency chains for `github.com/stretchr/testify`](./images/dependency-chains.jpg)

### Hidden `replace`'s

Revealing hidden replaces in the [Matterbridge](https://github.com/42wim/matterbridge) dependency
graph. One is accounted for, the other is not. This can be the cause of unexpected errors when
building the project.

```text
 -> gomod reveal
'github.com/Rhymen/go-whatsapp' is replaced:
   maunium.net/go/mautrix-whatsapp -> github.com/tulir/go-whatsapp @ v0.0.2-0.20190528182350-fde573a2a73b

'gopkg.in/russross/blackfriday.v2' is replaced:
 ✓ maunium.net/go/mautrix            -> github.com/russross/blackfriday/v2 @ v2.0.1
 ✓ maunium.net/go/mautrix-appservice -> github.com/russross/blackfriday/v2 @ v2.0.1
 ✓ maunium.net/go/mautrix-whatsapp   -> github.com/russross/blackfriday/v2 @ v2.0.1

[✓] Match with a top-level replace in 'github.com/42wim/matterbridge'
```

### Dependency statistics

Statistical analysis of the `gomod` dependency graph.

```text
 -> gomod analyse
-- Analysis for 'github.com/Helcaraxan/gomod' --
Dependency counts:
- Direct dependencies:   10
- Indirect dependencies: 28

Age statistics:
- Mean age of dependencies: 15 month(s) 18 day(s)
- Maximum dependency age:   58 month(s) 17 day(s)
- Age distribution per month:

  18.42 % |          #
          |          #
          |          #
          |          #
          |    _   _ #   _   _
          |    #   # #   #   #
          |    #   # #   #   #
          |  _ # _ # # _ #   #     _
          |  # # # # # # #   #     #
          |# # # # # # # # # # #   #           #     #         #     #
   0.00 % |___________________________________________________________
           0                                                        60

Update backlog statistics:
- Number of dependencies with an update:  12 (of which 0 are direct)
- Mean update backlog of dependencies:    10 month(s) 4 day(s)
- Maximum update backlog of dependencies: 18 month(s) 4 day(s)
- Update backlog distribution per month:

  25.00 % |            #
          |            #
          |            #
          |            #
          |            #
          |            #
          |            #
          |  # #       # #     # #     #   # # #
          |  # #       # #     # #     #   # # #
          |  # #       # #     # #     #   # # #
   0.00 % |_____________________________________
           0                                  19

Reverse dependency statistics:
- Mean number of reverse dependencies:    1.42
- Maximum number of reverse dependencies: 4
- Reverse dependency count distribution:

  76.32 % |  #
          |  #
          |  #
          |  #
          |  #
          |  #
          |  #
          |  #
          |  #
          |  # # # _
   0.00 % |_________
           0       5
```
