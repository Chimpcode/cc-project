package main

type DataType struct {
	Name string `json:"name"`
	Persistence string `json:"persistence"`
	Descriptor string `json:"descriptor"`
	ChimpcodeHelper bool `json:"chimpcode_helper"`
}

type ChimpCodeSeed struct {
	ProjectName string `json:"project_name"`
	Author string `json:"author"`
	Description string `json:"description"`

	BackendFramework string `json:"backend_framework"`
	UseGraphQL bool `json:"use_graphql"`

	DatabaseType string `json:"database_type"`

	StorageFramework string `json:"storage_framework"`

	DataTypes []DataType `json:"data_types"`
}
