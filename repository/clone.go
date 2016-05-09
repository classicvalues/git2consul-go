package repository

import (
	"os"

	"gopkg.in/libgit2/git2go.v23"
)

// Clone the repository
func (r *Repository) Clone() error {
	r.Lock()
	defer r.Unlock()

	err := os.Mkdir(r.store, 0755)
	if err != nil {
		return err
	}

	raw_repo, err := git.Clone(r.repoConfig.Url, r.store, &git.CloneOptions{
		CheckoutOpts: &git.CheckoutOpts{
			Strategy: git.CheckoutNone,
		},
	})
	if err != nil {
		return err
	}

	r.Repository = raw_repo

	err = r.checkoutConfigBranches()
	if err != nil {
		return err
	}

	r.cloneCh <- struct{}{}

	return nil
}
