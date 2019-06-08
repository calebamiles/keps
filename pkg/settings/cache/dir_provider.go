package cache

type Dir string

// DirProvider is satisfied by settings.Runtime
type DirProvider interface {
	CacheDir() string
}
