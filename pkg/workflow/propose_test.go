package workflow_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"

	"github.com/calebamiles/keps/pkg/changes/auth"
	"github.com/calebamiles/keps/pkg/keps"
	"github.com/calebamiles/keps/pkg/orgs/sandbox"
	"github.com/calebamiles/keps/pkg/settings"
	"github.com/calebamiles/keps/pkg/workflow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Proposing a KEP", func() {
	const (
		authorOne         = "handleOne"
		title             = "A Great but Complicated Idea"
		kubernetesWideDir = "kubernetes-wide"
		metadataFilename  = "metadata.yaml"
	)

	Context("with a valid KEP", func() {
		FIt("proposes a KEP and returns a link to the GitHub Pull Request", func() {
			By("reading in the runtime settings")

			settingsLocation := os.Getenv("KEP_TEST_CONFIG")
			if settingsLocation == "" {
				Skip("`KEP_TEST_CONFIG` environment variable is unset and required for test")
			}

			// we need to peek at the settings to get the content root before creating the settings.Runtime
			var settingsHint struct {
				ContentRoot string `json:"content_root"`
			}

			settingsBytes, err := ioutil.ReadFile(settingsLocation)
			Expect(err).ToNot(HaveOccurred(), "expected no error reading in settings file to determine content root")

			err = yaml.Unmarshal(settingsBytes, &settingsHint)
			Expect(err).ToNot(HaveOccurred(), "expected no error when unmarshalling settings bytes to YAML to determine content root")

			// create a "kubernetes-wide" KEP
			kepDirName := "a-great-and-complex-idea"
			targetDir := filepath.Join("content", kepDirName)

			runtime, err := settings.Open(settingsLocation, targetDir)
			Expect(err).ToNot(HaveOccurred(), "expeted no error when reading settings file from location specified by environment variable")

			By("creating a user sandbox")

			sandbox, err := sandbox.New(runtime)
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a new user sandbox")

			token, err := auth.NewProvideTokenFromPath(runtime.TokenPath())
			Expect(err).ToNot(HaveOccurred(), "expected no error creating a new token provider using path from settings.Runtime")

			err = sandbox.CreateRepositories(token)
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating sandbox GitHub repositories in a user account")

			By("creating a KEP skeleton")

			kepLocation, err := workflow.Init(runtime)
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a new KEP skeleton")

			By("returning the URL that a human would care about")

			kep, err := keps.Open(kepLocation)
			Expect(err).ToNot(HaveOccurred(), "expected no error when opening the KEP created by workflow.Init()")

			prUrl, err := workflow.Propose(runtime, sandbox, kep)
			Expect(err).ToNot(HaveOccurred(), "expected no error when proposing a KEP created with workflow.Init()")

			var getPRResponse struct {
				Title string `json:"title"`
				Body  string `json:"body"`
			}

			Expect(prUrl).To(BeEmpty())
		})
	})

	Context("when the KEP has already been proposed", func() {
		It("returns an existing link to the GitHub Pull Request", func() {
			Fail("test not written")

			By("reading in the runtime settings")

			By("returning the URL that a human would care about")
		})
	})
})
