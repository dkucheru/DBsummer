package serviceDir

import (
	"DBsummer/dataBaseDir"
	"DBsummer/pdfReading"
	"DBsummer/structs"
	"context"
)

type studentsService struct {
	service    *Service
	repository dataBaseDir.Repository
}

func newStudentsService(service *Service, repository dataBaseDir.Repository) *studentsService {
	subServ := studentsService{
		service:    service,
		repository: repository,
	}
	return &subServ
}

func (s *studentsService) Create(ctx context.Context, subject structs.Student) (structs.Student, error) {
	panic("implement me")
}

func (s *studentsService) Get(ctx context.Context, idVid int) (*structs.Student, error) {
	panic("implement me")
}

func (s *studentsService) GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error) {
	subj, err := s.repository.Students().GetAllStudInfo(ctx)
	if err != nil {
		return nil, err
	}

	return subj, nil
}

func (s *studentsService) GetAllBorjniki(ctx context.Context) ([]*structs.AllStudInfo, error) {
	subj, err := s.repository.Students().GetAllBorjniki(ctx)
	if err != nil {
		return nil, err
	}

	return subj, nil
}

func (s *studentsService) GetStudentByPIB(ctx context.Context, fn string, ln string, mn string, year string) ([]*structs.StudentMarks, error) {
	subj, err := s.repository.Students().GetStudentByPIB(ctx, fn, ln, mn, year)
	if err != nil {
		return nil, err
	}

	return subj, nil
}

func (s *studentsService) AddStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error) {
	id, err := s.repository.Students().AddStudent(ctx, sheetMarks)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *studentsService) GetPIBAllStudents(ctx context.Context) ([]*structs.StudentPIB, error) {
	PIBs, err := s.repository.Students().GetPIBAllStudents(ctx)
	if err != nil {
		return nil, err
	}

	return PIBs, nil
}

func (s *studentsService) GetAllStudentMarksByID(ctx context.Context, id int) ([]*structs.StudentAllMarks, error) {
	marks, err := s.repository.Students().GetAllStudentMarksByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return marks, nil
}

func (s *studentsService) FindStudent(ctx context.Context, sheetMarks *pdfReading.StudInfoFromPDF) (*int, error) {
	id, err := s.repository.Students().FindStudent(ctx, sheetMarks)
	if err != nil {
		return nil, err
	}

	return id, nil
}
