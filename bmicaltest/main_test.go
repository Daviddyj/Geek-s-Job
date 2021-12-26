package main

import "testing"

func TestCase1(t *testing.T) {
	inputRecord("小明", 1.75, 75)
	height, weight := getRand()
	if height <= 0 {
		t.Fatalf("预期小明的身高为正数，但是得到的是:%f", height)
	}
	if weight <= 0 {
		t.Fatalf("预期小明的体重为正数，但是得到的是:%f", weight)
	}
	bmi := weight / (height * height)
	if bmi <= 0 {
		t.Fatalf("预期小明的bmi为正数，但是得到的是:%f", bmi)
	}

	inputRecordSecond(24, "男")
	age, sex := getRandSecond()
	if age <= 0 || age > 150 {
		t.Fatalf("预期小明的年龄为正数且小于150岁，但是得到的是:%d", age)
	}
	if sex != "男" && sex != "女" {
		t.Fatalf("预期小明的性别为男或者女，但是得到的是:%s", sex)
	}
	FatRate := CalFatRate(age, bmi, sex)
	if FatRate < 0.1 || FatRate > 0.26 {
		t.Fatalf("与预期的体脂率区别不符，得到的是:%s", sex)

	}

}
