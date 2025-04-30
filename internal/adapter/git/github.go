package git

import ()

type GitHub struct{}

func (gh *GitHub) checkRepoExist() (bool, error) {
	return true, nil
}

func (gh *GitHub) checkConnection() (bool, error) {
	return true, nil
}
