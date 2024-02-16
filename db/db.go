package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Department struct {
	DepartmentId   int64  `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
}

type Employee struct {
	EmployeeId   int64  `json:"employeeId"`
	EmployeeName string `json:"employeeName"`
	DepartmentId int64  `json:"departmentId"`
}

func CreateConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error

	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "level3",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return db, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return db, pingErr
	}
	fmt.Println("Connected!")

	return db, nil
}

// createDepartment adds the specified department to the database,
// returning the department ID of the new entry
func CreateDepartment(db *sql.DB, departmentName string) (int64, error) {

	result, err := db.Exec("INSERT INTO department (department_name) VALUES (?)", departmentName)
	if err != nil {
		return 0, fmt.Errorf("createDepartment: %v", err)
	}
	lastDepartmentId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createDepartment: %v", err)
	}
	return lastDepartmentId, nil
}

// updateDepartment update the specified department to the database,
// returning the amount of department of the updated entry
func UpdateDepartment(db *sql.DB, departmentName string, departmentId int64) (int64, error) {
	result, err := db.Exec("UPDATE department SET department_name = (?) WHERE department_id = (?)", departmentName, departmentId)
	if err != nil {
		return 0, fmt.Errorf("updateDepartment: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("createDepartment: %v", err)
	}
	return updatedRow, nil
}

// deleteDepartment delete the specified department from the database,
// returning the amount of department of the deleted entry
func DeleteDepartmentById(db *sql.DB, departmentId int64) (int64, error) {
	result, err := db.Exec("DELETE FROM department WHERE department_id = (?)", departmentId)
	if err != nil {
		return 0, fmt.Errorf("deleteDepartment: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteDepartment: %v", err)
	}
	return updatedRow, nil
}

// deleteAllDataDepartment delete all data department from the database,
// returning the amount of department of the deleted entry
func DeleteAllDataDepartment(db *sql.DB) (int64, error) {
	result, err := db.Exec("TRUNCATE TABLE department")
	if err != nil {
		return 0, fmt.Errorf("deleteAllDataDepartment: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAllDataDepartment: %v", err)
	}
	return updatedRow, nil
}

// deleteAllDataEmployee delete all data employee from the database,
// returning the amount of employee of the deleted entry
func DeleteAllDataEmployee(db *sql.DB) (int64, error) {
	result, err := db.Exec("TRUNCATE TABLE employee")
	if err != nil {
		return 0, fmt.Errorf("deleteAllDataEmployee: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAllDataEmployee: %v", err)
	}
	return updatedRow, nil
}

// viewDepartment view queries for all departments.
func ViewDepartment(db *sql.DB) ([]Department, error) {
	// An departments slice to hold data from returned rows.
	var departments []Department

	rows, err := db.Query("SELECT * FROM department")
	if err != nil {
		return nil, fmt.Errorf("viewDepartment %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dept Department
		if err := rows.Scan(&dept.DepartmentId, &dept.DepartmentName); err != nil {
			return nil, fmt.Errorf("employeesByDepartment %v", err)
		}
		departments = append(departments, dept)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeesByDepartment %v", err)
	}
	return departments, nil
}

// viewDepartment view queries for all departments.
func ViewDepartmentById(db *sql.DB, departmentId int64) ([]Department, error) {
	// An departments slice to hold data from returned rows.
	var departments []Department

	rows, err := db.Query("SELECT * FROM department WHERE department_id = (?)", departmentId)
	if err != nil {
		return nil, fmt.Errorf("viewDepartment %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var dept Department
		if err := rows.Scan(&dept.DepartmentId, &dept.DepartmentName); err != nil {
			return nil, fmt.Errorf("ViewDepartmentById %v", err)
		}
		departments = append(departments, dept)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ViewDepartmentById %v", err)
	}
	return departments, nil
}

// createEmployee adds the specified employee to the database,
// returning the employee ID of the new entry
func CreateEmployee(db *sql.DB, employeeName string, departmentId int64) (int64, error) {
	result, err := db.Exec("INSERT INTO employee (employee_name, department_id) VALUES (?, ?)", employeeName, departmentId)
	if err != nil {
		return 0, fmt.Errorf("createEmployee: %v", err)
	}
	lastEmployeeId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("createEmployee: %v", err)
	}
	return lastEmployeeId, nil
}

// updateEmployee update the specified employee to the database,
// returning the amount of employee of the updated entry
func UpdateEmployee(db *sql.DB, employeeName string, departmentName string, employeeId int64) (int64, error) {
	departmentDb, err := IdByDepartmentName(db, departmentName)
	if err != nil {
		return 0, fmt.Errorf("IdByDepartmentName: %v", err)
	}
	result, err := db.Exec("UPDATE employee SET employee_name = (?), department_id = (?) WHERE employee_id = (?)",
		employeeName, departmentDb.DepartmentId, employeeId)
	if err != nil {
		return 0, fmt.Errorf("updateEmployee: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateEmployee: %v", err)
	}
	return updatedRow, nil
}

// deleteEmployee delete the specified department from the database,
// returning the amount of employee of the deleted entry
func DeleteEmployeeById(db *sql.DB, employeeId int64) (int64, error) {
	result, err := db.Exec("DELETE FROM employee WHERE employee_id = (?)", employeeId)
	if err != nil {
		return 0, fmt.Errorf("deleteEmployeeById: %v", err)
	}
	updatedRow, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteEmployeeById: %v", err)
	}
	return updatedRow, nil
}

// viewEmployee view queries for all employees.
func ViewEmployee(db *sql.DB) ([]Employee, error) {
	// An employees slice to hold data from returned rows.
	var employees []Employee

	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		return nil, fmt.Errorf("viewEmployee %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var empl Employee
		if err := rows.Scan(&empl.EmployeeId, &empl.EmployeeName, &empl.DepartmentId); err != nil {
			return nil, fmt.Errorf("employees %v", err)
		}
		employees = append(employees, empl)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("viewEmployee %v", err)
	}
	return employees, nil
}

// idByDepartmentName queries for the department_id with the specified department_name.
func IdByDepartmentName(db *sql.DB, departmentName string) (Department, error) {
	// An album to hold data from the returned row.
	var dept Department

	row := db.QueryRow("SELECT * FROM department WHERE department_name = ?", departmentName)
	if err := row.Scan(&dept.DepartmentId, &dept.DepartmentName); err != nil {
		if err == sql.ErrNoRows {
			return dept, fmt.Errorf("idByDepartmentName %q: no such department", departmentName)
		}
		return dept, fmt.Errorf("IdByDepartmentName %q: %v", departmentName, err)
	}
	return dept, nil
}

// DepartmentNameById queries for the department_name with the specified department_id.
func DepartmentNameById(db *sql.DB, departmentId int64) (Department, error) {
	// An album to hold data from the returned row.
	var dept Department

	row := db.QueryRow("SELECT * FROM department WHERE department_id = ?", departmentId)
	if err := row.Scan(&dept.DepartmentId, &dept.DepartmentName); err != nil {
		if err == sql.ErrNoRows {
			return dept, fmt.Errorf("DepartmentNameById %q: no such department with ID ", departmentId)
		}
		return dept, fmt.Errorf("DepartmentNameById %q: %v", departmentId, err)
	}
	return dept, nil
}

// EmployeeNameById queries for the employee_name with the specified employee_id.
func EmployeeNameById(db *sql.DB, employeeId int64) (Employee, error) {
	// An album to hold data from the returned row.
	var empl Employee

	row := db.QueryRow("SELECT * FROM employee WHERE employee_id = ?", employeeId)
	if err := row.Scan(&empl.EmployeeId, &empl.EmployeeName, &empl.DepartmentId); err != nil {
		if err == sql.ErrNoRows {
			return empl, fmt.Errorf("EmployeeNameById %q: no such employee with ID ", employeeId)
		}
		return empl, fmt.Errorf("EmployeeNameById %q: %v", employeeId, err)
	}
	return empl, nil
}

// employeesByDepartment queries for employees that have the specified department.
func ViewEmployeesByDepartment(db *sql.DB, departmentName string) ([]Employee, error) {
	// An albums slice to hold data from returned rows.
	var employees []Employee

	// Get departmentId from departmentName.
	department, err := IdByDepartmentName(db, departmentName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Department ID found: %v\n", department.DepartmentId)

	rows, err := db.Query("SELECT * FROM employee WHERE department_id = ?", department.DepartmentId)
	if err != nil {
		return nil, fmt.Errorf("employeesByDepartment %q: %v", departmentName, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var empl Employee
		if err := rows.Scan(&empl.EmployeeId, &empl.EmployeeName, &empl.DepartmentId); err != nil {
			return nil, fmt.Errorf("employeesByDepartment %q: %v", departmentName, err)
		}
		employees = append(employees, empl)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("employeesByDepartment %q: %v", departmentName, err)
	}
	return employees, nil
}
