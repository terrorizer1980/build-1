// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Command makestatic writes the generated file "static.go".
// It is intended to be invoked via "go generate" (directive in "release.go").
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("releaselet.go")
	if err != nil {
		log.Fatalf("error reading releaselet.go: %v\n", err)
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, `// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by makestatic.go; DO NOT EDIT.

package main

const releaselet = %q
`, b)
	err = ioutil.WriteFile("static.go", buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("error writing static.go: %v\n", err)
	}
}