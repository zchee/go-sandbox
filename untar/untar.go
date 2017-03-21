package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse() // get the arguments from command line

	sourcefile := flag.Arg(0)
	if sourcefile == "" {
		fmt.Println("Usage : untar sourcefile.tar")
		os.Exit(1)
	}

	file, err := os.Open(sourcefile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var r io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if r, err = gzip.NewReader(file); err != nil {

			fmt.Println(err)
			os.Exit(1)
		}
		defer r.Close()
	}

	tarball := tar.NewReader(r)
	// Extracting tarred files
	for {
		header, err := tarball.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		case tar.TypeReg:
			// handle normal file
			fmt.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarball)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			writer.Close()
		default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}

}
