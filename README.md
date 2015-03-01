# [![Logo](http://i.imgur.com/IPUdmyu.png)](http://hansrodtang.github.io/runcom/) Runcom

 [![Coverage](https://img.shields.io/coveralls/hansrodtang/runcom.svg?style=flat)](https://coveralls.io/r/hansrodtang/runcom) [![Tip](https://img.shields.io/gratipay/hansrodtang.svg?style=flat)](https://gratipay.com/hansrodtang/)
[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](http://choosealicense.com/licenses/mit/)
[![Documentation](http://img.shields.io/badge/documentation-wiki-blue.svg?style=flat)](https://github.com/hansrodtang/runcom/wiki)
[![Chat](https://img.shields.io/badge/gitter-join%20chat%20%E2%86%92-brightgreen.svg?style=flat)](https://gitter.im/hansrodtang/runcom)

| System  | Status | Documentation |
| --- | :----------: | -------------: |
| OSX | [![Travis Build](https://img.shields.io/travis/hansrodtang/runcom.svg?style=flat)](https://travis-ci.org/hansrodtang/runcom) | Coming soon |
| Linux | [![Wercker Build](https://img.shields.io/wercker/ci/54ed3bb6b05d63312300510a.svg?style=flat)](https://app.wercker.com/project/bykey/803c5515e66c7ea2329cd411d6fa2407) | Coming soon |
| Windows | [![AppVeyor Build](https://img.shields.io/appveyor/ci/hansrodtang/runcom.svg?style=flat)](https://ci.appveyor.com/project/hansrodtang/runcom) | Coming soon |



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

Icon is created by [Arthur Shlain](http://thenounproject.com/ArtZ91/) and is licensed under [Creative Commons – Attribution (CC BY 3.0)](https://creativecommons.org/licenses/by/3.0/us/).
