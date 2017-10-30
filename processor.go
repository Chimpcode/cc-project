package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Create = "create"
)

func ProcessTask(task interface{}) error {
	switch v := task.(type) {
	default:
		return errors.New(fmt.Sprintf("Unexpected type %T", v))
	case ChimpcodeTaskNode:
		if strings.EqualFold(v.TypeOfProcess, Create) {
			if err := MakeDir(v.Name); err != nil {
				return err
			}
		}
		err := ProcessTask(v.ChildTask)
		if err != nil {
			return err
		}
	case ChimpcodeTaskLeaf:
		if strings.EqualFold(v.TypeOfProcess, Create) {
			if err := MakeDir(v.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

func ProcessChimpcodeSeed(config *ChimpCodeSeed) error {
	for _, t := range MainTasks {
		err := ProcessTask(t)
		if err != nil {
			return err
		}
	}
	return nil
}
