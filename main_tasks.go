package main

type ChimpcodeTaskNode struct {
	Name string `json:"name"`
	TypeOfProcess string `json:"type_of_process"`
	ChildTask interface{} `json:"child_task"`
}

type ChimpcodeTaskLeaf struct {
	Name string `json:"name"`
	TypeOfProcess string `json:"type_of_process"`

}

type ChimpcodeTasks []interface{}

var MainTasks ChimpcodeTasks = ChimpcodeTasks{
	ChimpcodeTaskLeaf{"api", "create"},
	ChimpcodeTaskLeaf{"database", "create"},
	ChimpcodeTaskLeaf{"db", "create"},
	ChimpcodeTaskLeaf{"storage", "create"},
	ChimpcodeTaskLeaf{"types", "create"},
	ChimpcodeTaskLeaf{"utils", "create"},
	ChimpcodeTaskLeaf{"globals", "create"},
}

