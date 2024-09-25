/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package builder

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TokenFromEnvOrBuilder", func() {
	var appID, appCert string
	var tokenExp uint32 = 3600

	BeforeEach(func() {
		os.Unsetenv("AC_TOKEN")
		appID = "test-app-id"
		appCert = "test-app-cert"
	})

	Context("when both AC_TOKEN and AppID & AppCert are provided", func() {
		BeforeEach(func() {
			os.Setenv("AC_TOKEN", "test-env-token")
		})

		AfterEach(func() {
			// Unset after the test
			os.Unsetenv("AC_TOKEN")
		})

		It("prioritizes the AppID and AppCert for token generation", func() {
			auth, err := NewAuth(appID, appCert, tokenExp)
			Expect(err).NotTo(HaveOccurred())

			token, err := auth.TokenFromEnvOrBuilder()
			Expect(err).NotTo(HaveOccurred())
			Expect(token).ToNot(BeNil())
		})
	})

	Context("when only AC_TOKEN is provided", func() {
		BeforeEach(func() {
			// Set environment token
			os.Setenv("AC_TOKEN", "test-env-token")
			appID = ""
			appCert = ""
		})

		AfterEach(func() {
			os.Unsetenv("AC_TOKEN")
		})

		It("uses the AC_TOKEN from the environment", func() {
			auth, err := NewAuth(appID, appCert, tokenExp)
			Expect(err).NotTo(HaveOccurred())

			token, err := auth.TokenFromEnvOrBuilder()
			Expect(err).NotTo(HaveOccurred())
			Expect(token).To(Equal("test-env-token")) // Should return the env token
		})
	})

	Context("when neither AC_TOKEN nor AppID & AppCert are provided", func() {
		It("returns an error", func() {
			auth, err := NewAuth("", "", tokenExp)
			Expect(err).NotTo(HaveOccurred())

			token, err := auth.TokenFromEnvOrBuilder()
			Expect(err).To(HaveOccurred())
			Expect(token).To(Equal("")) // No token should be available
		})
	})

	Context("when AppID & AppCert are provided but token generation fails", func() {
		It("returns an error from token generation", func() {
			auth, err := NewAuth(appID, appCert, tokenExp)
			Expect(err).NotTo(HaveOccurred())

			_, err = auth.TokenFromEnvOrBuilder()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to generate token"))
		})
	})
})
