# Migrating Code

Take the following steps to migrate code using the deprecated
`github.com/spiffe/spire/proto/spire` Go module.

- `go get` the desired version from this repository, e.g. `go get github.com/spiffe/spire-api-sdk@v1.0.0`.
- Rename the service package imports:
  - V1 Agent API
    - Old: `"github.com/spiffe/spire/proto/spire/api/server/agent/v1"`
    - New: `agentv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/agent/v1"`
  - V1 Bundle API
    - Old: `"github.com/spiffe/spire/proto/spire/api/server/bundle/v1"`
    - New: `bundlev1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bundle/v1"`
  - V1 Debug API
    - Old: `"github.com/spiffe/spire/proto/spire/api/server/debug/v1"`
    - New: `debugv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/debug/v1"`
  - V1 Entry API
    - Old: `"github.com/spiffe/spire/proto/spire/api/server/entry/v1"`
    - New: `entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"`
  - V1 SVID API
    - Old: `"github.com/spiffe/spire/proto/spire/api/server/svid/v1"`
    - New: `svidv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/svid/v1"`
- Rename the types package imports:
  - Old: `"github.com/spiffe/spire/proto/spire/types"`
  - New: `"github.com/spiffe/spire-api-sdk/proto/spire/api/types"`

Renaming the service package imports will almost certainly require code imports
to use the newly named imports unless existing code already named the import
for other reasons. This should be a mostly mechanical process and can be
aided by a variety of tools available in the Go community.
