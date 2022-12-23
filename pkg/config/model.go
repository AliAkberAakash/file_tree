package config

type FileStruct struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Children []FileStruct `json:"children,omitempty"`
}
