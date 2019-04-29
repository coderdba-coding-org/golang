package main

import (  
    "fmt"
)

// single return value
func calculateBill(price, no int) int {  
    var totalPrice = price * no
    return totalPrice
}

// multiple return value
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}

func main() {  

    // single return value function
    price, no := 90, 6
    totalPrice := calculateBill(price, no)
    fmt.Println("Total price is", totalPrice)

    // multiple return value function
    area, perimeter := rectProps(10.8, 5.6)
    fmt.Printlf("Area %f Perimeter %f \n", area, perimeter) 

}
