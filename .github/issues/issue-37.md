---
code: null
id: I_kwDOGMJgxc5gIV9C
databaseId: 1612799810
number: 37
url: https://github.com/kamilsk/godownloader/issues/37
title: "taskfile: algorithm is incorrect"
issueType: null
labels: []
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2023-03-07T06:15:48Z
updatedAt: 2023-03-10T19:51:24Z
lastEditedAt: null
closedAt: 2023-03-10T19:51:24Z
---

# taskfile: algorithm is incorrect

```bash
$ git rebase -i homebrew # <----- why?
$ go mod edit -require=github.com/goreleaser/goreleaser@latest
$ go mod tidy
$ git add .
$ amend
$ git rebase --continue
$ git fetch origin --force --prune --tags
$ git fetch upstream --force --prune --tags
$ git rebase --onto v0.1.0 upstream/master # <---- and then, they must be equal
$ git tag -m 'sync with upstream' v0.1.4
$ git tag -d homebrew
$ git push origin :refs/tags/homebrew
$ git tag -m 'extend by homebrew' homebrew 0ca80b45b4e9a11df85f7266c39dd05438d77e2c
$ git push --force-with-lease origin extended
$ git push origin refs/tags/homebrew
$ git push origin refs/tags/v0.1.4
$ gh release -R kamilsk/godownloader delete v0.1.4 -y
$ gh release -R kamilsk/godownloader edit homebrew --draft=false --prerelease
$ goreleaser --clean --release-notes CHANGES.md
$ rm -rf dist
$ gh release -R kamilsk/godownloader view v0.1.4 -w
```

<!-- 2023-03-10T19:51Z https://github.com/kamilsk/godownloader/issues/37#issuecomment-1464333855
fixed
-->
