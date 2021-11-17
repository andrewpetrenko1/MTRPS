package main

import (
  "bufio"
  "fmt"
  "os"
  "flag"
  "main/treeSort"
  "strings"
)

func getFileLines(fileName string) []string {
  file, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  arrInp := make([]string, 0, 10)

  for scanner.Scan() {
    arrInp = append(arrInp, scanner.Text())
  }
  return arrInp
}

func writeToFile(fileName string, content []string) bool {
  file, err := os.Create(fileName) 
  writer := bufio.NewWriter(file) 
  if err != nil { 
    fmt.Println(err) 
    return false
  } 
  defer file.Close()

  for _, data := range content {
    writer.WriteString(data)
    writer.WriteString("\n")  
  } 
  writer.Flush()
  return true
}

func getUserInput() []string {
	arrInp := make([]string, 0, 10)
  fmt.Println("Enter line:")
	for {
		var input string
    fmt.Scanln(&input)
		if input == "" {
			break
		}
		arrInp = append(arrInp, input)
	}
  return arrInp;
}

func main() {
  inputFile := flag.String("i", "testFile.txt", "Input file")
  outputFile := flag.String("o", "testOutFile.txt", "Input file")
  header := flag.Bool("h", false, "The first line is a header tha must be ignored during sorting but included in the output.")
  nLineValue := flag.Int("f", 0, "Sort input lines by value number N.")
  reverse := flag.Bool("r", false, "Sort input lines in reverse order.")
  flag.Parse()

  fmt.Println("Entered options:")
  fmt.Println("inputFile:", *inputFile)
  fmt.Println("outputFile:", *outputFile)
  fmt.Println("header:", *header)
  fmt.Println("nLineValue:", *nLineValue)
  fmt.Println("reverse:", *reverse)
  fmt.Println("---------------")

  fileLines:= getFileLines(*inputFile)
  headerSl := ""

  if *header {
    headerSl = fileLines[0]
    fileLines = fileLines[1:len(fileLines)]
  }
  if *nLineValue != 0 {
    fileLines = fileLines[*nLineValue:len(fileLines)]
  }

  sortedResult:=treeSort.Treesort(fileLines, *reverse)
  
  if *header && headerSl != "" {
    sortedResult = append([]string{headerSl}, sortedResult...)
  }
  fmt.Println("Sorted: \n" + strings.Join(sortedResult, "\n"))

  if writeToFile(*outputFile, sortedResult) {
    fmt.Println("Output at " + *outputFile)
  }
}
