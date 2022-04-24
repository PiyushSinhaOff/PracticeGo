package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to be fixed"

	file, err := os.Create("./myTestingFile.txt")
	defer file.Close()
	CheckNilError(err)

	length, err := io.WriteString(file, content)
	CheckNilError(err)

	fmt.Println("length: ", length)
	readFile("./myTestingFile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	CheckNilError(err)
	fmt.Println("dataByte: ", dataByte)
	fmt.Println("dataByte: ", string(dataByte))
}

func CheckNilError(err error){
	if err != nil {
		panic(err)
	} 
}