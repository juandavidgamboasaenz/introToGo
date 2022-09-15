package structures

import (
	"fmt"
)

type User struct {
	age            int
	name, lastName string
}

// I just created a new interface, there we'll allocate the functions
// Needed to our Cart interface

type Cart interface {
	//Here we create the Auth func
	Permissions() int
	Name() string
}

type Admin struct {
	name    string
	authLVL int
}

// type Editor struct {
// 	name    string
// 	authLVL int
// }

func MyStructures() {
	u := User{
		age:      69,
		name:     "David",
		lastName: "Gamboa",
	}

	u2 := User{
		420, "juansho", "juanshoso",
	}

	u3 := new(User)
	fmt.Println(u.name, u2, "Pointer Data", u3, "Inner Data", *u3, "Memory Direction", &u3)

}

//Here in the RECIEVER we implement into an Admin struct type
//the Auth function we declared earlyer

func (admin Admin) Permissions() int {
	return admin.authLVL
}

func (admin Admin) Name() string {
	return admin.name
}

func auth(cart Cart) string {
	if cart.Permissions() > 4 {
		return cart.Name() + " tienen permisos de admin"
	}
	return cart.Name() + " no tiene permisos de admin"
}

func ReceiverInterface() {
	admin := Admin{
		name:    "David",
		authLVL: 4,
	}
	fmt.Println(auth(admin))
}
