package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	//"github.com/sjwhitworth/golearn/ensemble"
	//"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	//"github.com/sjwhitworth/golearn/trees"
	"math"
	"math/rand"
	"reflect"
)

func findBestSplit(partition base.FixedDataGrid) {
	var delta float64
	delta = math.MinInt64

	attrs := partition.AllAttributes()
	classAttrs := partition.AllClassAttributes()
	candidates := base.AttributeDifferenceReferences(attrs, classAttrs)

	fmt.Println(delta)
	fmt.Println(classAttrs)
	fmt.Println(reflect.TypeOf(partition))
	fmt.Println(reflect.TypeOf(candidates))

	for i, n := range attrs {
		fmt.Println(i)
		//fmt.Println(partition)
		fmt.Println(reflect.TypeOf(n))
		attributeSpec, _ := partition.GetAttribute(n)

		fmt.Println(partition.GetAttribute(n))
		_, rows := partition.Size()
		for j := 0; j < rows; j++ {
			data := partition.Get(attributeSpec, j)
			fmt.Println(base.UnpackBytesToFloat(data))
		}

	}
}

func findBestSplitsJOnXj(partion base.FixedDataGrid, attribute base.Attribute) {
	delta := 0.0
	k := 0.0
	vK := -999999.99 //TODO
	for i := 1; i < 99; {
		//TODO second while with i++
		i++

	}
	fmt.Println(delta)
	fmt.Println(k)
	fmt.Println(vK)
}

func main() {

	var tree base.Classifier

	rand.Seed(44111342)

	// Load in the iris dataset
	iris, err := base.ParseCSVToInstances("/home/kralli/go/src/github.com/sjwhitworth/golearn/examples/datasets/iris_headers.csv", true)
	if err != nil {
		panic(err)
	}

	// Discretise the iris dataset with Chi-Merge
	filt := filters.NewChiMergeFilter(iris, 0.999)
	for _, a := range base.NonClassFloatAttributes(iris) {
		filt.AddAttribute(a)
	}
	filt.Train()
	irisf := base.NewLazilyFilteredInstances(iris, filt)

	// Create a 60-40 training-test split
	//testData
	trainData, _ := base.InstancesTrainTestSplit(iris, 0.60)

	findBestSplit(trainData)

	//fmt.Println(trainData)
	//fmt.Println(testData)

	fmt.Println(tree)
	fmt.Println(irisf)
}
