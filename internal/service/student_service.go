package service

type Student struct {
}

func NewStudentHanle() *Student {
	return &Student{}
}
func (student *Student) GetStudentName(name string) (r string, err error) {
	return "你好", nil
}
