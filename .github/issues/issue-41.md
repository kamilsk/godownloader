---
code: null
id: I_kwDOGMJgxc5jdaJA
databaseId: 1668653632
number: 41
url: https://github.com/kamilsk/godownloader/issues/41
title: "dev: publishing used incorrect tag"
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
createdAt: 2023-04-14T17:27:24Z
updatedAt: 2023-04-14T18:44:19Z
lastEditedAt: null
closedAt: 2023-04-14T18:44:19Z
---

# dev: publishing used incorrect tag

**Details**

- see [[issue-40]]
- after repeating
```bash
$ run patch
Stopped at 5e43fc6...  patch godownloader and update goreleaser
You can amend the commit now, with

  git commit --amend '-S'

Once you are satisfied with your changes, run

  git rebase --continue
[detached HEAD b87f697] patch godownloader and update goreleaser
 Date: Tue Feb 14 16:12:09 2023 +0800
 8 files changed, 3343 insertions(+), 358 deletions(-)
 create mode 100644 patch.go
Successfully rebased and updated refs/heads/extended.
From github.com:kamilsk/godownloader
 t [tag update]      homebrew   -> homebrew
Current branch extended is up to date.
Deleted tag 'homebrew' (was eeb926a)
To github.com:kamilsk/godownloader.git
 - [deleted]         homebrew
Enumerating objects: 36, done.
Counting objects: 100% (36/36), done.
Delta compression using up to 8 threads
Compressing objects: 100% (20/20), done.
Writing objects: 100% (26/26), 130.80 KiB | 1.04 MiB/s, done.
Total 26 (delta 11), reused 13 (delta 5), pack-reused 0
remote: Resolving deltas: 100% (11/11), completed with 3 local objects.
To github.com:kamilsk/godownloader.git
 + 4f7a11b...ee9388c extended -> extended (forced update)
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 810 bytes | 810.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:kamilsk/godownloader.git
 * [new tag]         homebrew -> homebrew
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 811 bytes | 811.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:kamilsk/godownloader.git
 * [new tag]         v0.1.9 -> v0.1.9
release not found
https://github.com/kamilsk/godownloader/releases/tag/homebrew
  • starting release...
  • loading config file                              file=.goreleaser.yml
  • loading environment variables
    • using token from "$GITHUB_TOKEN"
  • getting and validating git state
    • building...                                    commit=ee9388cea9017b10472acf0a89311306d1563274 latest tag=v0.1.9
  • parsing tag
  • setting defaults
  • checking distribution directory
  • loading go mod information
  • build prerequisites
  • writing effective config file
    • writing                                        config=dist/config.yaml
  • building binaries
    • building                                       binary=dist/godownloader_linux_arm64/godownloader
    • building                                       binary=dist/godownloader_darwin_amd64_v1/godownloader
    • building                                       binary=dist/godownloader_darwin_arm64/godownloader
    • building                                       binary=dist/godownloader_linux_amd64_v1/godownloader
    • took: 13s
  • generating changelog
  • archives
    • creating                                       archive=dist/godownloader_0.1.9_darwin-arm64.tar.gz
    • creating                                       archive=dist/godownloader_0.1.9_darwin-amd64.tar.gz
    • creating                                       archive=dist/godownloader_0.1.9_linux-arm64.tar.gz
    • creating                                       archive=dist/godownloader_0.1.9_linux-amd64.tar.gz
    • took: 6s
  • calculating checksums
  • homebrew tap formula
    • writing                                        formula=dist/godownloader.rb
  • publishing
    • scm releases
      • creating or updating release                 repo=kamilsk/godownloader tag=v0.1.9
      • release created                              name=v0.1.9 release-id=99576600 request-id=966A:98F4:3F2DF22:3FD4FB3:64398AD3
      • uploading to release                         file=dist/godownloader_0.1.9_linux-amd64.tar.gz name=godownloader_0.1.9_linux-amd64.tar.gz
      • uploading to release                         file=dist/checksums.txt name=checksums.txt
      • uploading to release                         file=dist/godownloader_0.1.9_darwin-arm64.tar.gz name=godownloader_0.1.9_darwin-arm64.tar.gz
      • uploading to release                         file=dist/godownloader_0.1.9_linux-arm64.tar.gz name=godownloader_0.1.9_linux-arm64.tar.gz
      • uploading to release                         file=dist/godownloader_0.1.9_darwin-amd64.tar.gz name=godownloader_0.1.9_darwin-amd64.tar.gz
      • published                                    url=https://github.com/kamilsk/godownloader/releases/tag/v0.1.9
      • took: 29s
    • homebrew tap formula
      • pushing                                      formula=Formula/godownloader.rb repo=kamilsk/homebrew-tap
      • took: 1s
  • took: 30s
  • storing release metadata
    • writing                                        file=dist/artifacts.json
    • writing                                        file=dist/metadata.json
  • announcing
  • release succeeded after 49s
Opening github.com/kamilsk/godownloader/releases/tag/v0.1.9 in your browser.
```

```bash
$ git tag --list | cat
homebrew
v0.1.0
v0.1.1
v0.1.2
v0.1.3
v0.1.4
v0.1.5
v0.1.6
v0.1.7
v0.1.8 # result of #40 issue
v0.1.9
```

<img width="514" alt="image" src="https://user-images.githubusercontent.com/1165416/232115016-feceaf14-ccd8-44e9-acba-2e8ecf3bfeec.png">

<!-- 2023-04-14T17:29Z https://github.com/kamilsk/godownloader/issues/41#issuecomment-1508996877
```bash
$ run -d patch
git rebase -i homebrew
go mod edit -require=github.com/goreleaser/goreleaser@latest
go mod tidy
git add .
eval GIT_COMMITTER_DATE=2023-04-14T20:17:30+03:00 git commit --amend --date=2023-02-14T16:12:09+08:00 --no-edit
git rebase --continue
git fetch origin --force --prune --tags
git fetch upstream --force --prune --tags
git rebase --onto v0.1.0 upstream/master
git tag -m 'sync with upstream' v0.1.10
git tag -d homebrew || true
git push origin :refs/tags/homebrew || true
git tag -m 'extend by homebrew' homebrew 9ea6f46b5e377dc658ad5adc44d3472212721d47
git push --force-with-lease origin extended
git push origin refs/tags/homebrew
git push origin refs/tags/v0.1.10
gh release -R kamilsk/godownloader delete v0.1.10 -y || true
gh release -R kamilsk/godownloader edit homebrew --draft=false --prerelease
goreleaser --clean --release-notes CHANGES.md
rm -rf dist
gh release -R kamilsk/godownloader view v0.1.10 -w
```
-->

<!-- 2023-04-14T18:44Z https://github.com/kamilsk/godownloader/issues/41#issuecomment-1509074463
fixed by https://github.com/kamilsk/godownloader/blob/9af943d2c984e590e1da082ede315419a6d5463a/Taskfile#L179
-->
