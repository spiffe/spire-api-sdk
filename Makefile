.PHONY: default help

# There is no reason GOROOT should be set anymore. Unset it so it doesn't mess
# with our go toolchain detection/usage.
ifneq ($(GOROOT),)
    export GOROOT=
endif

default: generate

help:
	@echo "Usage: make <target>"
	@echo
	@echo "  generate         - generate protocol buffer and gRPC stub code (default)"
	@echo "  generate-check   - ensure generated code is up to date"

protos := \
	proto/spire/api/types/agent.proto \
	proto/spire/api/types/attestation.proto \
	proto/spire/api/types/bundle.proto \
	proto/spire/api/types/entry.proto \
	proto/spire/api/types/federationrelationship.proto \
	proto/spire/api/types/federateswith.proto \
	proto/spire/api/types/jointoken.proto \
	proto/spire/api/types/jwtsvid.proto \
	proto/spire/api/types/logger.proto \
	proto/spire/api/types/selector.proto \
	proto/spire/api/types/spiffeid.proto \
	proto/spire/api/types/status.proto \
	proto/spire/api/types/x509svid.proto \


apiprotos := \
	proto/spire/api/agent/debug/v1/debug.proto \
	proto/spire/api/agent/delegatedidentity/v1/delegatedidentity.proto \
	proto/spire/api/agent/logger/v1/logger.proto \
	proto/spire/api/server/agent/v1/agent.proto \
	proto/spire/api/server/bundle/v1/bundle.proto \
	proto/spire/api/server/debug/v1/debug.proto \
	proto/spire/api/server/entry/v1/entry.proto \
	proto/spire/api/server/localauthority/v1/localauthority.proto \
	proto/spire/api/server/logger/v1/logger.proto \
	proto/spire/api/server/trustdomain/v1/trustdomain.proto \
	proto/spire/api/server/svid/v1/svid.proto \

# Used to force some rules to run every time
FORCE: ;

############################################################################
# OS/ARCH detection
############################################################################
os1=$(shell uname -s)
os2=
ifeq ($(os1),Darwin)
os1=darwin
os2=osx
else ifeq ($(os1),Linux)
os1=linux
os2=linux
else
$(error unsupported OS: $(os1))
endif

arch1=$(shell uname -m)
ifeq ($(arch1),x86_64)
arch2=amd64
else ifeq ($(arch1),aarch64)
arch2=arm64
else ifeq ($(arch1),arm64)
arch2=arm64
else
$(error unsupported ARCH: $(arch1))
endif

build_dir := ${CURDIR}/.build/$(os1)-$(arch1)

#############################################################################
# Go
#############################################################################

go_version_full := $(shell cat .go-version)
go_version := $(go_version_full:.0=)
go_dir := $(build_dir)/go/$(go_version)
go_bin_dir := $(go_dir)/bin
go_url = https://storage.googleapis.com/golang/go$(go_version).$(os1)-$(arch2).tar.gz
go_path := PATH="$(go_bin_dir):$(PATH)"

# go-check checks to see if there is a version of Go available matching the
# required version. The build cache is preferred. If not available, it is
# downloaded into the build cache. Any rule needing to invoke tools in the go
# toolchain should depend on this rule and then prepend $(go_bin_dir) to their
# path before invoking go or use $(go_path) go which already has the path prepended.
# Note that some tools (e.g. anything that uses golang.org/x/tools/go/packages)
# execute on the go binary and also need the right path in order to locate the
# correct go binary.
go-check:
ifneq (go$(go_version), $(shell $(go_path) go version 2>/dev/null | cut -f3 -d' '))
ifeq ($(go_version),)
	$(error unable to determine go version)
endif
	@echo "Installing go$(go_version)..."
	@rm -rf $(dir $(go_dir))
	@mkdir -p $(go_dir)
	@curl -sSfL $(go_url) | tar xz -C $(go_dir) --strip-components=1
endif

#############################################################################
# protoc
#############################################################################

