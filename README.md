# gpassmanager

Simple Password Manager Application

A Commandline application for managing your passwords securely.
It's a very secure password manager.

## INSTALLING

#### Installing from go

Using Go Password Manager is easy. First, use `go get` to install the latest version. This command will install the `gpassmanager` and its dependencies:

`go get -u github.com/mrinjamul/gpassmanager`

#### Installing from Binaries

[Releases](https://github.com/mrinjamul/gpassmanager/releases)

Download for your platform and put the executable file into your env `$PATH` .

For Linux,

```sh
unzip gpassmanager-linux-[whatever].zip
chmod +x gpassmanager
sudo mv gpassmanager /usr/bin
```

## Usage

    Simple Password Manager Application version 0.3.0

    Usage:
    gpassmanager [command]

    Available Commands:
    add         Add new password
    change      Change Master Password
    export      export your data to a file (master key will be also exported)
    help        Help about any command
    import      import data from a file
    remove      remove an Account from password manager
    reset       Erase all passwords including master key
    version     Prints version
    view        view all passwords

    Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)
    -h, --help            help for gpassmanager

    Use "gpassmanager [command] --help" for more information about a command.

## Documentations

[Getting Started](docs/README.md)

## Links

- [Websites](https://mrinjamul.github.io/gpassmanager)

## CODE OF CONDUCT

[CODE OF CONDUCT](CODE_OF_CONDUCT.md)

## License

- Apache-2.0
