package main

type Person struct {
	name   string
	height float64
	weight float64
	age    int
	sex    string
}

var InsertPeople Person

func inputRecord(name string, height float64, weight float64) {
	InsertPeople.name = name
	InsertPeople.height = height
	InsertPeople.weight = weight

}

func getRand() (height float64, weight float64) {
	return InsertPeople.height, InsertPeople.weight
}

func inputRecordSecond(age int, sex string) {
	InsertPeople.age = age
	InsertPeople.sex = sex
}
func getRandSecond() (age int, sex string) {
	return InsertPeople.age, InsertPeople.sex
}

func CalFatRate(age int, bmi float64, sex string) (FatRate float64) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
