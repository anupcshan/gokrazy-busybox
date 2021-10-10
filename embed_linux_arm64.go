package main

import _ "embed"

//go:embed bins/busybox.arm64
var bin []byte
