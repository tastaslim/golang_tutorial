package main

import (
	"fmt"

	// "github.com/tastaslim/golang_tutorial/src/code/arrays"
	// "github.com/tastaslim/golang_tutorial/src/code/conditionals"
	// loops "github.com/tastaslim/golang_tutorial/src/code/loops"
	shape "github.com/tastaslim/golang_tutorial/src/code/interfaces"
	// "github.com/tastaslim/golang_tutorial/src/code/maps"

	// variable "github.com/tastaslim/golang_tutorial/src/code/variables"
	// "github.com/tastaslim/golang_tutorial/src/code/takinginput"
	// "github.com/tastaslim/golang_tutorial/src/code/dataTypes"
	// "github.com/tastaslim/golang_tutorial/src/code/functions"
	// "github.com/tastaslim/golang_tutorial/src/code/pointers"
	// "github.com/tastaslim/golang_tutorial/src/code/structures"
	"github.com/tastaslim/golang_tutorial/src/pkg/logger"
)

func main() {
	logger.Init()
	logger.Info("Service started")
	// arrays.MorePractice()
	// fmt.Println(arrays.Factorial(5))
	// conditionals.ConditionalOperations()
	// conditionals.SwitchStatement()
	// loops.LoopOperations()
	// person := structures.Person{
	// 	Name:       "Taslim",
	// 	Age:        20,
	// 	Address:    "NPP",
	// 	Salary:     100,
	// 	IsResident: true,
	// }
	// fmt.Println(structures.PracticeStruct(person))
	// maps.MapsPractice()
	// arrays.SlicesPractice()
	// takinginput.TakeInput()

	// var tas *string
	// fmt.Println(&tas)

	// ans := arrays.VariadicInput(1, 2, 3, 4, 5)
	// arr2 := []int{1, 2, 3, 4, 5}
	// ans2 := arrays.VariadicInput(arr2...)
	// fmt.Println(ans2)
	// fmt.Println(ans)
	// variables.TestVariableShadowing()
	// dataTypes.TestDataTypes()
	fmt.Println()
	// arrays.SlicesGrowth()
	// logger.Info("msg=The Operation has completed")
	// summation, different := functions.TupleSum(1, 2)
	// fmt.Println(summation, different)
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// pairSum := arrays.PairSum(arr, 7)
	// fmt.Println(pairSum)
	// arr2 := []int{1, 2, 3, 10, 12}
	// fmt.Println(arrays.UnionArray(arr, arr2))

	// fmt.Println(arrays.ElementWithMaxFrequency([]int{1, 2, 3, 1, 1, 2, 3, 3, 3, 4, 1, 5, 3}))
	// fmt.Println(arrays.SecondLargestElement([]int{12, 4, 12}))
	// pointers.PointerPractice()
	// a := 5
	// pointers.PassByReference(&a)
	// employeer := pointers.Employee{
	// 	Name:        "Taslim",
	// 	Age:         26,
	// 	Address:     "NPP",
	// 	Salary:      12,
	// 	Position:    "SSE",
	// 	CreatedTime: time.Now(),
	// }
	// pointers.PassByReferenceOnCustomType(&employeer)
	// fmt.Println(employeer)
	// 	planterEssentials := structures.PlanetEssentials{Water: true, Fire: false, Air: false, Soil: true, Space: true}
	// 	planet := structures.Planet{Name: "Mercury", PlanetEssentials: planterEssentials}
	// 	fmt.Println(planet.HasAir())
	// 	fmt.Println(planet.HasSoil())
	// arr := []int{1, 2, 3, 4, 5}
	// maxi := arrays.MaximumElement(arr)
	// fmt.Println(maxi)
	// variable.TestConstant()
	// loops.EvenNumbers(20)
	// loops.BreakOnCondition(10)
	// fmt.Println(loops.ReturnOnCondition("Tas"))
	// finalArray := arrays.InitializeArray(2)
	// fmt.Println(finalArray)
	// 	employees := []string{"Asad", "Diksha", "Rohan", "Raghav"}
	// 	targetName := "Asad1"
	// 	exists := arrays.SlicesCheck(employees, targetName)
	// 	fmt.Println(exists)
	// arr := []int{1, 2, 3, 4, 1, 2, 1, 1, 1, 1, 2}
	// fmt.Println(maps.CountFrequency(arr))
	// arr2 := []int{1, 2, 5}
	// arr3 := []int{2, 5, 1}
	// fmt.Println(maps.UnionArray(arr, arr2))
	// fmt.Println(maps.CheckEqual(arr3, arr2))
	// fmt.Println(maps.IsSubset(arr2, arr3))
	// circle := shape.Circle{Radius: 10}
	// square := shape.Square{Length: 10, Breadth: 10}
	// fmt.Println("The Area of Circle is", shape.CalculateArea(circle))
	// fmt.Println("The Area of Square is", shape.CalculateArea(square))
	// cat := shape.Cat{Sound: "Mew"}
	// dog := shape.Dog{Sound: "Bark"}
	// cow := shape.Cow{Sound: "Moo"}
	// fmt.Println(shape.MakeSound(cat))
	// fmt.Println(shape.MakeSound(dog))
	// fmt.Println(shape.MakeSound(cow))
	payment := shape.NetBanking{Name: "Net Banking"}
	payment.Pay(1000)
}
