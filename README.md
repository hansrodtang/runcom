# Runcom

[![Build](https://img.shields.io/travis/hansrodtang/runcom.svg?style=flat)](https://travis-ci.org/hansrodtang/runcom) [![Coverage](https://img.shields.io/coveralls/hansrodtang/runcom.svg?style=flat)](https://coveralls.io/r/hansrodtang/runcom) [![Tip](https://img.shields.io/gratipay/hansrodtang.svg?style=flat)](https://gratipay.com/hansrodtang/)
[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](http://choosealicense.com/licenses/mit/)
[![Documentation](http://img.shields.io/badge/documentation-wiki-blue.svg?style=flat)](https://github.com/hansrodtang/runcom/wiki)
[![Chat](https://img.shields.io/badge/gitter-join%20chat%20%E2%86%92-brightgreen.svg?style=flat)](https://gitter.im/hansrodtang/runcom)



A [Dotfiles](http://dotfiles.github.io/) manager written in [Go](http://golang.org).

__:heavy_exclamation_mark: Still in heavy development. Not everything in this document is accurate. Please wait for the 1.0 release :heavy_exclamation_mark:__

## Installation

Prebuilt executables are available [here](https://github.com/hansrodtang/runcom/releases).


Install with `go get`
```
  go get github.com/hansrodtang/runcom
```
For those who prefer stable releases, you can also use [gopkg.in](http://gopkg.in):

```
  go get gopkg.in/hansrodtang/runcom.v0
```
You can also download a prebuilt executable
```
  curl -O https://github.com/hansrodtang/runcom/releases
```

## Usage
Here are some example commands.

```sh
# Download from GitHub
runcom get hansrodtang/dotfiles
# Pull changes from GitHub
runcom pull hansrodtang/dotfiles
# Backup local configuration
runcom backup
# Upload to GitHub
runcom push hansrodtang/dotfiles
# Restore from configuration
runcom restore
```

## Plugins

- [List of available plugins and documentation](https://github.com/hansrodtang/runcom/wiki/Plugins)

Fork this repository and add your own today! Pull requests are welcomed!

## Backends

Runcom supports storing your configuration in multiple different locations/services.

- Git (additional integration with GitHub)
- Dropbox
- Compressed files

More coming.

## Naming

The name comes from [Run Commands](https://en.wikipedia.org/wiki/Run_commands).

## License

This software is licensed under the [MIT license](LICENSE.md).
