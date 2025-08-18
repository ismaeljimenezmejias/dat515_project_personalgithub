package main

import (
	"testing"
)

func TestRoot(t *testing.T) {
	testRoot(t, func() {}, func() {})
}

func TestNonExisting(t *testing.T) {
	testNonExisting(t, func() {}, func() {})
}

func TestRedirect(t *testing.T) {
	testRedirect(t, func() {}, func() {})
}

func TestCounter(t *testing.T) {
	testCounter(t, func() {}, func() {})
}

func TestFizzBuzz(t *testing.T) {
	testFizzBuzz(t, func() {}, func() {})
}

func TestServerFull(t *testing.T) {
	testServerFull(t, func() {}, func() {})
}
