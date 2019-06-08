package orgs

import (
	"github.com/calebamiles/keps/pkg/settings"
)

type Instance interface {
	// where KEP content lives
	EnhancementsRepositoryOwner() string
	EnhancementsRepository() string
	EnhancementsRepositoryDefaultSourceBranch() string
	EnhancementsRepositoryDefaultTargetBranch() string
	EnhancementsRepositoryChangesetSource() string
	EnhancementsRepositoryChangesetTarget() string

	// where release tracking happens
	EnhancementsTrackingRepositoryOwner() string
	EnhancementsTrackingRepository() string
	EnhancementsTrackingRepositoryDefaultSourceBranch() string
	EnhancementsTrackingRepositoryDefaultTargetBranch() string
	EnhancementsTrackingRepositoryChangesetSource() string
	EnhancementsTrackingRepositoryChangesetTarget() string

	// where API review happens
	ApiReviewRepositoryOwner() string
	ApiReviewRepository() string
	ApiReviewRepositoryDefaultSourceBranch() string
	ApiReviewRepositoryDefaultTargetBranch() string
	ApiReviewRepositoryChangesetSource() string
	ApiReviewRepositoryChangesetTarget() string

	// security
	IsAuthorized(runtime settings.Runtime) (bool, error)
}
