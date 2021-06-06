package dataBaseDir

import (
	"DBsummer/structs"
	"context"
)

type Repository interface {
	Close() error
	BeginTx(ctx context.Context) (Transaction, error)
	TableNew() TableNewInterface
	Subjects() SubjectsInterface
	Groups() GroupsInterface
	Students() StudentsInterface
	Teachers() TeachersInterface
	Sheets() SheetsInterface
	RunnerMarks() RunnerMarksInterface
	Runners() RunnerInterface
	SheetMarks() SheetMarksInterface
}

type Transaction interface {
	Repository

	Commit() error
	Rollback() error
}

type TableNewInterface interface { //functions to be used for the table 'tablen'
	Get(ctx context.Context, id int) error
	Create(ctx context.Context) (id int, err error)
}

type SubjectsInterface interface { //functions to be used for the table 'subjects'
	Get(ctx context.Context, id int) (*structs.Subject, error)
	GetAll(ctx context.Context) ([]*structs.Subject, error)
	Create(ctx context.Context) (id int, err error)
}

type GroupsInterface interface {
}

type StudentsInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context, id int) (*structs.Student, error)
	GetAllStudInfo(ctx context.Context) ([]*structs.AllStudInfo, error)
	GetAllBorjniki(ctx context.Context) ([]*structs.AllStudInfo, error)
	GetStudentByPIB(ctx context.Context, fn string, ln string, mn string) ([]*structs.AllStudInfo, error)
}

type TeachersInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type SheetsInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type RunnerMarksInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type RunnerInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}

type SheetMarksInterface interface {
	Create(ctx context.Context) (id int, err error)
	Get(ctx context.Context) error
}
