# Repofuel Ingest

[![Go](https://https://github.com/repofuel/repofuel-ingest/workflows/Go/badge.svg)](https://https://github.com/repofuel/repofuel-ingest/actions?query=workflow%3AGo)
[![Docker Image CI](https://https://github.com/repofuel/repofuel-ingest/workflows/Docker%20Image%20CI/badge.svg?branch=dev-server)](https://https://github.com/repofuel/repofuel-ingest/actions?query=workflow%3A%22Docker+Image+CI%22)

Here where we collect and analyze the data from repositories.

## Build

### Build libgit2

You must have the following installed on your system to run the necessary install script.

- `cmake`
- `wget`

Once the packages are installed run the following:

```shell script
bash ./libgit2/build-libgit2-dynamic.sh
```
