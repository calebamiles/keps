package routing

type Info interface {
	ChangeReceipt() string

	SourceRepositoryOwner() string
	SourceRepository() string
	SourceBranch() string

	TargetRepositoryOwner() string
	TargetRepository() string
	TargetBranch() string

	PathToRepo() string
	ChangesUnderPath() string
	CacheDir() string

	Title() string           // used for pull request title
	FullDescription() string // used for pull request description
	ShortSummary() string    // used as commit message

	OwningSIG() string
	PrincipalName() string
	PrincipalEmail() string

	ChangesetSource() string
	ChangesetTarget() string
}

type SourceOwner string
type SourceRepository string
type SourceBranch string
type TargetOwner string
type TargetRepository string
type TargetBranch string
type Title string
type FullDescription string
type ShortSummary string
type OwningSIG string
type PrincipalName string
type PrincipalEmail string
type PathToRepo string
type PathToChanges string
type CacheDir string
type ChangesetSource string
type ChangesetTarget string

func NewInfo(
	changeReceipt func() string,
	sourceRepositoryOwner SourceOwner,
	sourceRepository SourceRepository,
	sourceBranch SourceBranch,
	targetRepositoryOwner TargetOwner,
	targetRepository TargetRepository,
	targetBranch TargetBranch,
	title Title,
	fullDescription FullDescription,
	shortSummary ShortSummary,
	owningSIG OwningSIG,
	principalName PrincipalName,
	principalEmail PrincipalEmail,
	pathToRepo PathToRepo,
	pathToChanges PathToChanges,
	cacheDir CacheDir,
	changesetSource ChangesetSource,
	changesetTarget ChangesetTarget,
) (Info, error) {
	info := &routingInfo{
		changeReceipt:         changeReceipt,
		sourceRepositoryOwner: string(sourceRepositoryOwner),
		sourceRepository:      string(sourceRepository),
		sourceBranch:          string(sourceBranch),
		targetRepositoryOwner: string(targetRepositoryOwner),
		targetRepository:      string(targetRepository),
		targetBranch:          string(targetBranch),
		title:                 string(title),
		fullDescription:       string(fullDescription),
		shortSummary:          string(shortSummary),
		owningSIG:             string(owningSIG),
		principalName:         string(principalName),
		principalEmail:        string(principalEmail),
		pathToRepo:            string(pathToRepo),
		pathToChanges:         string(pathToChanges),
		cacheDir:              string(cacheDir),
		changesetSource:       string(changesetSource),
		changesetTarget:       string(changesetTarget),
	}

	return info, nil
}

type routingInfo struct {
	changeReceipt func() string

	sourceRepositoryOwner string
	sourceRepository      string
	sourceBranch          string
	targetRepositoryOwner string
	targetRepository      string
	targetBranch          string

	title           string
	fullDescription string
	shortSummary    string
	owningSIG       string
	principalName   string
	principalEmail  string

	pathToRepo    string
	pathToChanges string
	cacheDir      string

	changesetSource string
	changesetTarget string
}

func (i *routingInfo) ChangeReceipt() string         { return i.changeReceipt() }
func (i *routingInfo) SourceRepositoryOwner() string { return i.sourceRepositoryOwner }
func (i *routingInfo) SourceRepository() string      { return i.sourceRepository }
func (i *routingInfo) SourceBranch() string          { return i.sourceBranch }
func (i *routingInfo) TargetRepositoryOwner() string { return i.targetRepositoryOwner }
func (i *routingInfo) TargetRepository() string      { return i.targetRepository }
func (i *routingInfo) TargetBranch() string          { return i.targetBranch }
func (i *routingInfo) PathToRepo() string            { return i.pathToRepo }
func (i *routingInfo) ChangesUnderPath() string      { return i.pathToChanges }
func (i *routingInfo) Title() string                 { return i.title }
func (i *routingInfo) FullDescription() string       { return i.fullDescription }
func (i *routingInfo) ShortSummary() string          { return i.shortSummary }
func (i *routingInfo) OwningSIG() string             { return i.owningSIG }
func (i *routingInfo) PrincipalName() string         { return i.principalName }
func (i *routingInfo) PrincipalEmail() string        { return i.principalEmail }
func (i *routingInfo) CacheDir() string              { return i.cacheDir }
func (i *routingInfo) ChangesetSource() string       { return i.changesetSource }
func (i *routingInfo) ChangesetTarget() string       { return i.changesetTarget }
