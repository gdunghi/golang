package golang

type langStruct struct {
	ID   int    `xml:"id" json:"id"`
	Lang string `xml:"lang" json:"lang"`
}

//Book struct
type Book struct {
	Auther string `xml:"auther"`
	Name   string `xml:"name"`
}
