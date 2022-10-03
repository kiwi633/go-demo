package rander

type Data struct {
	ContentType string
	Data        []byte
	Name        string
	DateTime    string
	CreateTime  string
}

type User struct {
	Name   string
	IdType string
	IdNo   string
}
