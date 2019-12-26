package main

import (
	"log"
	"sync"
  "os"
  "fmt"
	"strings"
	"path/filepath"
	"bufio"
)

type chunk struct {
	bufsize int
	offset  int64
}

var separators = [3]string{". ", "! ", "? "}
var lineEndings = [3]string{".", "!", "?"}
var inputs = os.Args[1]
var outputs = os.Args[2]

func find(counter int, searchArea string) int {
	inputString := searchArea[:]
	stringLength := len(inputString)
  for _, separator := range separators {
    counter += strings.Count(inputString, separator)
	}
   for _, lineEnding := range lineEndings {
	 	if string(inputString[stringLength - 1]) == lineEnding {
	 		counter++
	 	}
	 }
  return counter
}

func writeFile(fileName string, counter int) {
    err := os.Mkdir(outputs, 0644)
	  f, err := os.Create(outputs + "/" + strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".res")
		if err != nil {
			fmt.Println(err)
			return
		}
    defer f.Close()
    f.Sync()

    w := bufio.NewWriter(f)
    _, err = w.WriteString(fmt.Sprintf("%d", counter))
		if err != nil {
			fmt.Println(err)
			return
		}

    w.Flush()
}


func readFile(fileName string) {
  counter := 0
  //const BufferSize = 128
	fmt.Println(fileName)
  file, err := os.Open(inputs + "/" + fileName)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  //fileinfo, err := file.Stat()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if string(text) == "" {
      return
		}
		counter = find(counter, text)
	}
	writeFile(fileName, counter)

}
  // filesize := int(fileinfo.Size())
  // // Number of go routines we need to spawn.
  // concurrency := filesize / BufferSize
  // // buffer sizes that each of the go routine below should use. ReadAt
  // 	// returns an error if the buffer size is larger than the bytes returned
  // 	// from the file.
  // chunkArray:= make([]chunk, concurrency)
	//
  // 	// All buffer sizes are the same in the normal case. Offsets depend on the
  // 	// index. Second go routine should start at 100, for example, given our
  // 	// buffer size of 100.
 	// for i := 0; i < concurrency; i++ {
  // 		chunkArray[i].bufsize = BufferSize
  // 		chunkArray[i].offset = int64(BufferSize * i)
  // 	}
	//
  // 	// check for any left over bytes. Add the residual number of bytes as the
  // 	// the last chunk size.
  // 	if remainder := filesize % BufferSize; remainder != 0 {
  // 		c := chunk{bufsize: remainder, offset: int64(concurrency * BufferSize)}
  // 		concurrency++
  // 		chunkArray = append(chunkArray, c)
  // 	}
	//
  // var wg sync.WaitGroup
	// var m sync.Mutex
  // wg.Add(concurrency)
	//
  // for i := 0; i < concurrency; i++ {
  //   go func(chunkArray[]chunk, i int) {
  //     defer wg.Done()
	//
  //   chunk := chunkArray[i]
  //     buffer := make([]byte, chunk.bufsize)
  //     _, err := file.ReadAt(buffer, chunk.offset)
	// 		m.Lock()
	// 		counter = find(counter, buffer)
	// 		m.Unlock()
  //     if err != nil{
  //       fmt.Println(err)
  //       return
  //     }
  //     //fmt.Println("bytes read, string(bytestream): ", bytesread)
  //     //fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	// 		//fmt.Println("Sentences found in file ", counter)
  //   }(chunkArray, i)


//   wg.Wait()
// 	writeFile(fileName, counter)
//   //fmt.Println("Sentences found in file ", counter)
// }

func readMultipleFiles (files []os.FileInfo) {
	var wg sync.WaitGroup
	var m sync.Mutex

 //map[string][]byte
	filesLength := len(files)
	//contents := make(map[string][]byte, filesLength)
	wg.Add(filesLength)

	for _, file := range files {
		go func(file string) {

			m.Lock()
			readFile(file)
			m.Unlock()
			wg.Done()
		}(file.Name())
	}

	wg.Wait()
  fmt.Println("Total number of processed files:", filesLength)
	//return contents
}

func main() {
  reader, err := os.Open(inputs)

  if err !=nil {
    log.Fatal(err)
  }
  files, err := reader.Readdir(-1)
  reader.Close()
  if err !=nil {
    log.Fatal(err)
  }

  readMultipleFiles(files)
}
