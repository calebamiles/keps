package enhancements

import (
	"github.com/calebamiles/keps/pkg/changes"
	"github.com/calebamiles/keps/pkg/changes/auth"
	"github.com/calebamiles/keps/pkg/changes/inplace"
	"github.com/calebamiles/keps/pkg/keps"
	"github.com/calebamiles/keps/pkg/orgs"
	"github.com/calebamiles/keps/pkg/settings"
)

// TODO add docstring
func OpenRepo(runtime settings.Runtime, org orgs.Instance, kep keps.Instance) (changes.Submitter, error) {
	routingInfo, err := GatherProposalInfoFrom(runtime, org, kep)
	if err != nil {
		return nil, err
	}

	// build auth.tokenProvider
	token, err := auth.NewProvideTokenFromPath(runtime.TokenPath())
	if err != nil {
		return nil, err
	}

	// open repo
	repo, err := inplace.Open(routingInfo, token)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
