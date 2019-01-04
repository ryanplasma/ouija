# Ouija

Ouija is a CLI for communicating with a local Phantom VM while building custom Phantom apps.

This was built with the intent of being used with a locally run Phantom VM and not with a production environment so security was not a top priority.

Please do not use this with a production system or with production credentials.

## Dependencies

- rsync
- scp

Only tested on OSX so far

## Installation

1. `go build`

   - Clone this repo
   - `go build`
   - Move the binary where you want it to be (`/usr/bin/`, etc.)

2. `go get`
   - `go get github.com/ryanplasma/ouija`
   - `cd $GOPATH/src/github.com/ryanplasma/ouija`
   - `go install`

## Usage

### Commands

- `push, p` - Push the app code to the Phantom server
- `build, b` - Build the app code on the Phantom server
- `download, d` - Download the app tarball from the Phantom server
- `help, h` - Shows a list of commands or help for one command

### Flags

Flags can be set in multiple ways following this order of precedence:

1. Command line flag value from user
2. Environment variable (if specified)
3. Configuration file (if specified)
4. Default defined on the flag

At the moment, ouija requires a `ouija.yml` file in your app directory, even if it remains blank.

Flags include:

- User
  - Command Line Flag: `--user`
  - Environment Variable: `$OUIJA_USER`
  - Config File Value: `user`
  - Defaults to your current system user
  - Used in: Push, Build, Download
- Password
  - Command Line Flag: `--password`
  - Environment Variable: `$OUIJA_PASSWORD`
  - Config File Value: N/A
  - Defaults to an empty string
  - Used in: Build
- Host
  - Command Line Flag: `--host`
  - Environment Variable: `$OUIJA_HOST`
  - Config File Value: `host`
  - Defaults to 127.0.0.1
  - Used in: Push, Build, Download
- Port
  - Command Line Flag: `--port`
  - Environment Variable: `$OUIJA_PORT`
  - Config File Value: `port`
  - Defaults to 22
  - Used in: Build
- App
  - Command Line Flag: `--app`
  - Environment Variable: `$OUIJA_APP`
  - Config File Value: `app`
  - Defaults to the name of the current directory
  - Used in: Push, Build, Download
- Load
  - Command Line Flag: `--load`
  - Environment Variable: N/A
  - Config File Value: N/A
  - Defaults to `ouija.yml`
  - Used in: Push, Build, Download
  - Use this to override the config file name

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
