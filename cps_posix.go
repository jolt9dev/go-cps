//go:build aix || darwin || dragonfly || freebsd || hurd || illumos || ios || linux || netbsd || openbsd || plan9 || solaris || zos

package cps

import "os"

// IsElevated returns true if the process is running with elevated privileges.
func IsElevated() bool {
	return os.Geteuid() == 0
}
