package git

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

// Remote struct is simply a collection of remote branches and wraps it with the
// name of the remote and fetch/push urls. It also holds the *selected* remote
// branch
type Remote struct {
	Name     string
	URL      []string
	RefSpecs []string
	Branches []*RemoteBranch
}

// search for remotes in go-git way. It is the short way to get remotes but it
// does not give any insght about remote branches
func (r *Repository) initRemotes() error {
	rp := r.Repo
	r.Remotes = make([]*Remote, 0)

	rms, err := rp.Remotes()
	for _, rm := range rms {
		rfs := make([]string, 0)
		for _, rf := range rm.Config().Fetch {
			rfs = append(rfs, string(rf))
		}
		remote := &Remote{
			Name:     rm.Config().Name,
			URL:      rm.Config().URLs,
			RefSpecs: rfs,
		}
		remote.loadRemoteBranches(r)
		r.Remotes = append(r.Remotes, remote)

	}
	if err != nil {
		log.Warn("Cannot find remotes " + err.Error())
		return err
	}
	if len(r.Remotes) <= 0 {
		return errors.New("no remote for repository: " + r.Name)
	}
	r.State.Remote = r.Remotes[0]
	return err
}
