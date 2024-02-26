//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package prg

import (
	"testing"
)

func TestLoad(t *testing.T) {
	prg, err := Load("hello.prg")
	if err != nil {
		t.Fatal(err)
	}
	err = prg.Print()
	if err != nil {
		t.Error(err)
	}
}
