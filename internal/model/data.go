package model

import "strings"

type Data struct {
	FirstName string
	LastName  string
	Number    string
}

// New parse a line to create an instance
// of data.
func New(line string) Data {
	line = strings.TrimSuffix(line, "\n")
	data := strings.Split(line, ",")

	model := Data{
		FirstName: data[0],
		LastName:  data[1],
		Number:    data[2],
	}

	return model
}
