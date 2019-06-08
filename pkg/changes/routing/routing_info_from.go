package routing

import (
	"path/filepath"

	"github.com/calebamiles/keps/pkg/changes/changeset"
	"github.com/calebamiles/keps/pkg/keps"
	"github.com/calebamiles/keps/pkg/orgs"
	"github.com/calebamiles/keps/pkg/settings"
)

func NewInfoFrom(runtime settings.Runtime, kep keps.Instance, org orgs.Instance, changeSet changeset.Description) (Info, error) {
	sourceRepositoryOwner := runtime.PrincipalGithubHandle()
	sourceRepository := org.EnhancementsRepository()
	sourceBranch := org.EnhancementsRepositoryDefaultSourceBranch()

	targetRepositoryOwner := org.EnhancementsRepositoryOwner()
	targetRepository := org.EnhancementsRepository()
	targetBranch := org.EnhancementsRepositoryDefaultTargetBranch()

	committerEmail := runtime.PrincipalEmail()
	committerDisplayName := runtime.PrincipalDisplayName()

	title := changeSet.Title()
	fullDescription := changeSet.FullDescription()
	shortSummary := changeSet.ShortSummary()

	owningSig := kep.OwningSIG()

	localRepositoryLocation := runtime.RepoRoot()

	// TODO test this
	// changes under path must be *relative* to runtime.ContentRoot otherwise go-git will fail to understand that an
	// absolute path may very well be inside the repo ¯\_(ツ)_/¯
	changesUnderPath, err := filepath.Rel(runtime.RepoRoot(), kep.ContentDir())
	if err != nil {
		return nil, err
	}

	cacheDir := runtime.CacheDir()

	changesetSource := org.EnhancementsRepositoryChangesetSource()
	changesetTarget := org.EnhancementsRepositoryChangesetTarget()

	info, err := NewInfo(
		changeSet.Receipt,
		SourceOwner(sourceRepositoryOwner),
		SourceRepository(sourceRepository),
		SourceBranch(sourceBranch),
		TargetOwner(targetRepositoryOwner),
		TargetRepository(targetRepository),
		TargetBranch(targetBranch),
		Title(title),
		FullDescription(fullDescription),
		ShortSummary(shortSummary),
		OwningSIG(owningSig),
		PrincipalName(committerDisplayName),
		PrincipalEmail(committerEmail),
		PathToRepo(localRepositoryLocation),
		PathToChanges(changesUnderPath),
		CacheDir(cacheDir),
		ChangesetSource(changesetSource),
		ChangesetTarget(changesetTarget),
	)

	if err != nil {
		return nil, err
	}

	return info, nil
}
