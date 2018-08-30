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
})

func execute(url string) string {
	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})

	cmd := exec.Command(binaryPath, url)

	session, err := gexec.Start(cmd, stdout, stderr)
	Expect(err).NotTo(HaveOccurred())
	Eventually(session, 1*time.Minute).Should(gexec.Exit(0))

	return strings.TrimSpace(string(stdout.Bytes()))
}
