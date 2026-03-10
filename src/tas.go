package main

import (
	"fmt"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/arrays"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/conditionals"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/loops"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/maps"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/variables"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/takinginput"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/dataTypes"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/functions"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/pointers"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/structures"
	"git.druva.org/cloudapps/shareddrive-node/src/pkg/logger"
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

}
