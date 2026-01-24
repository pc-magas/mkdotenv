package types

import (
	"path" 
)

type ContextPath struct {
	providedPath string // User provided path
	workingDirectory string // Directory where path related upon
}

func NewContextPath(providedPath string, workingDirectory string) ContextPath {
	return ContextPath{providedPath: providedPath,workingDirectory:workingDirectory}
}

func (p ContextPath) Value() string {
	if p.providedPath == "" {
		return ""
	}

	if path.IsAbs(p.providedPath) {
		return p.providedPath
	}

	return path.Join(p.workingDirectory, p.providedPath)
}
