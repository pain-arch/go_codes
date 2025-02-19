package main

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the simple sum app")
}

func getUserName ()  string{
	var name string
	fmt.Println("Enter your name please -")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter the first Number -")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second Number -")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(n1 int , n2 int) int {
	sum := n1 + n2
	return sum
}

func display(name string, sum int) {
	fmt.Println("\nHello - ",name)
	fmt.Println("The sum of the number is : ",sum)
}

func endingMessage() {
	fmt.Println("Thank you for using our application.")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1,num2)
	display(name,sum)
	endingMessage()
}