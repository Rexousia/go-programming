package main

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
