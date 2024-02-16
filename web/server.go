package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/db"
)

var dbConnection *sql.DB
var err error

// func main() {
// 	// fmt.Println("This is not a web server")
// 	// http.HandleFunc("/sum", sumHandler)
// 	// http.ListenAndServe(":8080", nil)

// 	http.ListenAndServe(":8080", handler())
// }

func Handler(dbConn *sql.DB) http.Handler {
	// Connect and get a database handle.
	dbConnection = dbConn

	mux := http.NewServeMux()
	mux.HandleFunc("/department", departmentHandler)
	mux.HandleFunc("/employee", employeeHandler)
	return mux
}

func departmentHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		departmentName := r.FormValue("departmentName")
		fmt.Printf("Post Li %s", departmentName)

		var deptIdDb int64
		if departmentName == "" {
			http.Error(w, "missing department name", http.StatusBadRequest)
			return
		}
		if deptIdDb, err := db.CreateDepartment(dbConnection, departmentName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintln(os.Stdout, []any{"deptIdDb %v", deptIdDb}...)
			return
		}
		w.WriteHeader(http.StatusCreated)
		department, err := db.ViewDepartmentById(dbConnection, int64(deptIdDb))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(department)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	case http.MethodGet:
		departmentId := r.FormValue("departmentId")
		fmt.Printf("MethodGetLIa %s", departmentId)

		if departmentId == "" {
			departments, err := db.ViewDepartment(dbConnection)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b, err := json.Marshal(departments)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)

			fmt.Printf("MethodGetLIa %s", b)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		vDeptId, _ := strconv.Atoi(departmentId)
		department, err := db.ViewDepartmentById(dbConnection, int64(vDeptId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(department)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	//todo
}

// func sumHandler(w http.ResponseWriter, r *http.Request) {
// 	a := r.FormValue("a")
// 	b := r.FormValue("b")
// 	va, _ := strconv.Atoi(a)
// 	vb, _ := strconv.Atoi(b)
// 	s := sum.Ints(va, vb)

// 	w.Write([]byte(fmt.Sprintf("%d", s)))
// }
