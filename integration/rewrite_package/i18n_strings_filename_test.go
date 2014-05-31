package rewrite_package_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/maximilien/i18n4cf/integration/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rewrite-package -i18n-strings-filename some-file", func() {
	var (
		rootPath          string
		fixturesPath      string
		inputFilesPath    string
		expectedFilesPath string
	)

	Context("input file only contains simple strings", func() {
		BeforeEach(func() {
			dir, err := os.Getwd()
			Ω(err).ShouldNot(HaveOccurred())
			rootPath = filepath.Join(dir, "..", "..")

			fixturesPath = filepath.Join("..", "..", "test_fixtures", "rewrite_package")
			inputFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "input_files")
			expectedFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "expected_output")

			session := Runi18n(
				"-rewrite-package",
				"-f", filepath.Join(inputFilesPath, "test.go"),
				"-o", filepath.Join(rootPath, "tmp"),
				"-i18n-strings-filename", filepath.Join(inputFilesPath, "strings.json"),
				"-v",
			)

			Ω(session.ExitCode()).Should(Equal(0))
		})

		It("rewrites the input file with T() wrappers around the strings specified in the -i18n-strings-filename flag", func() {
			expectedOutputFile := filepath.Join(expectedFilesPath, "test.go")
			bytes, err := ioutil.ReadFile(expectedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			expectedOutput := string(bytes)

			generatedOutputFile := filepath.Join(rootPath, "tmp", "test.go")
			bytes, err = ioutil.ReadFile(generatedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			actualOutput := string(bytes)

			Ω(actualOutput).Should(Equal(expectedOutput))
		})
	})

	Context("input file contains some templated strings", func() {
		BeforeEach(func() {
			dir, err := os.Getwd()
			Ω(err).ShouldNot(HaveOccurred())
			rootPath = filepath.Join(dir, "..", "..")

			fixturesPath = filepath.Join("..", "..", "test_fixtures", "rewrite_package")
			inputFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "input_files")
			expectedFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "expected_output")

			session := Runi18n(
				"-rewrite-package",
				"-f", filepath.Join(inputFilesPath, "test_templated_strings.go"),
				"-o", filepath.Join(rootPath, "tmp"),
				"-i18n-strings-filename", filepath.Join(inputFilesPath, "test_templated_strings.go.en.json"),
				"-v",
			)

			Ω(session.ExitCode()).Should(Equal(0))
		})

		It("rewrites the input file with T() wrappers around the strings (templated and not) specified in the -i18n-strings-filename flag", func() {
			expectedOutputFile := filepath.Join(expectedFilesPath, "test_templated_strings.go")
			bytes, err := ioutil.ReadFile(expectedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			expectedOutput := string(bytes)

			generatedOutputFile := filepath.Join(rootPath, "tmp", "test_templated_strings.go")
			bytes, err = ioutil.ReadFile(generatedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			actualOutput := string(bytes)

			Ω(actualOutput).Should(Equal(expectedOutput))
		})
	})

	Context("input file contains some interpolated strings", func() {
		BeforeEach(func() {
			dir, err := os.Getwd()
			Ω(err).ShouldNot(HaveOccurred())
			rootPath = filepath.Join(dir, "..", "..")

			fixturesPath = filepath.Join("..", "..", "test_fixtures", "rewrite_package")
			inputFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "input_files")
			expectedFilesPath = filepath.Join(fixturesPath, "i18n_strings_filename_option", "expected_output")

			originalJsonData, err := ioutil.ReadFile(filepath.Join(inputFilesPath, "original.test_interpolated_strings.go.en.json"))
			Ω(err).ShouldNot(HaveOccurred())

			err = ioutil.WriteFile(filepath.Join(inputFilesPath, "test_interpolated_strings.go.en.json"), originalJsonData, 0666)
			Ω(err).ShouldNot(HaveOccurred())

			session := Runi18n(
				"-rewrite-package",
				"-f", filepath.Join(inputFilesPath, "test_interpolated_strings.go"),
				"-o", filepath.Join(rootPath, "tmp"),
				"-i18n-strings-filename", filepath.Join(inputFilesPath, "test_interpolated_strings.go.en.json"),
				"-v",
			)

			Ω(session.ExitCode()).Should(Equal(0))
		})

		It("converts interpolated strings to templated and rewrites the input file with T() wrappers around the strings (templated and not) specified in the -i18n-strings-filename flag", func() {
			expectedOutputFile := filepath.Join(expectedFilesPath, "test_interpolated_strings.go")
			bytes, err := ioutil.ReadFile(expectedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			expectedOutput := string(bytes)

			generatedOutputFile := filepath.Join(rootPath, "tmp", "test_interpolated_strings.go")
			bytes, err = ioutil.ReadFile(generatedOutputFile)
			Ω(err).ShouldNot(HaveOccurred())

			actualOutput := string(bytes)

			Ω(actualOutput).Should(Equal(expectedOutput))
		})

		It("updates the i18n strings JSON file with the converted interpolated JSON strings", func() {
			expectedOutputFile := filepath.Join(expectedFilesPath, "test_interpolated_strings.go.en.json")
			generatedOutputFile := filepath.Join(inputFilesPath, "test_interpolated_strings.go.en.json")
			CompareExpectedToGeneratedTraslationJson(expectedOutputFile, generatedOutputFile)
		})
	})
})