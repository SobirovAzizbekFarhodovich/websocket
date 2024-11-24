package domain

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FileData string `json:"file_data,omitempty"`
	IsJoin   bool   `json:"is_join,omitempty"`
}
