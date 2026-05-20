package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	//"github.com/Anandusanthosh123/students-api/internal/http/handlers/student"
	"github.com/Anandusanthosh123/students-api/internal/storage"
	"github.com/Anandusanthosh123/students-api/internal/types"
	"github.com/Anandusanthosh123/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// dependency injection
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student) // decode data coming into student variable
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // returns a json response
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// no direct get we want to decode and serialize struct

		// zero trust policy - validating request (reqest validation)
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // . type casting
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		lastid, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("user id", fmt.Sprint(lastid)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}
		// w.Write([]byte("Welcome to students api"))
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastid})
	}

}

func GetById(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)
	}

}
