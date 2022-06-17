package orm

import (
	"database/sql"
	"os"
	"strings"
	"venkin/logger"
)

// GlobalDB global database handle
var GlobalDB *sql.DB

// global preprocessing sql collection
var preprocessSqlMap = make(map[string]*sql.Stmt)

// InitGlobalDB initialization global database handle
func InitGlobalDB(db *sql.DB) {
	GlobalDB = db
}

// InitGlobalSqlMap global sql statement collection
func InitGlobalSqlMap() {

	// Get the current directory of the project
	currentDir, _ := os.Getwd()

	// Spell out the agreed sql configuration file directory path
	sqlConf := currentDir + "/sqlmap"

	// Check if a directory exists in the agreed path
	if !fileExist(sqlConf) {
		logger.LogRun("It is detected that the sqlmap directory does not exist in the project root directory, and the initialization of some ORM function parameters has not been completed")
		logger.LogRun("At this point, an error will occur when using the functions provided by the framework to operate the database")
		return
	}
	// If the directory exists, retrieve the .ini file in the directory
	sqlmap, _ := os.Open(sqlConf)

	// Deferred execution closes directory reads
	defer sqlmap.Close()

	// Read all files in a directory
	f, err := sqlmap.Readdir(-1)

	// exception handling
	if err != nil {
		logger.LogRun("An exception occurred while reading the sqlmap directory")
		return
	}

	// Iterate over the read collection of files
	for _, fInfo := range f {
		// Skip if file is a directory
		if fInfo.IsDir() {
			continue
		}

		// If it is a file, determine whether it is an .ini file
		if strings.HasSuffix(fInfo.Name(), ".ini") {
			// If it is to start reading file information
			readIni(sqlConf + "/" + fInfo.Name())
		}
	}
}

// Check if a folder exists in a given path
func fileExist(fileUrl string) bool {
	f, err := os.Stat(fileUrl)
	if err != nil {
		return false
	}
	return f.IsDir()
}
