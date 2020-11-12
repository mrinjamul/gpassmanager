# gpassmanager

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
wget https://github.com/mrinjamul/gpassmanager/releases/download/v0.5.0/gpassmanager-linux-amd64-current.zip
unzip gpassmanager-linux-amd64-current.zip
chmod +x gpassmanager
sudo mv gpassmanager /usr/bin
```

or you can put the executable file into your env variables `$PATH`

For Android (Termux),

You need to have `zip unzip wget`. To install simply type `pkg install zip unzip wget`

```sh
cd ~
wget https://github.com/mrinjamul/gpassmanager/releases/download/v0.5.0/gpassmanager-linux-arm-current.zip
unzip gpassmanager-linux-arm-current.zip
chmod +x gpassmanager
mv gpassmanager ../usr/bin
```

[Note: if new version available you need to download and install by the same process. The above instructions will install v0.5.0 .]

## Usage

    Simple Password Manager Application version 0.6.0

    Usage:
    gpassmanager [command]

    Available Commands:
    add         Add new password
    change      Change Master Password
    export      export your data to a file (master key will be also exported)
    generate    Generate secure password
    help        Help about any command
    import      import password(s) from a file
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

[Getting Started](docs/README.md)

## Links

- [Websites](https://mrinjamul.github.io/gpassmanager)

## CODE OF CONDUCT

[CODE OF CONDUCT](CODE_OF_CONDUCT.md)

## License

- Apache-2.0
