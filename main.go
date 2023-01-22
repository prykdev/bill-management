package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	mybill := createBill() 
	promptOptions(mybill)

}
func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("choose option a- add item, s - save bill, t - add tip", reader)

	switch opt {
	case "a": 
	name, _ := getInput("Item Name: ",reader)
	price,_ := getInput("Item Price: ",reader)
	p,err := strconv.ParseFloat(price,64) //parse to float64
	if err!= nil{
		fmt.Println("the price must be a number")
		promptOptions(b)
	}
	b.addItem(name,p)
	promptOptions(b)
	fmt.Println("Item added -",name,price)
	case "t":
		tip,_ := getInput("Enter Tip Ammount ($): ",reader)
		t,err := strconv.ParseFloat(tip,64) //parse to float64
	if err!= nil{
		fmt.Println("the tip must be a number")
		promptOptions(b)
	}
	b.updateTip(t)
	fmt.Println("Tip added",tip)
	promptOptions(b)
	case "s":
		b.save()

		fmt.Println("You saved the bill",b.name)
	default:
		fmt.Println("that was not a valid option ...")
		promptOptions(b)
	}
	

}
func getInput(prompt string,r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input,err := r.ReadString('\n')
	return strings.TrimSpace(input),err
	
}
func createBill() bill{
	reader := bufio.NewReader(os.Stdin)
	name ,_ := getInput("Create a new bill",reader)

	b:= makenewbill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}
