---
code: null
id: I_kwDOGMJgxc8AAAABTM1IDw
databaseId: 5583489039
number: 84
url: https://github.com/kamilsk/godownloader/issues/84
title: "sync: support goreleaser v2 configuration"
issueType: null
labels:
  - "type: bug"
  - "severity: major"
  - "scope: code"
  - "impact: high"
  - "effort: medium"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2026-09-25T13:29:47Z
updatedAt: 2026-11-08T14:15:44Z
lastEditedAt: null
closedAt: null
---

# sync: support goreleaser v2 configuration

**Motivation:** godownloader stops at a GoReleaser v2 configuration, so a project on v2 cannot regenerate its installer, and the committed one silently drifts from the releases. [octomation/maintainer](https://github.com/octomation/maintainer/blob/603a171/bin/install) still ships `bin/install` generated on 2021-12-11, when its `.goreleaser.yml` had `archives.replacements`. The replacements were dropped in [d6fff63](https://github.com/octomation/maintainer/commit/d6fff63), and the script keeps asking for `maintainer_0.1.0-rc11_macOS-64bit_arm.tar.gz` (404), while the release has `maintainer_0.1.0-rc11_darwin-arm64.tar.gz`.

**PoC**, Go 1.27.1, `extended` at 9213317:

```bash
go build .
# golang.org/x/tools/internal/tokeninternal/tokeninternal.go:64:9: invalid array length -delta * delta
go get golang.org/x/tools@v0.50.0 && go build -o /tmp/godownloader .  # v0.24.0 still fails

git clone https://github.com/octomation/maintainer && cd maintainer && git checkout 603a171
/tmp/godownloader --repo=octomation/maintainer .goreleaser.yml
# ⨯ failed error=unable to parse: only configurations files on version: 1 are supported, yours is version: 2

git show d6fff63:.goreleaser.yml > /tmp/v1.yml
/tmp/godownloader --repo=octomation/maintainer /tmp/v1.yml
# the same archives in the v1 schema generate fine:
# adjust_os and adjust_arch are no-ops, NAME=${BINARY}_${VERSION}_${OS}-${ARCH}
```

**Where the cause likely is:**

- `main.go:17-19` import `pkg/config`, `pkg/context` and `pkg/defaults` of goreleaser v1.24.0, whose loader rejects `version: 2`.
- `shell_godownloader.go:129-132` and `:153` render `format_overrides[].format` and `archives.format`, while v2 uses the `formats` lists.
- `makeName` in `main.go:97-140` only warns about a conditional in `name_template` and renders it with `missingkey=zero`. GoReleaser replaced `archives.replacements` with such conditionals, e.g. `{{ if eq .Os "darwin" }}macOS{{ else }}{{ .Os }}{{ end }}`, so a v2 archive name can come out wrong without an error.
- `go.mod` resolves `golang.org/x/tools` to `v0.21.1-0.20240508182429-e35e4ccd0d2d`, which does not build with current Go.
