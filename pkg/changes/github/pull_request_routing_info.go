package github

type SourceOwner string
type SourceRepository string
type SourceBranch string
type TargetOwner string
type TargetRepository string
type TargetBranch string
type PullRequestTitle string
type PullRequestDescription string
type ChangesetSource string
type ChangesetTarget string

type PullRequestRoutingInfo interface {
	SourceRepositoryOwner() string
	SourceRepository() string
	SourceBranch() string
	TargetRepositoryOwner() string
	TargetRepository() string
	TargetBranch() string
	ChangesetSource() string
	ChangesetTarget() string
}

func NewPullRequestRoutingInfo(
	sourceOwner SourceOwner,
	sourceRepository SourceRepository,
	sourceBranch SourceBranch,
	targetOwner TargetOwner,
	targetRepository TargetRepository,
	targetBranch TargetBranch,
	changesetSource ChangesetSource,
	changesetTarget ChangesetTarget,
) PullRequestRoutingInfo {

	return &routingInfo{
		sourceRepositoryOwner: string(sourceOwner),
		sourceRepository:      string(sourceRepository),
		sourceBranch:          string(sourceBranch),
		targetRepositoryOwner: string(targetOwner),
		targetRepository:      string(targetRepository),
		targetBranch:          string(targetBranch),
		changesetSource:       string(changesetSource),
		changesetTarget:       string(changesetTarget),
	}
}

type routingInfo struct {
	sourceRepositoryOwner string
	sourceRepository      string
	sourceBranch          string
	targetRepositoryOwner string
	targetRepository      string
	targetBranch          string
	changesetSource       string
	changesetTarget       string
}

func (info *routingInfo) SourceRepositoryOwner() string { return info.sourceRepositoryOwner }
func (info *routingInfo) SourceRepository() string      { return info.sourceRepository }
func (info *routingInfo) SourceBranch() string          { return info.sourceBranch }
func (info *routingInfo) TargetRepositoryOwner() string { return info.targetRepositoryOwner }
func (info *routingInfo) TargetRepository() string      { return info.targetRepository }
func (info *routingInfo) TargetBranch() string          { return info.targetBranch }
func (info *routingInfo) ChangesetSource() string       { return info.changesetSource }
func (info *routingInfo) ChangesetTarget() string       { return info.changesetTarget }
