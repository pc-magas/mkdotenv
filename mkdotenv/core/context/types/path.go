package path

import("filepath")

struct ContextPath {
	path string, // User provided path
	workingDirectory string // Directory where path related upon
}

func NewContextPath(path string, workingDirectory String) ContextPath {
	return ContextPath{path: path,workingDirectory:workingDirectory}
}

func (p ContextPath) Value() (string, error) {
	if p.path == "" {
		return "", fmt.Errorf("empty path")
	}

	if filepath.IsAbs(p.path) {
		return p.path, nil
	}

	return filepath.Join(ctx.workingDirectory, p.path), nil
}
