## Get Started

    Simple Password Manager Application version 0.4.1

    Usage:
    gpassmanager [command]

    Available Commands:
    add         Add new password
    change      Change Master Password
    export      export your data to a file (master key will be also exported)
    generate    Generate secure password
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

    view all passwords

    Usage:
    gpassmanager view [flags]

    Flags:
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
    -h, --help   help for reset

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

    Usage: gpassmanager import "[file location]"

    Usage:
    gpassmanager import [flags]

    Flags:
    -h, --help   help for import

    Global Flags:
        --config string   config file (default is $HOME/.gpassmanager.yaml)
