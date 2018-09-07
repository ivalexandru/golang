// https://shapeshed.com/copy-a-file-in-go/

// To copy a file is therefore a case of glueing together a few functions from the os package. The process is

// Open the file that should be copied
// Read the contents
// Create and open the file that the contents should be copied into
// Write to the new file
// Close both files

package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// 	The Open function from the os module is used to read a file from disk.
	// A defer statement is used to close the file once the script has finished all other execution.

	from, err := os.Open("./sursa.txt")
	//   if err != nil {
	//     log.Fatal(err)
	//   }
	defer from.Close()

	//   The OpenFile function is used to open a file.
	//   The first argument is the name of the file to be opened or created if it does not exist.
	//   The second argument represents the flags to be used on the file. In this case it is read and write and should be created if it does not exist. Finally the permissions on the file are set.
	// Another defer statement is used to close this file after other execution has finished.

	to, err := os.OpenFile("./folderDestinatie/creatAutomat.txt", os.O_RDWR|os.O_CREATE, 0666)
	//   if err != nil {
	//     log.Fatal(err)
	//   }
	defer to.Close()

	//   The Copy function is then used from the io package. This copies from the source file and writes it to the destination.

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}
