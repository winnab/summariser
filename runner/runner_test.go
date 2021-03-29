package runner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/winnab/summariser/runner"
	"io/ioutil"
)

var _ = Describe("Runner", func() {
	It("succeeds", func() {
		Expect(runner.Run()).To(Succeed())
		//	runner.Run() should happen before each test
		// pass in name of file with random num on the end
	})

	It("creates a file with news content", func() {
		dat, err := ioutil.ReadFile("/tmp/news")
		Expect(err).NotTo(HaveOccurred())
		Expect(string(dat)).To(ContainSubstring("editions"))
		//	delete file after each
	})
})
