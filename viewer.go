package main

import (
	"path/filepath"
	"os"
)

func GetCurrentDir() (*string, error){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return &dir, err
}

