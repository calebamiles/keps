package settings_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/thanhpk/randstr"
	"sigs.k8s.io/yaml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/calebamiles/keps/pkg/settings"
	"github.com/calebamiles/keps/pkg/settings/settingsfakes"
)

var _ = FDescribe("working with runtime settings", func() {
	Describe("building runtime settings", func() {
		Context("from a disk location specified by environment variable", func() {
			It("returns runtime settings", func() {
				var kepConfig struct {
					ContentRoot  string `json:"content_root"`
					TokenPath    string `json:"token_path"`
					CacheDir     string `json:"cache_dir"`
					ContactEmail string `json:"contact_email"`
					DisplayName  string `json:"display_name"`
					GithubHandle string `json:"github_handle"`
				}

				configDir, err := ioutil.TempDir("", "kep-settings-test")
				Expect(err).ToNot(HaveOccurred(), "expected no error creating a temporary directory to store config file")
				defer os.RemoveAll(configDir)

				// will be removed when directory is removed
				configFile, err := os.Create(filepath.Join(configDir, settings.ConfigName))
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating an empty config file")

				targetDir := randstr.String(16)

				kepConfig.ContentRoot = randstr.String(16)
				kepConfig.TokenPath = randstr.String(16)
				kepConfig.CacheDir = randstr.String(16)
				kepConfig.ContactEmail = randstr.String(16)
				kepConfig.DisplayName = randstr.String(16)
				kepConfig.GithubHandle = randstr.String(16)

				configBytes, err := yaml.Marshal(kepConfig)
				Expect(err).ToNot(HaveOccurred(), "expected no error when marshalling config struct to YAML")

				_, err = configFile.Write(configBytes)
				Expect(err).ToNot(HaveOccurred(), "expected no error when writing YAML config to disk")
				err = configFile.Close()
				Expect(err).ToNot(HaveOccurred(), "expected no error when closing config file after write")

				err = os.Setenv("KEPCONFIG", configFile.Name())
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEPCONFIG environment variable to point at written config")
				defer os.Unsetenv("KEPCONFIG")

				readConfig, err := settings.Open("", targetDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error when opening settings when KEPCONFIG environment variable is set to an existing config")

				Expect(readConfig.TargetDir()).To(Equal(targetDir), "expected read `target_dir` to match written `target_dir`")
				Expect(readConfig.ContentRoot()).To(Equal(kepConfig.ContentRoot), "expected read `content_root` to match written `content_root`")
				Expect(readConfig.TokenPath()).To(Equal(kepConfig.TokenPath), "expected read `token_path` to match written `token_path`")
				Expect(readConfig.CacheDir()).To(Equal(kepConfig.CacheDir), "expected read `cache_dir` to match written `cache_dir`")
				Expect(readConfig.PrincipalEmail()).To(Equal(kepConfig.ContactEmail), "expected read `contact_email` to match written `contact_email`")
				Expect(readConfig.PrincipalGithubHandle()).To(Equal(kepConfig.GithubHandle), "expected read `github_handle` to match written `github_handle`")
				Expect(readConfig.PrincipalDisplayName()).To(Equal(kepConfig.DisplayName), "expeted read `display_name` to match written `display_name`")
			})
		})

		Context("from a disk location passed into settings.Open()", func() {
			It("returns runtime settings", func() {
				var kepConfig struct {
					ContentRoot  string `json:"content_root"`
					TokenPath    string `json:"token_path"`
					CacheDir     string `json:"cache_dir"`
					ContactEmail string `json:"contact_email"`
					DisplayName  string `json:"display_name"`
					GithubHandle string `json:"github_handle"`
				}

				configDir, err := ioutil.TempDir("", "kep-settings-test")
				Expect(err).ToNot(HaveOccurred(), "expected no error creating a temporary directory to store config file")
				defer os.RemoveAll(configDir)

				// will be removed when directory is removed
				configFile, err := os.Create(filepath.Join(configDir, settings.ConfigName))
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating an empty config file")

				targetDir := randstr.String(16)

				kepConfig.ContentRoot = randstr.String(16)
				kepConfig.TokenPath = randstr.String(16)
				kepConfig.CacheDir = randstr.String(16)
				kepConfig.ContactEmail = randstr.String(16)
				kepConfig.DisplayName = randstr.String(16)
				kepConfig.GithubHandle = randstr.String(16)

				configBytes, err := yaml.Marshal(kepConfig)
				Expect(err).ToNot(HaveOccurred(), "expected no error when marshalling config struct to YAML")

				_, err = configFile.Write(configBytes)
				Expect(err).ToNot(HaveOccurred(), "expected no error when writing YAML config to disk")
				err = configFile.Close()
				Expect(err).ToNot(HaveOccurred(), "expected no error when closing config file after write")

				readConfig, err := settings.Open(configFile.Name(), targetDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error when opening settings when KEPCONFIG environment variable is set to an existing config")

				Expect(readConfig.TargetDir()).To(Equal(targetDir), "expected read `target_dir` to match written `target_dir`")
				Expect(readConfig.ContentRoot()).To(Equal(kepConfig.ContentRoot), "expected read `content_root` to match written `content_root`")
				Expect(readConfig.TokenPath()).To(Equal(kepConfig.TokenPath), "expected read `token_path` to match written `token_path`")
				Expect(readConfig.CacheDir()).To(Equal(kepConfig.CacheDir), "expected read `cache_dir` to match written `cache_dir`")
				Expect(readConfig.PrincipalEmail()).To(Equal(kepConfig.ContactEmail), "expected read `contact_email` to match written `contact_email`")
				Expect(readConfig.PrincipalGithubHandle()).To(Equal(kepConfig.GithubHandle), "expected read `github_handle` to match written `github_handle`")
				Expect(readConfig.PrincipalDisplayName()).To(Equal(kepConfig.DisplayName), "expeted read `display_name` to match written `display_name`")
			})
		})

		Context("from a `standard` location such as the working directory", func() {
			It("returns runtime settings", func() {
				var kepConfig struct {
					ContentRoot  string `json:"content_root"`
					TokenPath    string `json:"token_path"`
					CacheDir     string `json:"cache_dir"`
					ContactEmail string `json:"contact_email"`
					DisplayName  string `json:"display_name"`
					GithubHandle string `json:"github_handle"`
				}

				configDir, err := ioutil.TempDir("", "kep-settings-test")
				Expect(err).ToNot(HaveOccurred(), "expected no error creating a temporary directory to store config file")
				defer os.RemoveAll(configDir)

				pwd, err := os.Getwd()
				Expect(err).ToNot(HaveOccurred(), "expected no error getting current directory")
				defer os.Chdir(pwd)

				err = os.Chdir(configDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error changing directory to location of config")

				// will be removed when directory is removed
				configFile, err := os.Create(filepath.Join(configDir, settings.ConfigName))
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating an empty config file")

				targetDir := randstr.String(16)

				kepConfig.ContentRoot = randstr.String(16)
				kepConfig.TokenPath = randstr.String(16)
				kepConfig.CacheDir = randstr.String(16)
				kepConfig.ContactEmail = randstr.String(16)
				kepConfig.DisplayName = randstr.String(16)
				kepConfig.GithubHandle = randstr.String(16)

				configBytes, err := yaml.Marshal(kepConfig)
				Expect(err).ToNot(HaveOccurred(), "expected no error when marshalling config struct to YAML")

				_, err = configFile.Write(configBytes)
				Expect(err).ToNot(HaveOccurred(), "expected no error when writing YAML config to disk")
				err = configFile.Close()
				Expect(err).ToNot(HaveOccurred(), "expected no error when closing config file after write")

				readConfig, err := settings.Open(configFile.Name(), targetDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error when opening settings when KEPCONFIG environment variable is set to an existing config")

				Expect(readConfig.TargetDir()).To(Equal(targetDir), "expected read `target_dir` to match written `target_dir`")
				Expect(readConfig.ContentRoot()).To(Equal(kepConfig.ContentRoot), "expected read `content_root` to match written `content_root`")
				Expect(readConfig.TokenPath()).To(Equal(kepConfig.TokenPath), "expected read `token_path` to match written `token_path`")
				Expect(readConfig.CacheDir()).To(Equal(kepConfig.CacheDir), "expected read `cache_dir` to match written `cache_dir`")
				Expect(readConfig.PrincipalEmail()).To(Equal(kepConfig.ContactEmail), "expected read `contact_email` to match written `contact_email`")
				Expect(readConfig.PrincipalGithubHandle()).To(Equal(kepConfig.GithubHandle), "expected read `github_handle` to match written `github_handle`")
				Expect(readConfig.PrincipalDisplayName()).To(Equal(kepConfig.DisplayName), "expeted read `display_name` to match written `display_name`")
			})
		})

		Context("from environment variables", func() {
			It("returns runtime settings", func() {
				const (
					principalEmail        = "kubernetes-sig-architecture+keps@googlegroups.com"
					principalGithubHandle = "ossKEPTool"
					principalDisplayName  = "The OSS KEP Tool"
				)

				By("setting environment variables")

				err := os.Setenv("KEP_GITHUB_HANDLE", principalGithubHandle)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_GITHUB_HANDLE environment variable")
				defer os.Unsetenv("KEP_GITHUB_HANDLE")

				err = os.Setenv("KEP_CONTACT_EMAIL", principalEmail)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_EMAIL_ADDRESS environment variable")
				defer os.Unsetenv("KEP_EMAIL_ADDRESS")

				err = os.Setenv("KEP_DISPLAY_NAME", principalDisplayName)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_DISPLAY_NAME environment variable")
				defer os.Unsetenv("KEP_DISPLAY_NAME")

				targetDir, err := ioutil.TempDir("", "kep-runtime-settings-test-target-dir")
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp directory")
				defer os.RemoveAll(targetDir)

				contentRoot, err := ioutil.TempDir("", "kep-runtime-settings-test-content-root")
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp directory")
				defer os.RemoveAll(contentRoot)

				err = os.Setenv("KEP_CONTENT_ROOT", contentRoot)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_CONTENT_ROOT environment variable")
				defer os.Unsetenv("KEP_CONTENT_ROOT")

				cacheDir, err := ioutil.TempDir("", "kep-runtime-settings-test-cache-dir")
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp directory")
				defer os.RemoveAll(cacheDir)

				err = os.Setenv("KEP_CACHE_DIR", cacheDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_CACHE_DIR environment variable")
				defer os.Unsetenv("KEP_CACHE_DIR")

				tokenPath, err := ioutil.TempDir("", "kep-runtime-settings-test-token-path")
				Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp directory")
				defer os.RemoveAll(tokenPath)

				err = os.Setenv("KEP_TOKEN_PATH", tokenPath)
				Expect(err).ToNot(HaveOccurred(), "expected no error when setting KEP_TOKEN_PATH environment variable")
				defer os.Unsetenv("KEP_TOKEN_PATH")

				runtime, err := settings.FromEnvironment(targetDir)
				Expect(err).ToNot(HaveOccurred(), "expected no error when reading runtime settings from the environment")

				By("returning the value of KEP_* environment variables")

				Expect(runtime.TargetDir()).To(Equal(targetDir), "expected #TargetDir() to return given targetDir")
				Expect(runtime.ContentRoot()).To(Equal(contentRoot), "expected #ContentRoot() to return the value from the environment")
				Expect(runtime.TokenPath()).To(Equal(tokenPath), "expected #TokenPath() to return the value from the environment")
				Expect(runtime.CacheDir()).To(Equal(cacheDir), "expected #CacheDir() to return the value from the environment")
				Expect(runtime.PrincipalEmail()).To(Equal(principalEmail), "expected #PrincipalEmail() to return the value from the environment")
				Expect(runtime.PrincipalGithubHandle()).To(Equal(principalGithubHandle), "expected #PrincipalGithubHandle() to return the value from the environment")
				Expect(runtime.PrincipalDisplayName()).To(Equal(principalDisplayName), "expected #PrincipalDisplayName() to return the value from the environment")
			})
		})

	})

	Describe("validating runtime settings", func() {
		Context("when a required value has not been set", func() {
			It("returns all validation errors", func() {
				emptyRuntime := &settingsfakes.FakeRuntime{}
				err := settings.Validate(emptyRuntime)
				Expect(err).To(HaveOccurred(), "expected validation errors when attempting to validate a settings.Runtime which returns empty data")

				merr, ok := err.(*multierror.Error)
				Expect(ok).To(BeTrue(), "expected it to be ok to cast an error returned from settings.Validate to a multierror.Error")

				Expect(merr.Errors).To(HaveLen(4), "expected 4 validation errors when attempting to validate a settings.Runtime with returns empty data")
				Expect(merr.Errors).To(ContainElement(errors.New("no display name set")))
				Expect(merr.Errors).To(ContainElement(errors.New("no GitHub token path set")))
				Expect(merr.Errors).To(ContainElement(errors.New("no GitHub handle set")))
				Expect(merr.Errors).To(ContainElement(errors.New("no contact email set")))
			})
		})

	})
})
