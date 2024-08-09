[![Go](https://img.shields.io/badge/Go-v1.22.4-blue.svg)](https://go.dev)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-coreutils/etc-hosts-editor)
[![GoReportCard](https://goreportcard.com/badge/github.com/go-coreutils/etc-hosts-editor)](https://goreportcard.com/report/github.com/go-coreutils/etc-hosts-editor)

# etc-hosts-editor

Utility for managing the OS /etc/hosts file.

## INSTALLATION

``` shell
> go install github.com/go-coreutils/etc-hosts-editor/cmd/eheditor@latest
```

## DOCUMENTATION

``` shell
> eheditor --help
NAME:
   eheditor - etc hosts editor

USAGE:
   eheditor [options] [/etc/hosts]

VERSION:
   v0.8.0 (trunk)

DESCRIPTION:
   command line utility for managing the OS /etc/hosts file

GLOBAL OPTIONS:
   --read-only, -r      do not write any changes to the etc hosts file
   --help, -h, --usage  display command-line usage information
   --version, -v        display the version
```


## LICENSE

```
Copyright 2024  The Go-CoreUtils Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
