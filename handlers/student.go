package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"fealtyx/models"
	"fealtyx/utils"
)

var (
	students = make(map[int]models.Student)
	mutex    = &sync.Mutex{}
)

// POST /students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil || s.Name == "" || s.Email == "" || s.Age <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	if _, exists := students[s.ID]; exists {
		http.Error(w, "Student already exists", http.StatusBadRequest)
		return
	}
	students[s.ID] = s
	json.NewEncoder(w).Encode(s)
}

// GET /students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var all []models.Student
	for _, s := range students {
		all = append(all, s)
	}
	json.NewEncoder(w).Encode(all)
}

// GET /students/{id}
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	studentID, _ := strconv.Atoi(id)

	mutex.Lock()
	defer mutex.Unlock()
	s, ok := students[studentID]
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// PUT /students/{id}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	studentID, _ := strconv.Atoi(id)

	var update models.Student
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := students[studentID]; !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	update.ID = studentID
	students[studentID] = update
	json.NewEncoder(w).Encode(update)
}

// DELETE /students/{id}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	studentID, _ := strconv.Atoi(id)

	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := students[studentID]; !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	delete(students, studentID)
	w.WriteHeader(http.StatusNoContent)
}

// GET /students/{id}/summary
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	studentID, _ := strconv.Atoi(id)

	mutex.Lock()
	student, ok := students[studentID]
	mutex.Unlock()

	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	summary, err := utils.GenerateSummary(student)
	if err != nil {
		http.Error(w, "Error generating summary", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}
