// Largely derived from
// https://github.com/gokrazy/serial-busybox/blob/b4f0c5c3f3aa2b12a43beb23d40647428259d851/wrapper.go

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gokrazy/gokrazy"
)

func main() {
	// Path hardcoded in
	// https://github.com/gokrazy/gokrazy/blob/54d80ad29da34ed8acc97fad18739b9da2ba1d57/gokrazy.go#L276
	const wellKnownSerialShell = "/tmp/serial-busybox/ash"
	if err := os.MkdirAll(filepath.Dir(wellKnownSerialShell), 0755); err != nil {
		log.Fatalf("Mkdir: %v", err)
	}
	err := ioutil.WriteFile(wellKnownSerialShell, bin, 0755)
	if err != nil {
		log.Fatalf("could not write busybox: %v", err)
	}
	gokrazy.DontStartOnBoot()
}
