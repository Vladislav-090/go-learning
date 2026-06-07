//

package main

import "fmt"


type User struct {
	Name string
	Age int
}

func IncreaseAge(u *User, years int) {
	u.Age = u.Age + years
}

func main() {
	user := User {
		Name : "Vladislav",
		Age: 30,
	}
	IncreaseAge(&user, 5)

	fmt.Println("hello my name is :", user.Name, "and my age is :", user.Age)

}