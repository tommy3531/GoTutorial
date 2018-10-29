package main
import (
	"fmt"
)

type user struct {
	name string
	email string
	ext int
	privileged bool
}

type admin struct {
	person user
	level string
}

func main() {

	// var tom user
	// fmt.Println(tom)

	// jerry := user{
	// 	name: "Jerry",
	// 	email: "jerry@hotmail.com",
	// 	ext: 123,
	// 	privileged: true,
	// }
	// fmt.Println(jerry)

	bob := admin {
		person: user{
			name: "Tom",
			email: "tom@hotmail.com",
			ext: 123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(bob)

}