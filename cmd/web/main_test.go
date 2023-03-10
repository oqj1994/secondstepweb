package main

import "testing"

func TestRun(t *testing.T) {
	db, err := run()
	if db == nil {
		t.Error("failed connect db")
	}
	if err != nil {
		t.Error("failed run()")
	}
}
