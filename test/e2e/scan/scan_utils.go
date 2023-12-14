

package scan

import (
	"io"
	"path/filepath"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/khulnasoft/terrasec/test/helper"
)

const (
	// ScanCommand is terrasec's scan command
	ScanCommand string = "scan"

	// ScanTimeout is default scan command execution timeout
	ScanTimeout int = 3

	// WebhookScanTimeout is default scan command webhook execution timeout
	WebhookScanTimeout int = 30

	// RemoteScanTimeout is default scan command remote execution timeout
	RemoteScanTimeout int = 30
)

// RunScanAndAssertGoldenOutputRegexWithTimeout runs the scan command with supplied paramters and compares actual and golden output
// it replaces variable parts in output with regex eg: timestamp, file path
// added to provide extra option for specifying timeout.
func RunScanAndAssertGoldenOutputRegexWithTimeout(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isJunitXML, isStdOut bool, outWriter, errWriter io.Writer, scanTimeout int, args ...string) {
	session, goldenFileAbsPath := RunScanCommandWithTimeOut(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, scanTimeout, args...)
	helper.CompareActualWithGoldenSummaryRegex(session, goldenFileAbsPath, isJunitXML, isStdOut)
}

// RunScanAndAssertGoldenOutputRegex runs the scan command with supplied parameters and compares actual and golden output
// it replaces variable parts in output with regex eg: timestamp, file path
func RunScanAndAssertGoldenOutputRegex(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isJunitXML, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualWithGoldenSummaryRegex(session, goldenFileAbsPath, isJunitXML, isStdOut)
}

// RunScanAndAssertGoldenOutput runs the scan command with supplied parameters and compares actual and golden output
func RunScanAndAssertGoldenOutput(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualWithGolden(session, goldenFileAbsPath, isStdOut)
}

// RunScanCommandWithTimeOut executes the scan command, validates exit code
// added to provide extra option for specifying timeout.
func RunScanCommandWithTimeOut(terrasecBinaryPath, relGoldenFilePath string, exitCode int, outWriter, errWriter io.Writer, scanTimeout int, args ...string) (*gexec.Session, string) {
	argList := []string{ScanCommand}
	argList = append(argList, args...)
	session := helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, argList...)
	gomega.Eventually(session, scanTimeout).Should(gexec.Exit(exitCode))
	goldenFileAbsPath, err := filepath.Abs(relGoldenFilePath)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	return session, goldenFileAbsPath
}

// RunScanAndAssertJSONOutput runs the scan command with supplied parameters and compares actual and golden output
func RunScanAndAssertJSONOutput(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isJunitXML, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualWithGoldenJSON(session, goldenFileAbsPath, isStdOut)
}

// RunScanAndAssertJSONOutputString runs the scan command with supplied parameters and compares actual and golden output
func RunScanAndAssertJSONOutputString(terrasecBinaryPath, goldenString string, exitCode int, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	argList := []string{ScanCommand}
	argList = append(argList, args...)
	session := helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, argList...)
	gomega.Eventually(session, ScanTimeout).Should(gexec.Exit(exitCode))
	helper.CompareActualWithGoldenJSONString(session, goldenString, isStdOut)
}

// RunScanAndAssertYAMLOutput runs the scan command with supplied parameters and compares actual and golden output
func RunScanAndAssertYAMLOutput(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isJunitXML, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualWithGoldenYAML(session, goldenFileAbsPath, isStdOut)
}

// RunScanAndAssertXMLOutput runs the scan command with supplied parameters and compares actual and golden output
func RunScanAndAssertXMLOutput(terrasecBinaryPath, relGoldenFilePath string, exitCode int, isJunitXML, isStdOut bool, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualWithGoldenXML(session, goldenFileAbsPath, isStdOut)
}

// RunScanAndAssertErrorMessage runs the scan command with supplied parameters and checks of error string is present
func RunScanAndAssertErrorMessage(terrasecBinaryPath string, exitCode, timeOut int, errString string, outWriter, errWriter io.Writer, args ...string) {
	session := helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, args...)
	gomega.Eventually(session, timeOut).Should(gexec.Exit(exitCode))
	helper.ContainsErrorSubString(session, errString)
}

// RunScanCommand executes the scan command, validates exit code
func RunScanCommand(terrasecBinaryPath, relGoldenFilePath string, exitCode int, outWriter, errWriter io.Writer, args ...string) (*gexec.Session, string) {
	argList := []string{ScanCommand}
	argList = append(argList, args...)
	session := helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, argList...)
	gomega.Eventually(session, ScanTimeout).Should(gexec.Exit(exitCode))
	goldenFileAbsPath, err := filepath.Abs(relGoldenFilePath)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	return session, goldenFileAbsPath
}

// RunScanAndAssertGoldenSarifOutputRegex runs the scan command with supplied parameters and compares actual and golden output
// it replaces variable parts in output with regex eg: uri, version path
func RunScanAndAssertGoldenSarifOutputRegex(terrasecBinaryPath, relGoldenFilePath string, exitCode int, outWriter, errWriter io.Writer, args ...string) {
	session, goldenFileAbsPath := RunScanCommand(terrasecBinaryPath, relGoldenFilePath, exitCode, outWriter, errWriter, args...)
	helper.CompareActualSarifOutputWithGoldenSummaryRegex(session, goldenFileAbsPath)
}
