package orm

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Read .ini file utility function
func readIni(fileName string) {
	// open file
	fo, err := os.Open(fileName)
	if err != nil {
		panic("Error opening file")
	}

	defer fo.Close()

	// read file
	buf := bufio.NewReader(fo)
	var currentSection = ""
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic("Error reading file")
			}
		}
		// remove spaces
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = line[1 : len(line)-1]
			if err == io.EOF {
				break
			}
			continue
		}
		if line != "" {
			n := strings.Index(line, "=")
			preprocessSqlMap[currentSection+line[:n]], _ = GlobalDB.Prepare(line[n+1:])
		}
		if err == io.EOF {
			break
		}
	}
}
