package constants

import (
	"path"
)

// mod related constants
const (
	WorkspaceModDir         = "mods"
	WorkspaceDataDir        = ".steampipe"
	WorkspaceConfigFileName = "workspace.spc"
	WorkspaceIgnoreFile     = ".steampipeignore"
	WorkspaceDefaultModName = "local"
	WorkspaceModFileName    = "mod.sp"
	DefaultVarsFileName     = "steampipe.spvars"
)

func WorkspaceModPath(workspacePath string) string {
	return path.Join(workspacePath, WorkspaceDataDir, WorkspaceModDir)
}
func DefaultVarsFilePath(workspacePath string) string {
	return path.Join(workspacePath, DefaultVarsFileName)
}
