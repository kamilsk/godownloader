---
code: null
id: I_kwDOGMJgxc5aONL3
databaseId: 1513673463
number: 27
url: https://github.com/kamilsk/godownloader/issues/27
title: "dev: refactor Taskfile"
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
createdAt: 2022-12-29T12:06:59Z
updatedAt: 2023-01-01T15:36:02Z
lastEditedAt: null
closedAt: 2023-01-01T15:36:01Z
---

# dev: refactor Taskfile

**Motivation:** better DX.

Instead of

```
#  - run -d release v1.14.0 > bin/transaction.sh
#  - run release v1.14.0
#  - run release
#  - run -d patch v1.14.0 v0.1.0 > bin/transaction.sh
#  - run patch v1.14.0 v0.1.0
#  - run patch
```

I want to

```
#  - run release
#  - run release v0.1.0
#  - run -d release v0.1.0 > bin/transaction.sh
#  - run patch -- v1.14.0
#  - run patch v0.1.0 -- v1.14.0
#  - run -d patch -- v1.14.0 > bin/transaction.sh
```

<!-- 2023-01-01T15:36Z https://github.com/kamilsk/godownloader/issues/27#issuecomment-1368470228
fixed by https://github.com/kamilsk/godownloader/blob/fc135d340528a78ed49ac323bc0d55758c8bd992/Taskfile#L74-L93
-->
