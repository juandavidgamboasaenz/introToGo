package pointers

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func StructPointers() {

	u := User{
		FirstName: "Fulanito",
		LastName:  "Cosme",
		Email:     "cosmefu@papaia.com",
		IsActive:  true,
	}

	fmt.Printf("my user is named: %+v\nand the property IsActivated is equal: %+v\n", u, u.IsActive)

	updateStructPointered(&u)

	fmt.Printf("my user is named: %v\nIts IsActive Property: %v", u.FirstName, u.IsActive)

	updateStructPointeredSimplier(&u)

	fmt.Println("If i use a simplier function this will work as expected")
	fmt.Println("You can ommit the * operator when referencing a value in a Struct!")

	fmt.Printf("my user is named: %v\nIts simplier implementation of IsActive Property: %v", u.FirstName, u.IsActive)

}

func updateStructPointered(u *User) {
	(*u).IsActive = !(*u).IsActive
}

func updateStructPointeredSimplier(u *User) {
	u.IsActive = !u.IsActive
}

func NoPointers() {
	//When doing programs you need to search what is a better option, use pointers
	//or returning a value. In this case the returned value its funnier, but creates
	//more memory, if using the pointer option it's more processing. All will depend on your application
	var v int = 19
	iv := increase(v)
	fmt.Println("v value is : ", v)
	fmt.Println("its returned value is :", iv)
}

func increase(v int) int {
	v++
	fmt.Println("v value in increase func is :", v)
	return v
}

func WithPointers() {
	var v int = 19
	fmt.Printf("This is the v value: %d\n", v)
	increasePointered(&v)
	fmt.Printf("This is the v new value: %d\nAnd this is the v pointer %v\n", v, &v)
	//As we can see alocating the new value of the data using pointers is not that hard
	//The thing is now how it can be accesed and changed
}

func increasePointered(v *int) {
	//We have put *int as the value that will be received by the function
	//in this case v is the variable and *int the pointer so when accessing
	//to the memory more easly we can use *v inside the method to properly play
	//with it's inner values
	fmt.Printf("this is the v pointer %v\n", v)
	*v++
}

func IntPointers() {
	v := 19

	var p1 *int
	var p2 = new(int)
	p3 := &v

	fmt.Printf("p1 int number with default value %d type is: %T \n", p1, p1)
	fmt.Printf("p2 int number with default value %d  type is: %T \n", p2, p2)
	fmt.Printf("p3 int number with default value %d type is: %T \n", p3, p2)
}

func IntDirection() {
	var v int = 19
	fmt.Println("Memory direction of v int var is :", &v)
}

func DereferencePointers() {
	var v int = 19
	var p *int = &v

	fmt.Printf("v variable is: %d\n", v)
	//%v formats a Direction
	fmt.Printf("v memory direction is: %v\n", &v)
	fmt.Printf("pointer p references to memory direction: %v\n", p)
	fmt.Printf("when dereferencing p pointer you get : %d \n", *p)
}
