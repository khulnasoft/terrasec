package init

import (
	"io"

	"github.com/khulnasoft/terrasec/test/helper"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"gopkg.in/src-d/go-git.v4"
)

// InitCommandTimeout timeout
const InitCommandTimeout = 60

// RunInitCommand will execute the init command and verify exit code
func RunInitCommand(terrasecBinaryPath string, outWriter, errWriter io.Writer, exitCode int) *gexec.Session {
	session := helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, "init")
	gomega.Eventually(session, InitCommandTimeout).Should(gexec.Exit(exitCode))
	return session
}

// OpenGitRepo checks if a directory is a git repo
func OpenGitRepo(repoPath string) *git.Repository {
	repo, err := git.PlainOpen(repoPath)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(repo).NotTo(gomega.BeNil())
	return repo
}

// ValidateGitRepo validates a git repo and verifies the git url
func ValidateGitRepo(repo *git.Repository, gitURL string) {
	remote, err := repo.Remote("origin")
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(remote).NotTo(gomega.BeNil())
	remoteConfig := remote.Config()
	gomega.Expect(remoteConfig).NotTo(gomega.BeNil())
	err = remoteConfig.Validate()
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(remoteConfig.URLs[0]).To(gomega.BeEquivalentTo(gitURL))
}
