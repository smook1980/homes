package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHomes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Homes Suite")
}
