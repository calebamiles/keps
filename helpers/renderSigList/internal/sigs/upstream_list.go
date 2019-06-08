package sigs

type upstreamSIGList struct {
	SIGs []upstreamSIGEntry `json:"sigs"`
}

type upstreamSIGEntry struct {
	Name        string                    `json:"name"` // we actually want to look at what the SIG is called on disk
	Subprojects []upstreamSubprojectEntry `json:"subprojects"`
}

type upstreamSubprojectEntry struct {
	Name string `json:"name"`
}

const UpstreamSIGListURL = "https://raw.githubusercontent.com/kubernetes/community/master/sigs.yaml"
