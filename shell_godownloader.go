package main

import (
	"fmt"
	"log"
)

func processGodownloader(repo string, filename string) {
	cfg, err := Load(repo, filename)
	if err != nil {
		log.Fatalf("Unable to parse: %s", err)
	}
	// get name template
	name, err := makeName(cfg.Archive.NameTemplate)
	cfg.Archive.NameTemplate = name
	if err != nil {
		log.Fatalf("Unable generate name: %s", err)
	}

	shell, err := makeShell(shellGodownloader, cfg)
	if err != nil {
		log.Fatalf("Unable to generate shell: %s", err)
	}
	fmt.Println(shell)
}

var shellGodownloader = `#!/bin/sh
set -e
#  Code generated by godownloader. DO NOT EDIT.
#

usage() {
  this=$1
  cat <<EOF
$this: download go binaries for {{ $.Release.GitHub.Owner }}/{{ $.Release.GitHub.Name }}

Usage: $this [-b] bindir [version]
  -b sets bindir or installation directory, default "./bin"
   [version] is a version number from
   https://github.com/{{ $.Release.GitHub.Owner }}/{{ $.Release.GitHub.Name }}/releases
   If version is missing, then an attempt to find the latest will be found.

Generated by godownloader
 https://github.com/goreleaser/godownloader

EOF
  exit 2
}` + shellfn + `
is_supported_platform() {
  platform=$1
  found=1
  case "$platform" in
  {{- range $goos := $.Build.Goos }}{{ range $goarch := $.Build.Goarch }}
    {{ if not (eq $goarch "arm") }}{{ $goos }}/{{ $goarch }}) found=0 ;; {{ end }}
  {{- end }}{{ end }}
  {{- if $.Build.Goarm }}
  {{- range $goos := $.Build.Goos }}{{ range $goarch := $.Build.Goarch }}{{ range $goarm := $.Build.Goarm }}
  {{- if eq $goarch "arm" }}{{ $goos }}/armv{{ $goarm }}) found=0 ;;
{{ end }}
  {{- end }}{{ end }}{{ end }}
  {{- end }}
  esac
  {{- if $.Build.Ignore }}
  case "$platform" in 
    {{- range $ignore := $.Build.Ignore }}
    {{ $ignore.Goos }}/{{ $ignore.Goarch }}{{ if $ignore.Goarm }}v{{ $ignore.Goarm }}{{ end }}) found=1 ;; 
    {{- end -}}
  esac
  {{- end }}
  return $found
}

parse_args() {
  #BINDIR is ./bin unless set be ENV
  # over-ridden by flag below

  BINDIR=${BINDIR:-./bin}
  while getopts "b:" arg; do
    case "$arg" in
      b) BINDIR="$OPTARG" ;;
      \?) usage "$0" ;;
    esac
  done
  shift $((OPTIND - 1))
  VERSION=$1
}

OWNER={{ $.Release.GitHub.Owner }}
REPO={{ $.Release.GitHub.Name }}
BINARY={{ .Build.Binary }}
FORMAT={{ .Archive.Format }}

parse_args "$@"

uname_os_check
uname_arch_check

OS=$(uname_os)
ARCH=$(uname_arch)
PREFIX="$OWNER/$REPO"
PLATFORM="${OS}/${ARCH}"
if is_supported_platform "$PLATFORM"; then
  # optional logging goes here
  true
else
  echo "${PREFIX}: platform $PLATFORM is not supported.  Make sure this script is up-to-date and file request at https://github.com/${PREFIX}/issues/new"
  exit 1
fi

if [ -z "${VERSION}" ]; then
  echo "$PREFIX: checking GitHub for latest version"
  VERSION=$(github_last_release "$OWNER/$REPO")
fi
# if version starts with 'v', remove it
VERSION=${VERSION#v}


# change format (tar.gz or zip) based on ARCH
{{- with .Archive.FormatOverrides }}
case ${ARCH} in
{{- range . }}
{{ .Goos }}) FORMAT={{ .Format }} ;;
esac
{{- end }}
{{- end }}

# adjust archive name based on OS
{{- with .Archive.Replacements }}
case ${OS} in
{{- range $k, $v := . }}
{{ $k }}) OS={{ $v }} ;;
{{- end }}
esac

# adjust archive name based on ARCH
case ${ARCH} in
{{- range $k, $v := . }}
{{ $k }}) ARCH={{ $v }} ;;
{{- end }}
esac
{{- end }}

echo "$PREFIX: found version ${VERSION} for ${OS}/${ARCH}"

{{ .Archive.NameTemplate }}
TARBALL=${NAME}.${FORMAT}
TARBALL_URL=https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${TARBALL}
CHECKSUM=${REPO}_checksums.txt
CHECKSUM_URL=https://github.com/${OWNER}/${REPO}/releases/download/v${VERSION}/${CHECKSUM}

# this function wraps all the destructive operations
# if a curl|bash cuts off the end of the script due to
# network, either nothing will happen or will syntax error
# out preventing half-done work
execute() {
  TMPDIR=$(mktmpdir)
  echo "$PREFIX: downloading ${TARBALL_URL}"
  http_download "${TMPDIR}/${TARBALL}" "${TARBALL_URL}"

  echo "$PREFIX: verifying checksums"
  http_download "${TMPDIR}/${CHECKSUM}" "${CHECKSUM_URL}"
  hash_sha256_verify "${TMPDIR}/${TARBALL}" "${TMPDIR}/${CHECKSUM}"

  (cd "${TMPDIR}" && untar "${TARBALL}")
  install -d "${BINDIR}"
  install "${TMPDIR}/${BINARY}" "${BINDIR}/"
  echo "$PREFIX: installed as ${BINDIR}/${BINARY}"
}

execute`
