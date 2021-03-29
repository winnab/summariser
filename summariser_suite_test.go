package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSummariser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Summariser Suite")
}
