package entities

type Task struct {
	Id        int64
	InputTask string `validate:"required" label:"Input Task"`
	Name      string `validate:"required" label:"Nama"`
	Deadline  string `validate:"required" label:"Deadline"`
}
