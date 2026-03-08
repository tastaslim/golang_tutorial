package structures

/*
Consider go structs like C++ structures. Since go is function programming language, for us to define custom types, we
can use struct.

type <STRUCT NAME> struct{
   KEY TYPE
}

It is not mandatory to pass all fields in struct definition. The Keys for which values are not assigned, those
will be defined with default null values
*/

type Person struct {
	Name       string
	Age        int
	Address    string
	IsResident bool
	Salary     float32
}

func helper(name string, age int, address string, salary float32, isResident bool) Person {
	person := Person{
		Name:       name,
		Age:        age,
		Address:    address,
		Salary:     salary,
		IsResident: isResident,
	}
	return person
}

func PracticeStruct(person Person) Person {
	return helper(person.Name, person.Age, person.Address, person.Salary, person.IsResident)
}

// Let's make below method part of Person struct. Below method is not part of struct Person
func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() string { // func (p Person) GetName() string ==> This is also fine
	return p.Name
}
