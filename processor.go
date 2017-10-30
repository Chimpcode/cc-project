package main

import (
	"errors"
	"fmt"
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
		err := MakeFile(task.GetName(), task.GetContent())
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessTask(task interface{}) error {
	switch v := task.(type) {
	default:
		return errors.New(fmt.Sprintf("Unexpected type %T", v))
	case ChimpcodeTaskNode:
		makeProcess(ChimpcodeTask(v))
		err := ProcessTask(v.ChildTask)
		if err != nil {
			return err
		}
	case ChimpcodeTaskLeaf:
		makeProcess(ChimpcodeTask(v))

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
