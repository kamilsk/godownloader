---
code: null
id: I_kwDOGMJgxc5jc-Eh
databaseId: 1668538657
number: 40
url: https://github.com/kamilsk/godownloader/issues/40
title: "dev: publishing is unstable"
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
createdAt: 2023-04-14T16:10:02Z
updatedAt: 2023-04-14T18:44:27Z
lastEditedAt: null
closedAt: 2023-04-14T18:44:27Z
---

# dev: publishing is unstable

**Details**

```bash
$ run patch
Stopped at 3c7874b...  patch godownloader and update goreleaser
You can amend the commit now, with

  git commit --amend '-S'

Once you are satisfied with your changes, run

  git rebase --continue
go: downloading github.com/goreleaser/goreleaser v1.17.1
[detached HEAD 5e43fc6] patch godownloader and update goreleaser
 Date: Tue Feb 14 16:12:09 2023 +0800
 8 files changed, 3343 insertions(+), 358 deletions(-)
 create mode 100644 patch.go
Successfully rebased and updated refs/heads/extended.
Current branch extended is up to date.
Deleted tag 'homebrew' (was eeb926a)
To github.com:kamilsk/godownloader.git
 ! [remote rejected] homebrew (missing necessary objects)
error: failed to push some refs to 'github.com:kamilsk/godownloader.git'
Enumerating objects: 36, done.
Counting objects: 100% (36/36), done.
Delta compression using up to 8 threads
Compressing objects: 100% (20/20), done.
Writing objects: 100% (26/26), 130.80 KiB | 1.23 MiB/s, done.
Total 26 (delta 11), reused 13 (delta 5), pack-reused 0
remote: Resolving deltas: 100% (11/11), completed with 3 local objects.
To github.com:kamilsk/godownloader.git
 + c50125f...4f7a11b extended -> extended (forced update)
To github.com:kamilsk/godownloader.git
 ! [rejected]        homebrew -> homebrew (already exists)
error: failed to push some refs to 'github.com:kamilsk/godownloader.git'
hint: Updates were rejected because the tag already exists in the remote.
git push origin refs/tags/homebrew
```

<!-- 2023-04-14T18:44Z https://github.com/kamilsk/godownloader/issues/40#issuecomment-1509074582
fixed by https://github.com/kamilsk/godownloader/blob/9af943d2c984e590e1da082ede315419a6d5463a/Taskfile#L179
-->
