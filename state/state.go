// Package state :
//
// The state package provides unified state management utilities for the installation and for
// packages to consume, without creating import cycles.
//
// This package should not import any other steampipe packages.
package state

// SteampipeState :: the current steampipe state
type SteampipeState struct {
	InstallationID string
	LastTaskRun    string
	Passwords      DBPasswords
	Service        RunningDBInstanceInfo
}

// RunningDBInstanceInfo :: contains data about the running process
// and it's credentials
type RunningDBInstanceInfo struct {
	Pid        int
	Port       int
	Listen     []string
	ListenType string
	Invoker    string
	Password   string
	User       string
	Database   string
}

// DBPasswords :: contains the passwords that were set in the DB during installation
type DBPasswords struct {
	RootPassword      string
	SteampipePassword string
}

var currentState SteampipeState

func init() {
	if load() != nil {
		panic("Error loading state information")
	}
}

func load() error {
	// stateFile := filepath.Join(constants.InternalDir(), "steampipe.json")
	return nil
}
