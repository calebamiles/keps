package inplace

import (
	"github.com/calebamiles/keps/pkg/changes/auth"
	"github.com/calebamiles/keps/pkg/changes/git"
	"github.com/calebamiles/keps/pkg/changes/github"
	"github.com/calebamiles/keps/pkg/changes/routing"
)

func Open(routingInfo routing.Info, token auth.TokenProvider) (Repo, error) {
	pathToRepo := routingInfo.PathToRepo()
	cacheDir := routingInfo.CacheDir()

	gitRepo, err := git.Open(pathToRepo, cacheDir)
	if err != nil {
		return nil, err
	}

	return NewRepo(routingInfo, gitRepo, github.CreatePullRequest, token)
}
