## INSTALLING

#### Installing from go

Using Go Password Manager is easy. First, use `go get` to install the latest version. This command will install the `gpassmanager` and its dependencies:

`go get -u github.com/mrinjamul/gpassmanager`

#### Installing from Binaries

[Download](https://github.com/mrinjamul/gpassmanager/releases) for your platform

For Linux,

```sh
wget https://github.com/mrinjamul/gpassmanager/releases/download/v1.0.2/gpassmanager-linux-amd64-v1.0.2.tar.gz
tar xvf gpassmanager-linux-amd64-v1.0.2.tar.gz
chmod +x gpassmanager
sudo mv gpassmanager /usr/bin
```

or you can put the executable file into your env variables `$PATH`

For Android (Termux),

You need to have `tar wget`. To install simply type `pkg install tar wget`

```sh
cd ~
wget https://github.com/mrinjamul/gpassmanager/releases/download/v1.0.2/gpassmanager-linux-arm-v1.0.2.tar.gz
tar xvf gpassmanager-linux-arm-v1.0.2.tar.gz
chmod +x gpassmanager
mv gpassmanager ../usr/bin
```

[Note: if new version available you need to download and install by the same process. The above instructions will install v1.0.2 .]

## Get Started

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

## Help

    Help provides help for any command in the application.
    Simply type gpassmanager help [path to command] for full details.

    Usage:
    gpassmanager help [command] [flags]

    Flags:
    -h, --help   help for help

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Add

    Add new password

    Usage:
    gpassmanager add [flags]

    Flags:
    -h, --help   help for add

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## View

    view a particular password or entire passwords
    Example: gpassmanager view
    then gpassmanager view 1

    Usage:
    gpassmanager view [flags]

    Flags:
    -a, --all    view all passwords in the store
    -h, --help   help for view

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Remove

    remove an Account from password manager

    Usage:
    gpassmanager remove [flags]

    Aliases:
    remove, rm

    Flags:
    -h, --help   help for remove

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Search

    Search password account in the password store

    Usage:
    gpassmanager search [flags]

    Flags:
    -h, --help   help for search

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Change Password

    Change Master Password
    Example, gpassmanager change --passwd

    Usage:
    gpassmanager change [flags]

    Flags:
    -h, --help     help for change
    -p, --passwd   change master key for the Data

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Reset

    It's like a hard reset.
    If you forget your master key then you have to perform a hard reset

    Usage:
    gpassmanager reset [flags]

    Flags:
    -h, --help      help for reset
    -r, --restore   restore last reset database

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Generate

    Generate secure password

    Usage:
    gpassmanager generate [flags]

    Aliases:
    generate, gen

    Flags:
    -h, --help   help for generate

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## License

    Print the license

    Usage:
    gpassmanager license [flags]

    Flags:
    -h, --help   help for license

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Version

    Prints version

    Usage:
    gpassmanager version [flags]

    Flags:
    -h, --help   help for version

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Export

    Usage: gpassmanager export "export filename"
            or
    gpassmanager export

    Usage:
    gpassmanager export [flags]

    Flags:
    -h, --help   help for export

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)

## Import

    gpassmanager import "[file location]"

    Usage:
    gpassmanager import [flags]

    Flags:
    -c, --csv    Import CSV file into the password manager (Currently Google password csv file is supported)
    -h, --help   help for import

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)
