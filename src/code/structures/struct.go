package structures

/*
Consider go structs like C++ structures. Since go is function programming language, for us to define custom types, we
can use struct.

type <STRUCT NAME> struct{
   KEY TYPE
}
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
