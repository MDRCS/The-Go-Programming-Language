package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {

	counts := make(map[string]int)
	counts_filename := make(map[string]string)
	if len(os.Args) > 1 {
		for _, filename := range os.Args[1:] {
			data,err := ioutil.ReadFile(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr,"Show the error when opening file : %v",err)
				continue
			}

			for _, line := range strings.Split(string(data),"\n") {
				counts[line]++
				counts_filename[line] = strings.TrimSuffix(filename,path.Ext("txt"))
			}
		}

		//Display Each line and the number of occurences
		for line, occ := range counts {
			if occ > 1 {
				fmt.Printf("Number of occurence of each line | occ : %d , line : %s \n",occ,line)
			}
		}

		//Display Each Line and the file where we got it
		for line, filename :=range counts_filename {
			fmt.Printf("This line %s got it from thus file %s \n",line,filename)
		}

	} else {

		newfile := 2

		for newfile>0 {

			data := bufio.NewScanner(os.Stdin)

			for data.Scan() {
				if data.Text() == " " {
					break
				}

				counts[data.Text()]++
			}

			newfile--
		}
		for line, occ := range counts {
			if occ > 1 {
				fmt.Printf("Number of occurence of each line | occ : %d , line : %s \n",occ,line)
			}
		}
	}



// go run duplicate_lines_v3.go file.txt file1.txt file2.txt
}
