---
code: null
id: I_kwDOGMJgxc5Rc-bi
databaseId: 1366550242
number: 17
url: https://github.com/kamilsk/godownloader/issues/17
title: "Taskfile: refactor run sync"
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
createdAt: 2022-09-08T15:17:39Z
updatedAt: 2022-12-18T10:36:34Z
lastEditedAt: 2022-10-22T15:45:55Z
closedAt: 2022-12-18T10:36:34Z
---

# Taskfile: refactor run sync

**Motivation:** prepare to automate job instead of manual fasion.

**Algorythm**

```bash
$ pull
$ git untag v0.1.0
$ git rebase -i HEAD~2
# -> e, up goreleaser, git amend
$ go mod edit -require='github.com/goreleaser/goreleaser@v1.11.5'
$ go mod tidy
$ git add .
$ git amend
$ git rebase --contrinue
$ run sync v0.1.0
$ rm -rf dist/
```

<!-- 2022-12-18T10:36Z https://github.com/kamilsk/godownloader/issues/17#issuecomment-1356769296
fixed https://github.com/kamilsk/godownloader/blob/c47685b527e10f92be60d3a2a4a65b96b35ffd84/Taskfile#L67-L75
-->
