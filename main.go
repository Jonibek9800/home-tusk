package main

import "fmt"

type User struct {
	id				int64
	name			string
	surname			string
	login			string
	password		string
	sex				string
	age				int64
	userInfo		string
}

// func ChangeInfo(){

// }


// func main() {
// 	// NewUser := User{
// 	// 	id:				0,
// 	// 	name:			"",
// 	// 	surname:		"",
// 	// 	login:			"",
// 	// 	password:		"",
// 	// 	sex:			"",
// 	// 	age:			0,
// 	// 	userInfo:		"",
// 	// }
// 	var newUserSecond User
// 	newUserSecond.id =
// 	newUserSecond.id =
// 	newUserSecond.id =
// 	newUserSecond.id =
// 	newUserSecond.id =
	
	










// 	//var a int64 = 10
// 	a := int64 (10)
	
// }

type pur struct {
	  purchase 				[]int64
	  percentCashback     	int64
	  cash	       			int64	
}


func (new pur) cashback() int64 {
   var sum int64
   for _, item := range new.purchase{
	   sum += item
   }
   var cash int64
   cash = sum * new.percentCashback / 100
   return cash
}


func main(){
	purchases := pur{
		purchase:   []int64{123, 200, 520, 400, 500},
		percentCashback:    20,
		cash:        0,
	}
	purchases.cash = purchases.cashback()
	fmt.Println(purchases.cash)
}


