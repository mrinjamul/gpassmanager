# gpassmanager

Simple Password Manager cli Application

## INSTALLING

Using Go Password Manager is easy. First, use go get to install the latest version of the library. This command will install the `gpassmanager` generator executable along with the library and its dependencies:

`go get -u github.com/mrinjamul/gpassmanager`

## Usage

    Simple Password Manager Application

    Usage:
    gpassmanager [command]

    Available Commands:
    add         Add new password
    help        Help about any command
    reset       Erase all passwords including master key
    version     Prints version
    view        view all passwords

    Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)
    -h, --help            help for gpassmanager

    Use "gpassmanager [command] --help" for more information about a command.

## License

- Apache-2.0
