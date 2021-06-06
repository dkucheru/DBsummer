-- AUTOGENERATED BY storj.io/dbx
-- DO NOT EDIT
CREATE TABLE groups_s (
	cipher text NOT NULL,
	groupname text NOT NULL,
	educationalyear text NOT NULL,
	semester text NOT NULL,
	course text NOT NULL,
	subject integer NOT NULL,
	PRIMARY KEY ( cipher )
);
CREATE TABLE runners (
	runner_number integer NOT NULL,
	date_of_compilation timestamp with time zone NOT NULL,
	date_of_expiration timestamp with time zone NOT NULL,
	postponing_reason text NOT NULL,
	type_of_control text NOT NULL,
	teacher text NOT NULL,
	PRIMARY KEY ( runner_number )
);
CREATE TABLE runner_marks (
	check_mark integer NOT NULL,
	runner_mark_number integer NOT NULL,
	national_mark text NOT NULL,
	semester_mark integer NOT NULL,
	together_mark integer NOT NULL,
	ects_mark text NOT NULL,
	sheet_mark integer NOT NULL,
	runner integer NOT NULL,
	PRIMARY KEY ( runner_mark_number ),
	UNIQUE ( sheet_mark )
);
CREATE TABLE sheets (
	sheetid integer NOT NULL,
	number_of_attendees integer NOT NULL,
	number_of_absent integer NOT NULL,
	number_of_ineligible integer NOT NULL,
	type_of_control text NOT NULL,
	date_of_compilation timestamp with time zone NOT NULL,
	teacher text NOT NULL,
	group_cipher text NOT NULL,
	PRIMARY KEY ( sheetid ),
	UNIQUE ( group_cipher )
);
CREATE TABLE sheet_marks (
	check_mark integer NOT NULL,
	mark_number integer NOT NULL,
	national_mark text NOT NULL,
	semester_mark integer NOT NULL,
	together_mark integer NOT NULL,
	ects_mark text NOT NULL,
	sheet integer NOT NULL,
	student text NOT NULL,
	PRIMARY KEY ( mark_number )
);
CREATE TABLE students (
	student_cipher text NOT NULL,
	firstname text NOT NULL,
	last_name text NOT NULL,
	middle_name text NOT NULL,
	record_book_number text NOT NULL,
	PRIMARY KEY ( student_cipher )
);
CREATE TABLE subjects (
	subjectid integer NOT NULL,
	subjectname text NOT NULL,
	educationallevel text NOT NULL,
	faculty text NOT NULL,
	PRIMARY KEY ( subjectid )
);
CREATE TABLE tableNews (
	id_t text NOT NULL,
	PRIMARY KEY ( id_t )
);
CREATE TABLE teachers (
	teacher_cipher text NOT NULL,
	firstname text NOT NULL,
	lastname text NOT NULL,
	middlename text NOT NULL,
	scientificdegree text NOT NULL,
	academictitles text NOT NULL,
	post text NOT NULL,
	PRIMARY KEY ( teacher_cipher )
);
