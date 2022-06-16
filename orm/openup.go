package orm

import (
	"database/sql"
	"errors"
	"unsafe"
)

func getStmt(section, key string) (*sql.Stmt, error) {
	// get sql
	sqlStr, ok := getSqlMapValue(section, key)
	if !ok {
		return nil, errors.New("SQL Does Not Exist")
	}
	// Preprocessing to prevent sql injection
	stmt, err := GlobalDB.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

// SelectOne Query a single data tool
func SelectOne[T interface{}](section, key string, args ...any) (*T, error) {

	stmt, err := getStmt(section, key)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(args...)
	defer stmt.Close()

	// type conversion
	// Converting the built-in type of golang to the framework copy is convenient for subsequent operations
	copyRow := (*Row)(unsafe.Pointer(row))

	// Invokes an overridden Scan of the framework copy type
	err = copyRow.Scan()

	// A collection of database fields to query
	columns := copyRow.rows.rowsi.Columns()
	// Database field result collection
	lastcols := copyRow.rows.lastcols

	// Generic type pointer definition
	// reflection assignment
	obj := reflectTypeObj[T](columns, lastcols)

	return obj, nil
}

// SelectList Query multiple pieces of data tool
func SelectList[T interface{}](section, key string, args ...any) ([]*T, error) {

	stmt, err := getStmt(section, key)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	// type conversion
	// Converting the built-in type of golang to the framework copy is convenient for subsequent operations
	copyRows := (*Rows)(unsafe.Pointer(row))
	defer copyRows.Close()

	var s []*T

	// cyclic read
	for copyRows.Next() {

		// A collection of database fields to query
		columns, err := copyRows.Columns()
		if err != nil {
			return nil, err
		}
		// Database field result collection
		lastcols := copyRows.lastcols

		// Generic type pointer definition
		// reflection assignment
		obj := reflectTypeObj[T](columns, lastcols)
		s = append(s, obj)
	}

	return s, nil
}

// InsertInto Insert info tool
// If the operation is successful and there is an auto-incrementing primary key inserted into the table
// the auto-incrementing primary key will be returned as the first parameter
// and if there is no auto-incrementing primary key, the number of affected rows will be returned
// use 0 on failure
func InsertInto(section, key string, args ...any) (int64, error) {
	return insertAndUpdateAndDelete(section, key, 1, args...)
}

// UpdateInfo Update Information Tool
func UpdateInfo(section, key string, args ...any) (int64, error) {
	return insertAndUpdateAndDelete(section, key, 0, args...)
}

// DeleteInfo delete information tool
func DeleteInfo(section, key string, args ...any) (int64, error) {
	return insertAndUpdateAndDelete(section, key, 0, args...)
}

// insert update delete operation common part tool
// The sign parameter determines whether to attempt to return the primary key
// If an exception occurs in obtaining the primary key and the number of affected rows
// but the data operation does not fail, the first parameter will be supplemented with 1 to indicate that the operation is successful and the error returns nil
func insertAndUpdateAndDelete(section, key string, sign int8, args ...any) (int64, error) {

	stmt, err := getStmt(section, key)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, nil
	}

	var id int64

	if sign == 1 {
		// Auto-increment primary key acquisition
		id, err = res.LastInsertId()
		if err == nil {
			return id, err
		}
	}

	// If the primary key acquisition fails, get the number of affected rows
	id, err = res.RowsAffected()

	if err != nil {
		return 1, nil
	}

	return id, nil
}
