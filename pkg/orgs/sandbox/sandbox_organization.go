package sandbox

import (
	"github.com/hashicorp/go-multierror"

	"github.com/calebamiles/keps/pkg/changes/auth"
	"github.com/calebamiles/keps/pkg/changes/github"
	"github.com/calebamiles/keps/pkg/orgs"
	"github.com/calebamiles/keps/pkg/settings"
)

const (
	upstreamOwner                 = "Planctae"
	sandboxEnhancementsRepository = "enhancements"
	sandboxTrackingRepository     = "enhancements-tracking"
	sandboxApiReivewRepository    = "api-review"
	sandboxDefaultSourceBranch    = "master"
	sandboxDefaultTargetBranch    = "sandbox"
)

type Instance interface {
	orgs.Instance
	CreateRepositories(token auth.TokenProvider) error
}

func New(r settings.Runtime) (Instance, error) {
	owner := r.PrincipalGithubHandle()

	token, err := auth.NewProvideTokenFromPath(r.TokenPath())
	if err != nil {
		return nil, err
	}

	return &sandbox{owner: owner, token: token}, nil
}

type sandbox struct {
	owner string
	token auth.TokenProvider
}

// KEP content
func (s *sandbox) EnhancementsRepository() string      { return sandboxEnhancementsRepository }
func (s *sandbox) EnhancementsRepositoryOwner() string { return s.owner }
func (s *sandbox) EnhancementsRepositoryDefaultSourceBranch() string {
	return sandboxDefaultSourceBranch
}
func (s *sandbox) EnhancementsRepositoryDefaultTargetBranch() string {
	return sandboxDefaultTargetBranch
}
func (s *sandbox) EnhancementsRepositoryChangesetSource() string { return sandboxDefaultTargetBranch }
func (s *sandbox) EnhancementsRepositoryChangesetTarget() string { return sandboxDefaultSourceBranch }

// release tracking
func (s *sandbox) EnhancementsTrackingRepository() string      { return sandboxTrackingRepository }
func (s *sandbox) EnhancementsTrackingRepositoryOwner() string { return s.owner }
func (s *sandbox) EnhancementsTrackingRepositoryDefaultSourceBranch() string {
	return sandboxDefaultSourceBranch
}
func (s *sandbox) EnhancementsTrackingRepositoryDefaultTargetBranch() string {
	return sandboxDefaultTargetBranch
}
func (s *sandbox) EnhancementsTrackingRepositoryChangesetSource() string {
	return sandboxDefaultTargetBranch
}
func (s *sandbox) EnhancementsTrackingRepositoryChangesetTarget() string {
	return sandboxDefaultSourceBranch
}

// API review

func (s *sandbox) ApiReviewRepository() string                    { return sandboxApiReivewRepository }
func (s *sandbox) ApiReviewRepositoryOwner() string               { return s.owner }
func (s *sandbox) ApiReviewRepositoryDefaultSourceBranch() string { return sandboxDefaultSourceBranch }
func (s *sandbox) ApiReviewRepositoryDefaultTargetBranch() string { return sandboxDefaultTargetBranch }
func (s *sandbox) ApiReviewRepositoryChangesetSource() string     { return sandboxDefaultTargetBranch }
func (s *sandbox) ApiReviewRepositoryChangesetTarget() string     { return sandboxDefaultSourceBranch }

// authz
func (s *sandbox) IsAuthorized(_ settings.Runtime) (bool, error) {

	// the user is always allowed to use the sandbox

	return true, nil
}

func (s *sandbox) CreateRepositories(token auth.TokenProvider) error {
	var errs *multierror.Error

	_, err := github.Fork(s.token, upstreamOwner, sandboxEnhancementsRepository)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	_, err = github.Fork(s.token, upstreamOwner, sandboxTrackingRepository)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	_, err = github.Fork(s.token, upstreamOwner, sandboxApiReivewRepository)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	return errs.ErrorOrNil()
}
