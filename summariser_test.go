package main_test

import (
	"os/exec"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

)

var _ = Describe("Summariser", func() {
	// fails if you give invalid section name
	// defaults to something if you don't give section name?
	It("tells you if you give a valid news section argument", func() {
		var session *gexec.Session

		summariserPath, err := gexec.Build("github.com/winnab/summariser")
		Expect(err).NotTo(HaveOccurred())

		sectionName := "foo"
		arg := fmt.Sprintf("-section=%v",sectionName)
		cmd := exec.Command(summariserPath, arg)
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))
		Eventually(session).Should(gbytes.Say(fmt.Sprintf("Yes, you can request %v news.", sectionName)))

		gexec.CleanupBuildArtifacts()
	})


})
