package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strings"
	"strconv"
)




func main(){

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done]")
		return
	}




		command:=os.Args[1]

		// fmt.Println(command)
		switch command{
		case "add":
			
			// todo add buy milk from store
			// os.Args[3:] = ["buy", "milk", "from", "store"]
			contentArray:=os.Args[2:]
			stringArray:=strings.Join(contentArray," ")
			
			fmt.Println(stringArray)

			file,err :=	os.OpenFile(
				"todo.csv",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, //append to end, if DNE then create, write-only mode
				0644, //perms owner RW / group R / others R
			)

			if err !=nil {
				panic(err)
			}

			// file closes when main() finishes
			defer file.Close()

			// CSV writer that writes to the file
			writer:=csv.NewWriter(file)
			defer writer.Flush()

			writer.Write([]string {stringArray})


		case "list":
			file,err:= os.Open("todo.csv")
			 
			if err!=nil{
				panic(err)
			}

			defer file.Close()

			reader :=csv.NewReader(file)

			records,err :=reader.ReadAll()

			if err!=nil{
				panic(err)
			}

			if len(records)==0 {
				fmt.Println("No Todos")
				return
			}

			for i,record := range records{

			fmt.Printf("%d. %s\n", i+1, record[0])
			}

		case "done":
			//read csv,recreate the todos and write to a new same name file
			if len(os.Args) < 3{
				fmt.Println("Usage: todo done <id>")
				return
			}
			id, err := strconv.Atoi(os.Args[2])
			if err != nil || id < 1 {
			fmt.Println("Invalid id")
			return
			}

			file,err := os.Open("todo.csv")

			if err!=nil {
			fmt.Println("No todos found")
			return
			}
			
			reader:=csv.NewReader(file)
			records,err :=reader.ReadAll()

			file.Close()

			if err!=nil {
				panic(err)
			}

			if id>len(records){
				fmt.Println("Todo not found")
				return
			}

			newRecords :=[][]string{}

			for i,record :=range records{

				if i+1!=id{
					newRecords=append(newRecords,record)
				}
			}

			file,err = os.Create("todo.csv")

			if err !=nil{
				panic(err)
			}
			writer:=csv.NewWriter(file)
			writer.WriteAll(newRecords)
			writer.Flush()
			file.Close()

			fmt.Println("Completed todo",id)
		

		default:
			fmt.Println("Invalid command")
		}


}