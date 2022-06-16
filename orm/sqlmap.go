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

	kMap := make(map[string]string)

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
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if currentSection != "" {
				sqlMap[currentSection] = kMap
				kMap = make(map[string]string)
			}
			currentSection = line[1 : len(line)-1]
			if err == io.EOF {
				sqlMap[currentSection] = kMap
				break
			}
			continue
		}
		if line != "" {
			n := strings.Index(line, "=")
			kMap[line[:n]] = line[n+1:]
		}
		if err == io.EOF {
			if currentSection != "" {
				sqlMap[currentSection] = kMap
			}
			break
		}
	}
}

func getSqlMapValue(section, key string) (string, bool) {
	s, ok := sqlMap[section][key]
	if ok {
		return s, true
	} else {
		return "", false
	}
}
