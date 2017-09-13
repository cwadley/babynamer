package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Name struct {
	State string
	Sex string
	Year string
	Name string
	Count string
}

func main() {
	echo "hello"
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertNames(sqlCon sql.DB, names []Name) {

}

func ProcessFiles(dataDir string) []Name {
	files := GetFiles(dataDir)
	var names []Name
	for _, currFilePath := range files {
		rawFile = ReadFile(currFilePath)
		names = append(names, ParseData(rawFile))
	}

	return names
}

func GetFiles(dirPath string) []FileInfo {
	files, err := ioutil.ReadDir(dirPath)
	CheckErr(err)
	return files
}

func ReadFile(currFile File) []string {
	var lines []string
	reader = bufio.NewReader(currFile)
	var currString string, err error
	currString, err = reader.ReadString("\n")
	for err == nil {
		lines = append(lines, currString)
		currString, err = reader.ReadString("\n")
	}

	return lines
}

func ParseData(rawStrings []string) []Name {
	var names []Name

	for _, v := range rawStrings {
		splitString := Split(v, ",")
		currName := Name{State:splitString[0], Sex:splitString[1], Year:splitString[2], Name:splitString[3], Count:splitString[4]}
		names = append(names, currName)
	}

	return names
}