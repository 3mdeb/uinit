// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is a basic init script.
package main

import (
	"log"
	"os"
	"os/exec"
)

var (
	commands = [][]string{
		{"/bbin/mount", "-t", "ext4", "/dev/sda1", "/var"},
		{"/bbin/kexec", "-c", `loglvl=all guest_loglvl=all dom0_mem=512M apic_verbosity=debug sched=credit console=com1 no-real-mode`, "--module", `/var/bzImage console=hvc0 earlyprintk=xen nomodeset root=/dev/sda2`, "/var/xen.gz"},
	}
)

func main() {
	for _, line := range commands {
		log.Printf("Executing Command: %v", line)

		cmd := exec.Command(line[0], line[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Print(err)
		}

	}
	log.Print("Uinit Done!")
}
