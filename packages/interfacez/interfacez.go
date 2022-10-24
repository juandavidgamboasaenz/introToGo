package interfacez

import "fmt"

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type snake struct {
	Type      string
	Poisonous bool
}

func (s snake) description() string {
	return fmt.Sprintf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound: %s", c.Sound)
}

func Interfacez() {

	a := snake{Poisonous: true}
	fmt.Println(a.description())

	b := cat{Sound: "Nya!!!"}
	fmt.Println(b.description())
}
