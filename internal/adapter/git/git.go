package git

import ()

type GitStore struct {}

type GitAction interface {
  CheckRepoExist() (bool, error)
  CheckConnection() (bool, error)
}
