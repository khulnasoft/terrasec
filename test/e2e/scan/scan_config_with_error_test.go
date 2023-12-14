package scan_test

import (
	"path/filepath"

	scanUtils "github.com/khulnasoft/terrasec/test/e2e/scan"
	"github.com/khulnasoft/terrasec/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Scan With Config With Error Flag", func() {

	BeforeEach(func() {
		outWriter = gbytes.NewBuffer()
		errWriter = gbytes.NewBuffer()
	})

	AfterEach(func() {
		outWriter = nil
		errWriter = nil
	})

	var iacDir string
	var err error
	iacDir, err = filepath.Abs(filepath.Join(awsIacRelPath, "aws_ami_violation"))

	It("should not error out while getting absolute path", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("scan command is run using the --config-with-error flag for unsupported output types", func() {
		When("output type is human readable format", func() {
			Context("it doesn't support --config-with-error flag", func() {
				Context("human readable output format is the default output format", func() {
					It("should result in an error and exit with status code 1", func() {
						errString := "please use yaml or json output format when using --config-only or --config-with-error flags"
						scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error"}
						scanUtils.RunScanAndAssertErrorMessage(terrasecBinaryPath, helper.ExitCodeOne, scanUtils.ScanTimeout, errString, outWriter, errWriter, scanArgs...)
					})
				})
			})
		})

		When("output type is xml", func() {
			Context("it doesn't support --config-with-error flag", func() {
				It("should result in an error and exit with status code 1", func() {
					errString := "please use yaml or json output format when using --config-only or --config-with-error flags"
					scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "xml"}
					scanUtils.RunScanAndAssertErrorMessage(terrasecBinaryPath, helper.ExitCodeOne, scanUtils.ScanTimeout, errString, outWriter, errWriter, scanArgs...)
				})
			})
		})

		When("output type is junit-xml", func() {
			Context("it doesn't support --config-with-error flag", func() {
				It("should result in an error and exit with status code 1", func() {
					errString := "please use yaml or json output format when using --config-only or --config-with-error flags"
					scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "junit-xml"}
					scanUtils.RunScanAndAssertErrorMessage(terrasecBinaryPath, helper.ExitCodeOne, scanUtils.ScanTimeout, errString, outWriter, errWriter, scanArgs...)
				})
			})
		})
	})

	Describe("scan command is run using the --config-with-error flag for supported output types", func() {
		Context("for terraform files", func() {
			When("output type is json", func() {
				Context("it supports --config-with-error flag", func() {
					It("should display config json and exit with status code 4", func() {
						scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "json"}
						session = helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, scanArgs...)
						Eventually(session, scanUtils.ScanTimeout).Should(gexec.Exit(helper.ExitCodeFour))
					})
				})
			})

			When("output type is yaml", func() {
				Context("it supports --config-with-error flag", func() {
					It("should display config json and exit with status code 4", func() {
						scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "yaml"}
						session = helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, scanArgs...)
						Eventually(session, scanUtils.ScanTimeout).Should(gexec.Exit(helper.ExitCodeFour))
					})
				})
			})
		})

		Context("for yaml files", func() {
			JustBeforeEach(func() {
				iacDir, err = filepath.Abs(filepath.Join(k8sIacRelPath, "kubernetes_ingress_violation"))
			})
			When("output type is json", func() {
				Context("it supports --config-with-error flag", func() {
					It("should display config json and exit with status code 0", func() {
						scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "json", "-i", "k8s"}
						session = helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, scanArgs...)
						Eventually(session, scanUtils.ScanTimeout).Should(gexec.Exit(helper.ExitCodeZero))
					})
				})
			})

			When("output type is yaml", func() {
				Context("it supports --config-with-error flag", func() {
					It("should display config json and exit with status code 0", func() {
						scanArgs := []string{scanUtils.ScanCommand, "-d", iacDir, "--config-with-error", "-o", "yaml", "-i", "k8s"}
						session = helper.RunCommand(terrasecBinaryPath, outWriter, errWriter, scanArgs...)
						Eventually(session, scanUtils.ScanTimeout).Should(gexec.Exit(helper.ExitCodeZero))
					})
				})
			})
		})
	})
})
