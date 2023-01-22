package main

import (
	"fmt"
	"os"
)
type bill struct{
	name string
	item map[string]float64
	tip float64
}

//make new bill
func makenewbill(name string) bill{
	b := bill{
		name: name,
		item: map[string]float64{},
		tip: 0,
	}
	return b
}

func (b bill) format() string{
	fs := "bill breakdown \n"
	var total float64 = 0
	
	for k,v := range b.item{
		fs += fmt.Sprintf("%-25v ...$%v \n",k+":",v)
		total += v
	}
	//add tip
	fs += fmt.Sprintf("%-25v ...%0.2f", "tip:",b.tip)
	//adding total
	fs += fmt.Sprintf("\n%-25v ...%0.2f", "total:",total+b.tip)
	return fs
}
//update tip need to use the pointer to the tip value
func (b *bill) updateTip(tip float64){
	b.tip = tip
}

func (b *bill) addItem(n string,p float64){
	b.item[n] = p
}

//save a bill 
func (b *bill) save(){
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt",data,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}