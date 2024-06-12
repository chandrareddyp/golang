package main

import "fmt"


func main() {
	var s Student
	s = NewStudent("john", -1, "id is negative")
	err := s.ProcessStudent()
	if err != nil {	
		fmt.Println(err)
	}
}

type Student struct {
	Name string
	id int
	errorMsg string
}

func (s Student) Error() string {
	return fmt.Sprintf("there is an error with student %s, id: %d, error msg: %s", s.Name, s.id, s.errorMsg)
}
func NewStudent(name string, id int, errorMsg string) Student {
   return Student{
	Name: name,
	id: id,
	errorMsg: errorMsg,
	}
}

func (s Student) ProcessStudent() error {
	if s.id < 0 {
		return NewStudent(s.Name, s.id, "id is negative")
	}
	return nil
}