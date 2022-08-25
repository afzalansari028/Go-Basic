package main

import "fmt"

type Shop struct {
	Item  string
	Price int
}

var total int
var totalBill int
var itemSlice []string
var quantitySlice []int

func main() {

	fmt.Println("*********Welcome the grocery shop*******")
	for {
		fmt.Println("---------Below items are available---------")

		fmt.Println("1:Sugar ₹40/kg")
		fmt.Println("2:Salt ₹20/kg")
		fmt.Println("3:Daal ₹120/kg")
		fmt.Println("4:Oil ₹180/l")
		fmt.Println("5:Ghee ₹220/l")
		fmt.Println("6:Quit press exit")
		fmt.Println("-------------------------------------------")

		Sugar := Shop{Item: "Sugar", Price: 40}
		Salt := Shop{Item: "Salt", Price: 20}
		Daal := Shop{Item: "Daal", Price: 120}
		Oil := Shop{Item: "Oil", Price: 180}
		Ghee := Shop{Item: "Ghee", Price: 350}

		fmt.Println("Enter item name for purchase or type exit:")
		var item string
		fmt.Scan(&item)
		if item == "exit" {
			break
		}
		fmt.Println("Enter quantity:")
		var quantity int
		fmt.Scan(&quantity)

		if item == Sugar.Item {
			payment := Sugar.Price * quantity
			total = total + payment
			bill := total
			totalBill = bill
			itemSlice = append(itemSlice, item)
			quantitySlice = append(quantitySlice, quantity)
		} else if item == Salt.Item {
			payment := Salt.Price * quantity
			total = total + payment
			bill := total
			totalBill = bill
			itemSlice = append(itemSlice, item)
			quantitySlice = append(quantitySlice, quantity)
		} else if item == Daal.Item {
			payment := Daal.Price * quantity
			total = total + payment
			bill := total
			totalBill = bill
			itemSlice = append(itemSlice, item)
			quantitySlice = append(quantitySlice, quantity)
		} else if item == Oil.Item {
			payment := Oil.Price * quantity
			total = total + payment
			bill := total
			totalBill = bill
			itemSlice = append(itemSlice, item)
			quantitySlice = append(quantitySlice, quantity)
		} else if item == Ghee.Item {
			payment := Ghee.Price * quantity
			total = total + payment
			bill := total
			totalBill = bill
			itemSlice = append(itemSlice, item)
			quantitySlice = append(quantitySlice, quantity)
		} else {
			fmt.Println("Please Enter valid name of product")
		}
	}
	fmt.Println("*************************************")
	fmt.Println("Purchased items are :", itemSlice)
	fmt.Println("Items Quantity are :", quantitySlice)
	fmt.Println("Total Bill :", totalBill)
	fmt.Println("**************************************")
}