protoc_version = 3.20.1
ifeq ($(arch2),arm64)
protoc_url = https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-$(os2)-aarch_64.zip
else
protoc_url = https://github.com/protocolbuffers/protobuf/releases/download/v$(protoc_version)/protoc-$(protoc_version)-$(os2)-$(arch1).zip
endif
protoc_dir = $(build_dir)/protoc/$(protoc_version)
protoc_bin = $(protoc_dir)/bin/protoc

$(protoc_bin):
	@echo "Installing protoc $(protoc_version)..."
	@rm -rf $(dir $(protoc_dir))
	@mkdir -p $(protoc_dir)
	@curl -sSfL $(protoc_url) -o $(build_dir)/tmp.zip; unzip -q -d $(protoc_dir) $(build_dir)/tmp.zip; rm $(build_dir)/tmp.zip

#############################################################################
# protoc-gen-go
#############################################################################

protoc_gen_go_version := $(shell grep google.golang.org/protobuf go.mod | awk '{print $$2}')
protoc_gen_go_base_dir := $(build_dir)/protoc-gen-go
protoc_gen_go_dir := $(protoc_gen_go_base_dir)/$(protoc_gen_go_version)-go$(go_version)
protoc_gen_go_bin := $(protoc_gen_go_dir)/protoc-gen-go

$(protoc_gen_go_bin): | go-check
	@echo "Installing protoc-gen-go $(protoc_gen_go_version)..."
	@rm -rf $(protoc_gen_go_base_dir)
	@mkdir -p $(protoc_gen_go_dir)
	@$(go_path) go build -o $(protoc_gen_go_bin) google.golang.org/protobuf/cmd/protoc-gen-go

#############################################################################
# protoc-gen-go-grpc
#############################################################################

protoc_gen_go_grpc_version := v1.0.1
protoc_gen_go_grpc_base_dir := $(build_dir)/protoc-gen-go-grpc
protoc_gen_go_grpc_dir := $(protoc_gen_go_grpc_base_dir)/$(protoc_gen_go_grpc_version)-go$(go_version)
protoc_gen_go_grpc_bin := $(protoc_gen_go_grpc_dir)/protoc-gen-go-grpc

$(protoc_gen_go_grpc_bin): | go-check
	@echo "Installing protoc-gen-go-grpc $(protoc_gen_go_grpc_version)..."
	@rm -rf $(protoc_gen_go_grpc_base_dir)
	@mkdir -p $(protoc_gen_go_grpc_dir)
	@GOBIN=$(protoc_gen_go_grpc_dir) $(go_path) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(protoc_gen_go_grpc_version)

#############################################################################
# Code Generation
#############################################################################

.PHONY: generate
generate: $(protos:.proto=.pb.go) $(apiprotos:.proto=.pb.go) $(apiprotos:.proto=_grpc.pb.go)

%_grpc.pb.go: %.proto $(protoc_bin) $(protoc_gen_go_grpc_bin) FORCE
	@echo "compiling API $<..."
	@cd proto && \
		PATH="$(protoc_gen_go_grpc_dir):$(PATH)" \
		$(protoc_bin) \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(<:proto/%=%)

%.pb.go: %.proto $(protoc_bin) $(protoc_gen_go_bin) FORCE
	@echo "compiling $<..."
	@cd proto && \
		PATH="$(protoc_gen_go_dir):$(PATH)" \
		$(protoc_bin) \
		--go_out=. --go_opt=paths=source_relative \
		$(<:proto/%=%)

#############################################################################
# Code Generation Checks
#############################################################################

git_dirty := $(shell git status -s)

.PHONY: generate-check
generate-check:
ifneq ($(git_dirty),)
	$(error generate-check must be invoked on a clean repository)
endif
	@$(MAKE) generate
	@$(MAKE) git-clean-check

.PHONY: git-clean-check
git-clean-check:
ifneq ($(git_dirty),)
	git diff
	@echo "Git repository is dirty!"
	@false
else
	@echo "Git repository is clean."
endif
