# gpassmanager [![gpassmanager](https://snapcraft.io/gpassmanager/badge.svg)](https://snapcraft.io/gpassmanager)[![gpassmanager](https://snapcraft.io/gpassmanager/trending.svg?name=0)](https://snapcraft.io/gpassmanager) [![CodeFactor](https://www.codefactor.io/repository/github/mrinjamul/gpassmanager/badge)](https://www.codefactor.io/repository/github/mrinjamul/gpassmanager)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/gpassmanager)

Simple Password Manager Application

A Commandline application for managing your passwords securely.
It's a very secure password manager.

Using gpassmanager, you can also generate secure passwords with given length.

Now, you can import passwords from google accounts (Google Chrome passwords).

First, you need to export the passwords as csv file.

Then, import into gpassmanager by running the following,

```sh
gpassmanager import --csv [csv filename]
```

## INSTALLING

#### Installing from go

Using Go Password Manager is easy. First, use `go get` to install the latest version. This command will install the `gpassmanager` and its dependencies:

`go get -u github.com/mrinjamul/gpassmanager`

#### Installing from Binaries

[Download](https://github.com/mrinjamul/gpassmanager/releases) for your platform

For Linux,

```sh
wget https://github.com/mrinjamul/gpassmanager/releases/download/v0.7.1/gpassmanager-linux-amd64-v0.7.1.tar.gz
tar xvf gpassmanager-linux-amd64-v0.7.1.tar.gz
chmod +x gpassmanager
sudo mv gpassmanager /usr/bin
```

or you can put the executable file into your env variables `$PATH`

For Android (Termux),

You need to have `tar wget`. To install simply type `pkg install tar wget`

```sh
cd ~
wget https://github.com/mrinjamul/gpassmanager/releases/download/v0.7.1/gpassmanager-linux-arm-v0.7.1.tar.gz
tar xvf gpassmanager-linux-arm-v0.7.1.tar.gz
chmod +x gpassmanager
mv gpassmanager ../usr/bin
```

[Note: if new version available you need to download and install by the same process. The above instructions will install v0.7.1 .]

## Usage

    Simple Password Manager Application
    Licensed under the Apache License, Version 2.0 (the "License");

    Usage:
    gpassmanager [command]

    Available Commands:
    add         Add new password
    change      Change Master Password
    export      export your data to a file (master key will be also exported)
    generate    Generate secure password
    help        Help about any command
    import      import password(s) from a file
    license     Print the license
    remove      remove an Account from password manager
    reset       Erase all passwords including master key
    search      Search password account in the password store
    version     Prints version
    view        view all passwords

    Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)
    -h, --help            help for gpassmanager

    Use "gpassmanager [command] --help" for more information about a command.

## Documentations

- [Getting Started](docs/README.md)

## Links

- [Websites](https://mrinjamul.github.io/gpassmanager)

## Contributing

- [CONTRIBUTING](CONTRIBUTING.md)

## CODE OF CONDUCT

- [CODE OF CONDUCT](CODE_OF_CONDUCT.md)

## License

- Apache-2.0
