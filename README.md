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

Build the application binary and specify an output name, e.g. `intxctl`:

```
go build -o intxctl
```

To ensure your project's dependencies are up-to-date, run:
```
go mod tidy
```

To make your application easily accessible from any location, move the binary you created to a directory that's already in your system's PATH. For example, these are the commands to move `intxctl` to `/usr/local/bin`, as well as set permissions to reduce risk:

```
sudo mv intxctl /usr/local/bin/
chmod 755 /usr/local/bin/intxctl
```

To verify that your application is installed correctly and accessible from any location, run the following command. It will include all available requests:

```
intxctl
```

Finally, to run commands for each endpoint, use the following format to test each endpoint. Please note that many endpoints require flags, which are detailed with the `--help` flag.

```
intxctl list-portfolios
```

```
intxctl create-order --help
```