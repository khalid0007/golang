package main

import "fmt"

const LoginToken string = "ghabbhhjd" // Public

type Person struct {
	Name string
	Age  int
}

type NilPersonException struct{}

func (exp *NilPersonException) Error() string {
	return "Can not access nil value."
}

func (p *Person) SayHello() error {
	if p == nil {
		return &NilPersonException{}
	}
	fmt.Printf("Hello from %v\n", p.Name)
	return nil
}

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45544511254451885
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

	// Null point check
	var khalid Person = Person{"Khalid Hassan Sk", 23}

	var khalidPtr *Person = &khalid

	khalidPtr = nil

	// err := SayHelloPtr(khalidPtr)
	err := khalidPtr.SayHello()

	if err != nil {
		fmt.Printf("Fail to say hello with the following reason : %s\n", err.Error())
	}

	khalid.SayHello()

}
