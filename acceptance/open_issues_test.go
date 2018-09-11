package acceptance_test

import (
	"bytes"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Open Issues", func() {
	It("returns the list of open issues for a given repository", func() {
		url := "ghc-tdd/spike"

		output := execute(url)

		Expect(output).To(ContainSubstring("#1: List the names of the open issues"))
	})

	// Context("when filtering by help wanted label", func() {
	// 	It("returns the list of open issues that are tagged with help wanted", func() {
	// 		url := "ghc-tdd/spike"
	// 		output := execute(url, "--filter", "help wanted")
	// 		Expect(output).NotTo(ContainSubstring("#1: List the names of the open issues"))
	// 		Expect(output).To(ContainSubstring(`#2: User can see only those issues that have the "help wanted" label`))
	// 	})
	// })

	Context("when filtering by username", func() {
		It("returns the list of open issues that were opened by that user", func() {
			url := "ghc-tdd/spike"
			output := execute(url, "--username", "chinangela")
			Expect(output).To(ContainSubstring("Mock Issue for Username Filter Testing"))
			Expect(output).NotTo(ContainSubstring("#1: List the names of the open issues"))
		})
	})
})

func execute(args ...string) string {
	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})

	cmd := exec.Command(binaryPath, args...)

	session, err := gexec.Start(cmd, stdout, stderr)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session, 1*time.Minute).Should(gexec.Exit(0))

	return strings.TrimSpace(string(stdout.Bytes()))
}
