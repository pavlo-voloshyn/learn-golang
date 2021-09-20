package services

import "github.com/pavlo-voloshyn/init/entities"

type StudentService interface {
	GetById(id uint) entities.Student
	GetAll() []entities.Student
	Create(st entities.Student)
	Delete(id uint) error
}

type studentService struct {
	data []entities.Student
}

func NewStudentService() StudentService {
	return &studentService{}
}

func (s *studentService) GetById(id uint) entities.Student {
	for _, st := range s.data {
		if st.ID == id {
			return st
		}
	}
	panic("Not Found")
}
func (s *studentService) GetAll() []entities.Student {
	return s.data
}
func (s *studentService) Create(st entities.Student) {
	s.data = append(s.data, st)
}
func (s *studentService) Delete(id uint) error {
	for i, st := range s.data {
		if st.ID == id {
			s.data[i] = s.data[len(s.data)-1]
			s.data = s.data[:len(s.data)-1]
			break
		}
	}

	panic("Not Found")
}
