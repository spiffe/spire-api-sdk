# Contribution Guidelines

This document outlines the contribution guidelines for the SPIRE API SDK.

This project follows the contribution and governance guidelines from the SPIFFE
project (see
[CONTRIBUTING](https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md)
and [GOVERNANCE](https://github.com/spiffe/spiffe/blob/master/GOVERNANCE.md))

## Prerequisites

The [Makefile](/Makefile) in the project is set up to download required
dependencies for code generation.

### Updating Dependency Versions

The [Makefile](/Makefile) uses internal variables or inspects [go.mod](/go.mod)
to determine the versions of various tools in the toolchain. See the
[Makefile](/Makefile) for specifics.

## Generating Service Definitions

To (re)generate service definitions do the following:

```sh
$ make
```

If you are adding a new .proto file, you first need to update the `Makefile`
and add the .proto file to the relevant variables.

## Consuming Changes in SPIRE

SPIRE's main branch will always depend on the `next` tag. When changes from
this repository need to make their way into SPIRE ahead of a release, the
`next` tag will be updated.

When a SPIRE release is imminent, a tag for that version is first pushed to
this repository. The `go.mod` file in the SPIRE release branch is then updated
to consume the specific version tag for the release.

For testing changes locally, you can add a temporary `replace` directive to the
SPIRE `go.mod` file. However, care must be taken to not push this change up to
SPIRE.
