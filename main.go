package main

import (
  "crypto/sha256"
  "crypto/md5"
  "crypto/sha1"
  "golang.org/x/crypto/sha3"
  "fmt"
  "os"
  "log"
  "io/fs"
  "path/filepath"

  "github.com/akamensky/argparse"
)

func main() {
  var outFile *os.File
  var outErr error

  parser := argparse.NewParser("hashit", "Create hash of a file / all files within a specified directory")

  verb := parser.Flag("v", "verbose", &argparse.Options{Help: "Enable verbose mode"})
  inFileOption := parser.String("i","input", &argparse.Options{Required: false, Help: "Name of file to hash"})
  outFileOption := parser.String("o","output", &argparse.Options{Required: false, Help: "Name of file to hash results to"})
  inDirOption := parser.String("d","directory", &argparse.Options{Required: false, Help: "Name of directory containing files to be hashed"})

  // Parse input
  err := parser.Parse(os.Args)
  if err != nil {
    fmt.Print(parser.Usage(err))
    os.Exit(1)
  }

  if len(*inFileOption) == 0 && len(*inDirOption) == 0 {
    fmt.Println("You must specify a file or a directory to be hashed")
    os.Exit(1)
  }

  if len(*outFileOption) > 0 {
    outFile, outErr = os.OpenFile(*outFileOption, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
    if outErr != nil {
      fmt.Errorf("create / append file: %w", outErr)
    }
  }

  switch {
  case len(*inFileOption) > 0 :
    output := processFile(*inFileOption)
    if *verb {
      fmt.Println(output)
    }
    if len(*outFileOption) > 0 {
      outFile.WriteString(output + "\n")
    }
  case len(*inDirOption) > 0 :
    filesystem := os.DirFS(*inDirOption)
    fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
      if err != nil {
        log.Fatal(err)
      }
   
      if !d.IsDir() { 
        output := processFile(filepath.Join(*inDirOption,path))
        if *verb {
          fmt.Println(output)
        }
      
        if len(*outFileOption) > 0 {
          outFile.WriteString(output + "\n")
        }
      }
      return nil
    })
  }
  outFile.Close()
}

func processFile(filename string) string {
    data, err := os.ReadFile(filename)
    if err != nil {
      log.Fatal(err)
    }
    output := fmt.Sprintf("Filename:\t%s\n", filename)
    output += fmt.Sprintf("Size:\t\t%d\n", len(data))  
    output += generateHash(data)

    return output 
}

func generateHash(data []byte) string {
   output := ""
   output += fmt.Sprintf("MD5:\t\t%x\n",md5.Sum(data)) 
   output += fmt.Sprintf("SHA1:\t\t%x\n",sha1.Sum(data)) 
   output += fmt.Sprintf("SHA256:\t\t%x\n",sha256.Sum256(data))
   output += fmt.Sprintf("SHA3(512):\t%x\n",sha3.Sum512(data))
  return output
}
