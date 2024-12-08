---
code: null
id: I_kwDOGMJgxc4850xf
databaseId: 1021791327
number: 2
url: https://github.com/kamilsk/godownloader/issues/2
title: "up dependency of goreleaser"
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
createdAt: 2021-10-09T19:51:06Z
updatedAt: 2021-12-12T13:48:17Z
lastEditedAt: null
closedAt: 2021-12-12T13:48:17Z
---

# up dependency of goreleaser

```
$ make dist-dump
   • reading file ".goreleaser.yml"
   • loading config file       file=.goreleaser.yml
   ⨯ failed                    error=unable to parse: yaml: unmarshal errors:
  line 25: field tap not found in type config.Homebrew
make: *** [dist-dump] Error 1
```

<!-- 2021-10-10T13:02Z https://github.com/kamilsk/godownloader/issues/2#issuecomment-939478624
create issue to broken changelog for go-tools
-->

<!-- 2021-12-12T13:48Z https://github.com/kamilsk/godownloader/issues/2#issuecomment-991901856
fixed
-->
