package main

import (
        "fmt"
        "strconv"
        "strings"
)

func luhnCheck(cardNumber string) bool {
        var sum int
        alt := false

        for i := len(cardNumber) - 1; i > -1; i-- {
                n, _ := strconv.Atoi(string(cardNumber[i]))
                if alt {
                        n *= 2
                        if n > 9 {
                                n = (n % 10) + 1
                        }
                }
                sum += n
                alt = !alt
        }

        return sum%10 == 0
}

func main() {
        var cardNumber string
        fmt.Print("Enter credit card number: ")
        fmt.Scanln(&cardNumber)

        cardNumber = strings.ReplaceAll(cardNumber, " ", "")
        cardNumber = strings.ReplaceAll(cardNumber, "-", "")

        if luhnCheck(cardNumber) {
                fmt.Println("Valid credit card number")
        } else {
                fmt.Println("Invalid credit card number")
        }
}