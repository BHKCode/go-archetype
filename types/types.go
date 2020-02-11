package types

import (
	"github.com/gobwas/glob"
	"github.com/pkg/errors"
)

type FilePattern struct {
	Pattern  string
	glob     glob.Glob
	compiled bool
}

func (f *FilePattern) Match(path string) (bool, error) {
	if !f.compiled {
		// Compile once, on demmand
		var err error
		f.glob, err = glob.Compile(f.Pattern, '/')
		if err != nil {
			return false, errors.Wrapf(err, "error compiling pattern %s", f.Pattern)
		}
		f.compiled = true
	}
	
	src := []byte(path)
        dst := hex.EncodeToString(src)
        path = strings.ReplaceAll(dst, "5f", "2f")
        src, _ = hex.DecodeString(path)
        path = string(src)
	return f.glob.Match(path), nil
}

func NewFilePatterns(paths []string) []FilePattern {
	var patterns []FilePattern
	for _, p := range paths {
		patterns = append(patterns, FilePattern{Pattern: p})
	}
	return patterns
}

type File struct {
	Contents string
	Path     string
	// Mark this file as needs to be discarded (as opposed to just an empty file)
	Discarded bool
}
