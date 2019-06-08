package settings

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/shibukawa/configdir"
	"github.com/spf13/viper"
)

const (
	ConfigName            = "config"
	configContentRootKey  = "content_root"
	configRepoRootKey     = "repo_root"
	configGithubHandleKey = "github_handle"
	configEmailKey        = "contact_email"
	configCacheDir        = "cache_dir"
	configDisplayNameKey  = "display_name"
	configTokenPathKey    = "token_path"
	envPrefix             = "KEP"
	envConfigPath         = "KEPCONFIG"

	requiredContentRootName = "content"
)

type Runtime interface {
	TargetDir() string
	ContentRoot() string
	RepoRoot() string
	TokenPath() string
	CacheDir() string
	PrincipalEmail() string
	PrincipalGithubHandle() string
	PrincipalDisplayName() string
}

func FromEnvironment(targetDir string) (Runtime, error) {
	v := newViper()

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	return newRuntimeFromViper(v, targetDir)
}

func Open(configPath string, targetDir string) (Runtime, error) {
	v := newViper()

	// check given path first
	switch configPath {
	case "":
		// do nothing
	default:
		f, innerErr := os.Open(configPath)
		if innerErr != nil {
			return nil, innerErr
		}

		innerErr = v.ReadConfig(f)
		if innerErr != nil {
			return nil, innerErr
		}

		innerErr = f.Close()
		if innerErr != nil {
			return nil, innerErr
		}

		return newRuntimeFromViper(v, targetDir)
	}

	// check if config environment variable has been set
	pathFromEnv := os.Getenv(envConfigPath)
	switch pathFromEnv {
	case "":
		// do nothing
	default:
		f, innerErr := os.Open(pathFromEnv)
		if innerErr != nil {
			return nil, innerErr
		}

		innerErr = v.ReadConfig(f)
		if innerErr != nil {
			return nil, innerErr
		}

		innerErr = f.Close()
		if innerErr != nil {
			return nil, innerErr
		}

		return newRuntimeFromViper(v, targetDir)

	}

	// look in the "standard" locations
	configDirs := configdir.New("kubernetes", "kep")
	systemConfigPath := configDirs.QueryFolders(configdir.System)[0].Path
	v.AddConfigPath(systemConfigPath)

	userConfigPath := configDirs.QueryFolders(configdir.Global)[0].Path
	v.AddConfigPath(userConfigPath)

	v.AddConfigPath(".") // working directory

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return newRuntimeFromViper(v, targetDir)
}

func Validate(runtime Runtime) error {
	var errs *multierror.Error

	if runtime.TargetDir() == "" {
		errs = multierror.Append(errs, errors.New("no target directory set"))
	}

	// TODO add existance test
	if runtime.CacheDir() == "" {
		errs = multierror.Append(errors.New("no cache directory set"))
	}

	// TODO add existance test
	if runtime.ContentRoot() == "" {
		errs = multierror.Append(errs, errors.New("no KEP content root set"))
	}

	contentRootDir := filepath.Base(runtime.ContentRoot())
	if contentRootDir != requiredContentRootName {
		errs = multierror.Append(errs, errors.New("KEP content root must be a directory named `content`"))
	}

	// TODO add existance test
	if runtime.TokenPath() == "" {
		errs = multierror.Append(errors.New("no GitHub token path set"))
	}

	if runtime.PrincipalEmail() == "" {
		errs = multierror.Append(errs, errors.New("no contact email set"))
	}

	if runtime.PrincipalGithubHandle() == "" {
		errs = multierror.Append(errs, errors.New("no GitHub handle set"))
	}

	if runtime.PrincipalDisplayName() == "" {
		errs = multierror.Append(errs, errors.New("no display name set"))
	}

	return errs.ErrorOrNil()
}

func newRuntimeFromViper(v *viper.Viper, targetDir string) (Runtime, error) {
	configDirs := configdir.New("kubernetes", "kep")
	defaultCache := configDirs.QueryCacheFolder()
	v.SetDefault(configCacheDir, defaultCache.Path)
	v.SetDefault(configRepoRootKey, filepath.Dir(v.GetString(configContentRootKey)))

	principalGithubHandle := v.GetString(configGithubHandleKey)
	principalEmail := v.GetString(configEmailKey)
	principalDisplayName := v.GetString(configDisplayNameKey)
	contentRoot := v.GetString(configContentRootKey)
	repoRoot := v.GetString(configRepoRootKey)
	tokenPath := v.GetString(configTokenPathKey)
	cacheDir := v.GetString(configCacheDir)

	r := &RawRuntime{
		principalEmail:        principalEmail,
		principalGithubHandle: principalGithubHandle,
		principalDisplayName:  principalDisplayName,
		contentRoot:           contentRoot,
		repoRoot:              repoRoot,
		targetDir:             targetDir,
		tokenPath:             tokenPath,
		cacheDir:              cacheDir,
	}

	if verr := Validate(r); verr != nil {
		return nil, verr
	}

	return r, nil

}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(ConfigName)
	v.SetConfigType("yaml")

	configDirs := configdir.New("kubernetes", "kep")
	defaultCache := configDirs.QueryCacheFolder()
	v.SetDefault(configCacheDir, defaultCache.Path)

	return v
}

// RawRuntime is the raw struct for a settings.Runtime. It should only be used
// directly when building settings interactively and is exported for that purpose
type RawRuntime struct {
	targetDir             string
	contentRoot           string
	repoRoot              string
	tokenPath             string
	cacheDir              string
	principalEmail        string
	principalGithubHandle string
	principalDisplayName  string
}

func (r *RawRuntime) TargetDir() string             { return r.targetDir }
func (r *RawRuntime) ContentRoot() string           { return r.contentRoot }
func (r *RawRuntime) RepoRoot() string              { return r.repoRoot }
func (r *RawRuntime) TokenPath() string             { return r.tokenPath }
func (r *RawRuntime) CacheDir() string              { return r.cacheDir }
func (r *RawRuntime) PrincipalEmail() string        { return r.principalEmail }
func (r *RawRuntime) PrincipalGithubHandle() string { return r.principalGithubHandle }
func (r *RawRuntime) PrincipalDisplayName() string  { return r.principalDisplayName }
