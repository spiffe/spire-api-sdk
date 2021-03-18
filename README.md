This repository contains the service definitions and code generated stubs for
public [SPIRE](https://github.com/spiffe/spire) APIs.

## Versioning

This repository is tagged along with SPIRE releases with the same name, even if
there are no changes to the APIs between SPIRE versions. This allows consumers
to always pick a tag that matches up with their deployment. Even so, SPIRE
maintains API compatibility between SPIRE versions. SPIRE will clearly indicate
in the CHANGELOG when APIs are deprecated and issue warnings at runtime when
they are used well in advance of any removal.

## Migrating Code

To migrate existing code that consumed SPIRE service definitions from the
github.com/spiffe/spire/proto/spire Go module, see
[MIGRATING](/MIGRATING.md).

## Contributing

This repository follows the same governance and contribution guidelines as the
[SPIRE](https://github.com/spiffe/spire) project.

For specifics on getting started, see [CONTRIBUTING](/CONTRIBUTING.md).

Please open [Issues](https://github.com/spiffe/spire/issues) to request features or file bugs.
