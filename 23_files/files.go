package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("file name", fileInfo.Name())
	// fmt.Println("file or folder", fileInfo.IsDir())
	// fmt.Println("file size", fileInfo.Size())
	// fmt.Println("file permission", fileInfo.Mode())
	// fmt.Println("file modified at", fileInfo.ModTime())

	// read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// //Not the regualr use
	// //Load all data at once
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// // read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// // negative or 0 give all dir and file
	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// // create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("hi go")
	// // f.WriteString("nice languge")

	// bytes := []byte("Hello Golang")

	// f.Write(bytes)

	// // read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {

	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()

	// fmt.Println("Written to new file succesfully")

	// Delete a File

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File is deleted")

}
