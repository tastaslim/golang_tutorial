package interfaces

type Animal interface {
	speak() string
}

type Cat struct {
	Sound string
}

type Dog struct {
	Sound string
}

type Cow struct {
	Sound string
}

func (c Cat) speak() string {
	return c.Sound
}

func (d Dog) speak() string {
	return d.Sound
}

func (c Cow) speak() string {
	return c.Sound
}

func MakeSound(a Animal) string {
	return a.speak()
}
