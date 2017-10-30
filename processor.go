package main

import (
	"errors"
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	CreateDir = "create_dir"
	CreateFile = "create_file"
)

func makeProcess(task ChimpcodeTask) error {
	switch task.GetTypeOfProcess() {
	case CreateDir:
		if err := MakeDir(task.GetName()); err != nil {
			return err
		}
	case CreateFile:
		if strings.Contains(task.GetContent(), "file:") {
			log.Println("Into file content", task.GetContent())
			// The content es in a file and We needs read this file
			pathToFile := strings.Replace(task.GetContent(), "file:", "", -1)

			dir, err := GetCurrentDir()
			if err != nil {
				return err
			}

			absPath := filepath.Join(*dir, pathToFile)

			data, err := ioutil.ReadFile(absPath)

			if err != nil {
				return err
			}

			err = MakeFile(task.GetName(), string(data))
			if err != nil {
				return err
			}


		} else {
			// The content is explicit into content prop
			err := MakeFile(task.GetName(), task.GetContent())
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func ProcessTask(task interface{}) error {
	switch v := task.(type) {
	default:
		return errors.New(fmt.Sprintf("Unexpected type %T", v))
	case ChimpcodeTaskNode:
		err := makeProcess(ChimpcodeTask(v))
		if err != nil {
			return err
		}
		err = ProcessTask(v.ChildTask)
		if err != nil {
			return err
		}
	case ChimpcodeTaskLeaf:
		err := makeProcess(ChimpcodeTask(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessChimpcodeSeed(config *ChimpCodeSeed) error {
	content := "# " + config.ProjectName + "\n"
	content += "## Author: " + config.Author + "\n-----\n"
	content += config.Description

	finalTasks := append(PrincipalTasks, ChimpcodeTaskLeaf{
		Name: "README.md",
		TypeOfProcess: "create_file",
		Content: content,
	})

	for _, t := range finalTasks {
		err := ProcessTask(t)
		if err != nil {
			return err
		}
	}
	return nil
}
