package main

import "fmt"

func main() {
	fmt.Println("斜率计算器")
	for Whethercontinue := true; Whethercontinue; {
		var x1, x2, y1, y2, x3, x4, y3, y4 float64
		x1, y1 = getPointXYFromInput()
		x2, y2 = getPointXYFromInput()
		x3, y3 = getPointXYFromInput()
		x4, y4 = getPointXYFromInput()
		//fmt.Printf("x1:")
		//fmt.Scanln(&x1)
		//fmt.Printf("y1:")
		//fmt.Scanln(&y1)
		//fmt.Printf("x2:")
		//fmt.Scanln(&x2)
		//
		//fmt.Printf("y2:")
		//fmt.Scanln(&y2)
		//
		//fmt.Printf("x3:")
		//fmt.Scanln(&x3)
		//fmt.Printf("y3:")
		//fmt.Scanln(&y33
		//fmt.Scanln(&x4)
		//fmt.Printf("y4:")
		//fmt.Scanln(&y4)

		calulate(x2, x1, y4, y3, y2, y1, x4, x3)

		panduan(Whethercontinue)

	}

}

func panduan(Whethercontinue bool) {
	var continueFlag string
	fmt.Print("是否要继续，继续请输入y,退出请输入n:")
	fmt.Scanln(&continueFlag)
	if continueFlag == "y" {
		Whethercontinue = true
	} else {
		Whethercontinue = false
	}
}

func calulate(x2 float64, x1 float64, y4 float64, y3 float64, y2 float64, y1 float64, x4 float64, x3 float64) {
	if (x2-x1)*(y4-y3) == (y2-y1)*(x4-x3) {
		fmt.Println("两直线平行")
	} else {
		fmt.Println("两直线不平行")
	}
}

func getPointXYFromInput() (float64, float64) {
	var x, y float64
	fmt.Printf("x:")
	fmt.Scanln(&x)
	fmt.Printf("y:")
	fmt.Scanln(&y)
	return x, y
}
