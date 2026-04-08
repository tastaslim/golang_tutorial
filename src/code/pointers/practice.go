package pointers

import (
	"fmt"
	"time"
)

func PointerPractice() {
	tas := 10
	pointerAddress := &tas
	fmt.Printf("The address of tas %v\n", pointerAddress)
	fmt.Printf("The value of tas is %v\n", *pointerAddress)
}

func PassByReference(referencedVariable *int) {
	*referencedVariable = 10
}

func SlicesOperations(childrenAge []int) {
	//childrenAge = {1,2,3}
	childrenAge[0] = 4
	// childrenAge = {4,2,3}
}

type Employee struct {
	Name        string
	Age         int
	Position    string
	Address     string
	Salary      float32
	CreatedTime time.Time
}

func PassByReferenceOnCustomType(emp *Employee) {
	emp.Name = "Arif"
	emp.Age++
}

type Animal struct {
	Name  string
	Age   int
	Breed string
}

func getAnimalInfo(animal *Animal) Animal {
	return *animal
}

func main() {
	cow := Animal{
		Name:  "Cow",
		Age:   12,
		Breed: "JERSEY",
	}

	result := getAnimalInfo(&cow)
	fmt.Println(result)
}
