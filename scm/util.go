// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scm

import "strings"

// Split splits the full repository name into segments.
func Split(s string) (owner, name string) {
	parts := strings.SplitN(s, "/", 2)
	switch len(parts) {
	case 1:
		name = parts[0]
	case 2:
		owner = parts[0]
		name = parts[1]
	}
	return
}

// Join joins the repository owner and name segments to
// create a fully qualified repository name.
func Join(owner, name string) string {
	return owner + "/" + name
}
