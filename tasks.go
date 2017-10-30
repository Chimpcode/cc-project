package main

type ChimpcodeTaskNode struct {
	Name string `json:"name"`
	TypeOfProcess string `json:"type_of_process"`
	Content string `json:"content"`
	ChildTask interface{} `json:"child_task"`
}

type ChimpcodeTaskLeaf struct {
	Name string `json:"name"`
	TypeOfProcess string `json:"type_of_process"`
	Content string `json:"content"`
}

type ChimpcodeTasks []interface{}

type ChimpcodeTask interface {
	GetName() string
	GetTypeOfProcess() string
	GetContent() string
	//SetContent(content string)
}

func (t ChimpcodeTaskNode) GetName() string {
	return t.Name
}
func (t ChimpcodeTaskNode) GetTypeOfProcess() string {
	return t.TypeOfProcess
}
func (t ChimpcodeTaskNode) GetContent() string {
	return t.Content
}

//func (t ChimpcodeTaskNode) SetContent(content string) {
//	t.Content = content
//}


func (t ChimpcodeTaskLeaf) GetName() string {
	return t.Name
}
func (t ChimpcodeTaskLeaf) GetTypeOfProcess() string {
	return t.TypeOfProcess
}
func (t ChimpcodeTaskLeaf) GetContent() string {
	return t.Content
}
//func (t ChimpcodeTaskLeaf) SetContent(content string) {
//	t.Content = content
//}
