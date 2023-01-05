package main

type person struct {
}

type human interface {
	speak()
}

func (p *person) speak() {

}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{}

	saySomething(&p)
	//saySomething(p) ❌
}
