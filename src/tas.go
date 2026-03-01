package main

import (
	"fmt"

	"git.druva.org/cloudapps/shareddrive-node/src/code/arrays"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/conditionals"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/loops"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/maps"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/structures"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/variables"

	// "git.druva.org/cloudapps/shareddrive-node/src/code/takinginput"
	// "git.druva.org/cloudapps/shareddrive-node/src/code/dataTypes"
	"git.druva.org/cloudapps/shareddrive-node/src/pkg/logger"
	"git.druve.org/cloudapps/shareddrive-node/src/code/functions"
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
	arrays.SlicesGrowth()
	// logger.Info("msg=The Operation has completed")
	summation, different = functions.TupleSum()
}
