---
code: null
id: I_kwDOGMJgxc5DwtA1
databaseId: 1136840757
number: 3
url: https://github.com/kamilsk/godownloader/issues/3
title: "refactor Taskfile and sync with go-tool template"
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
createdAt: 2022-02-14T06:33:37Z
updatedAt: 2022-02-14T11:39:08Z
lastEditedAt: null
closedAt: 2022-02-14T11:39:08Z
---

# refactor Taskfile and sync with go-tool template

1. replace `git describe --tags --always` -> `run version`
2. make `clean` more safe, do not remove latest release
3. add `--rm-dist`
4. reorg script

check content for template

<!-- 2022-02-14T11:39Z https://github.com/kamilsk/godownloader/issues/3#issuecomment-1038983589
fixed in https://github.com/kamilsk/godownloader/blob/v0.1.0/Taskfile
-->
