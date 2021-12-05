package main

import "fmt"

func main() {
	fmt.Println("基于BMI的体脂计算器")
	var SumBMI float64
	SumBMI = 0
	var BMI float64
	var times int
	var averageSumBMI float64
	whetherContinue := true
	for i := 0; whetherContinue; i++ {

		var name string
		var weight float64
		var height float64
		var age int
		var sex string

		fmt.Printf("姓名:")
		fmt.Scanln(&name)
		fmt.Printf("身高:")
		fmt.Scanln(&height)
		fmt.Printf("体重:")
		fmt.Scanln(&weight)
		fmt.Printf("性别:")
		fmt.Scanln(&sex)
		fmt.Printf("年龄:")
		fmt.Scanln(&age)

		BMI = weight / (height * height)

		var sexvalue int

		if sex == "男" {
			sexvalue = 1
		} else {
			sexvalue = 0
		}

		var fatrate float64
		fatrate = (1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexvalue)) / 100
		fmt.Printf("我的体脂为%f\n", BMI)
		fmt.Printf("我的体脂率为%f\n", fatrate)

		if sex == "男" {
			if age >= 18 && age <= 39 {
				if fatrate >= 0 && fatrate <= 0.1 {
					fmt.Println("偏瘦")
				} else if fatrate >= 0.1 && fatrate <= 0.106 {
					fmt.Println("标准")
				} else if fatrate >= 0.17 && fatrate <= 0.21 {
					fmt.Println("偏重")
				} else if fatrate >= 0.22 && fatrate <= 0.26 {
					fmt.Println("肥胖")
				} else if fatrate >= 0.27 {
					fmt.Println("严重肥胖")
				}
			} else if age >= 40 && age <= 60 {
				fmt.Println("我还没到中年")
			} else if age > 60 {
				fmt.Println("你太老了")
			} else {
				fmt.Println("你太年轻，好好享受生活，不要太在意体脂率")
			}
		} else {
			fmt.Println("女")
		}

		var continueFlag string
		fmt.Print("是否继续？y表示继续，n表示不继续:")
		fmt.Scanln(&continueFlag)
		if continueFlag == "y" {
			whetherContinue = true
		} else {
			whetherContinue = false
		}

		fmt.Printf("循环打印了%d次\n", i+1)
		times = i + 1
		SumBMI += BMI

	}
	averageSumBMI = SumBMI / float64(times)
	fmt.Printf("所有人的体脂总和为%f", SumBMI)
	fmt.Printf("所有人的平均体脂为%f", averageSumBMI)
}
