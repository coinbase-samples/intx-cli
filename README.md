# INTX CLI README

## Overview

The INTX CLI is a sample Command Line Interface (CLI) application that generates requests to and receives responses from [Coinbase International Exchange (INTX)](https://international.coinbase.com/) [REST APIs](https://docs.cloud.coinbase.com/intx/reference). The INTX CLI is written in Go, using [Cobra](https://github.com/spf13/cobra) in conjunction with the [INTX Go SDK](https://github.com/coinbase-samples/intx-sdk-go).

## License

The INTX CLI is free and open source and released under the [Apache License, Version 2.0](LICENSE.txt).

The application and code are only available for demonstration purposes.

## Usage

To begin, navigate to your preferred directory for development and clone the INTX CLI repository and enter the directory using the following commands:

```
git clone https://github.com/coinbase-samples/intx-cli
cd intx-cli
```

Next, pass an environment variable via your terminal called `INTX_CREDENTIALS` with your API and portfolio information.

INTX API credentials can be created in the INTX web console under APIs.

`INTX_CREDENTIALS` should match the following format:
```
export INTX_CREDENTIALS='{
"accessKey":"ACCESSKEY_HERE",
"passphrase":"PASSPHRASE_HERE",
"signingKey":"SIGNINGKEY_HERE",
"portfolioId":"PORTFOLIOID_HERE",
}'
```

You may also pass an environment variable called `intxCliTimeout` which will override the default request timeout of 7 seconds. This value should be an integer in seconds.

To build the application binary, simply run:

```
make
```

This command compiles the application and creates a binary named `intxctl` in the current directory.

To ensure your project's dependencies are up-to-date, run:
```
go mod tidy
```

To install `intxctl` to `/usr/local/bin` for easy access from any location, run:

```
make install
```

This command moves the `intxctl` binary to `/usr/local/bin/` and sets the appropriate permissions. You might need `sudo` access to move the file to `/usr/local/bin/`.

To verify that `intxctl` is correctly installed and accessible from any location, you can run:

```
intxctl
```

This command should display all available requests if the installation was successful.

Finally, to run commands for each endpoint, use the following format to test each endpoint. Please note that many endpoints require flags, which are detailed with the `--help` flag.

```
intxctl list-portfolios
```

```
intxctl create-order --help
```