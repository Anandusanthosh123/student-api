package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Anandusanthosh123/students-api/internal/config"
	"github.com/Anandusanthosh123/students-api/internal/types"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {

		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil

} // function receiver
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stat, err := s.Db.Prepare("INSERT INTO students(name,email,age)VALUES (?,?,?)") // giving placeholders
	if err != nil {
		return 0, err
	}
	defer stat.Close()
	result, err := stat.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastid, nil

}
func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {

	stat, err := s.Db.Prepare("SELECT id,name,email,age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stat.Close()
	var student types.Student
	err = stat.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error:%w", err)
	}
	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []types.Student
	// cursor to select next row
	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
func (s *Sqlite) UpdateStudent(id int64, name string, email string, age int) error {
	stmt, err := s.Db.Prepare("UPDATE students SET name=?, email=?, age=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no student found with id %d", id)
	}

	return nil
}

func (s *Sqlite) DeleteStudent(id int64) error {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no student found with id %d", id)
	}

	return nil
}
