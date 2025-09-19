package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type DB map[string]string

func dataFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".folder_sticky_note.json"), nil
}

func NormalizePath(p string) (string, error) {
	if p == "" {
		return "", errors.New("empty path")
	}

	if strings.HasPrefix(p, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if p == "~" {
			p = home
		} else if strings.HasPrefix(p, "~/") {
			p = filepath.Join(home, p[2:])
		}
	}

	abs, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	abs = filepath.Clean(abs)
	if real, err := filepath.EvalSymlinks(abs); err == nil {
		abs = real
	}
	return abs, nil
}

func load() (DB, error) {
	path, err := dataFile()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return DB{}, nil
		}
		return nil, err
	}
	var db DB
	if len(b) == 0 {
		return DB{}, nil
	}

	if err := json.Unmarshal(b, &db); err != nil {
		return nil, err
	}
	return db, nil
}

func save(db DB) error {
	path, err := dataFile()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func Add(folder, memo string) (string, error) {
	abs, err := NormalizePath(folder)
	if err != nil {
		return "", err
	}
	db, err := load()
	if err != nil {
		return "", err
	}
	db[abs] = memo
	if err := save(db); err != nil {
		return "", err
	}
	return abs, nil
}
