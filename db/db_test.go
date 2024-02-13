package db_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/db"
)

func TestCRUD(t *testing.T) {
	// Connect and get a database handle.
	dbConnection, err := db.CreateConnection()
	require.NoError(t, err)

	// Clear all database content.
	db.DeleteAllDataEmployee(dbConnection)
	db.DeleteAllDataDepartment(dbConnection)

	// CREATE
	// Create a new department.
	deptName := "Biologi"
	deptIdDb, err := db.CreateDepartment(dbConnection, deptName)
	require.NoError(t, err)

	// Get the department name by ID.
	deptDb, err := db.DepartmentNameById(dbConnection, deptIdDb)
	require.NoError(t, err)
	require.Equal(t, deptName, deptDb.DepartmentName)

	// Create a new employee.
	emplName := "Rasuna Said"
	emplIdDb, err := db.CreateEmployee(dbConnection, emplName, deptIdDb)
	require.NoError(t, err)

	// Get the employee name and department by ID.
	emplDb, err := db.EmployeeNameById(dbConnection, emplIdDb)
	require.NoError(t, err)
	require.Equal(t, emplName, emplDb.EmployeeName)
	require.Equal(t, emplIdDb, emplDb.DepartmentId)

	// UPDATE
	// Update existing department data.
	deptName = "Biologi Molekuler"
	deptIdDb, err = db.UpdateDepartment(dbConnection, deptName, deptIdDb)
	require.NoError(t, err)

	// Get the department name by ID.
	deptDb, err = db.DepartmentNameById(dbConnection, deptIdDb)
	require.NoError(t, err)
	require.Equal(t, deptName, deptDb.DepartmentName)

	// Update existing employee data.
	emplName = "Rasuna Aulia"
	emplIdDb, err = db.UpdateEmployee(dbConnection, emplName, deptName, emplIdDb)
	require.NoError(t, err)

	// Get the employee name and department by ID.
	emplDb, err = db.EmployeeNameById(dbConnection, emplIdDb)
	require.NoError(t, err)
	require.Equal(t, emplName, emplDb.EmployeeName)
	require.Equal(t, emplIdDb, emplDb.DepartmentId)

	// DELETE
	// Delete existing employee data.
	deletedRowNum := 1
	var deletedRowNumInt64 = int64(deletedRowNum)
	rowEmplAffectedNum, err := db.DeleteEmployeeById(dbConnection, emplIdDb)
	require.NoError(t, err)
	require.Equal(t, deletedRowNumInt64, rowEmplAffectedNum)

	// Delete existing department data.
	rowDeptAffectedNum, err := db.DeleteDepartmentById(dbConnection, deptIdDb)
	require.NoError(t, err)
	require.Equal(t, deletedRowNumInt64, rowDeptAffectedNum)

}
