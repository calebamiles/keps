package orgs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrgs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Orgs Suite")
}
