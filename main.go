package main

import (
  "crypto/sha256"
  "crypto/md5"
  "crypto/sha1"
  "golang.org/x/crypto/sha3"
  "fmt"
  "os"
  "log"
)

func main() {

  // Check if an argument is provided
  if len(os.Args) < 2 || len(os.Args) > 2 {
    fmt.Println("Usage: hashit <file to hash>")
    os.Exit(1)
  }

  // The first argument is always the program name,
  // so the actual input will be the second argument.
  inputFile := os.Args[1]

  data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
  output := fmt.Sprintf("Filename:\t%s\n", inputFile)
  output += fmt.Sprintf("Size:\t\t%d\n", len(data))  
  output += generateHash(data)
  fmt.Println(output)
}


func generateHash(data []byte) string {
   output := ""

   output += fmt.Sprintf("MD5:\t\t%x\n",md5.Sum(data)) 
   output += fmt.Sprintf("SHA1:\t\t%x\n",sha1.Sum(data)) 
   output += fmt.Sprintf("SHA256:\t\t%x\n",sha256.Sum256(data))
   output += fmt.Sprintf("SHA3(512):\t%x\n",sha3.Sum512(data))
  return output
}
