package run

type pr struct {
	Link   string
	Title  string
	ID     string
	Points int32
}

type inpr struct {
	Name  string `json:"name" binding:"required"`
	Link  string `json:"link" binding:"required"`
	Title string `json:"title" binding:"required"`
	ID    string `json:"id" binding:"required"`
}

type participant struct {
	Name   string
	Prs    []pr
	Points int32
}
