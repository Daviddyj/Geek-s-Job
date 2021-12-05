package main

import "fmt"

func main() {
	fmt.Println("斜率计算器")
	for Whethercontinue := true; Whethercontinue; {
		var x1, x2, y1, y2, x3, x4, y3, y4 float64
		fmt.Printf("x1:")
		fmt.Scanln(&x1)
		fmt.Printf("y1:")
		fmt.Scanln(&y1)
		fmt.Printf("x2:")
		fmt.Scanln(&x2)

		fmt.Printf("y2:")
		fmt.Scanln(&y2)

		fmt.Printf("x3:")
		fmt.Scanln(&x3)
		fmt.Printf("y3:")
		fmt.Scanln(&y3)
		fmt.Printf("x4:")
		fmt.Scanln(&x4)
		fmt.Printf("y4:")
		fmt.Scanln(&y4)

		if (x2-x1)*(y4-y3) == (y2-y1)*(x4-x3) {
			fmt.Println("两直线平行")
		} else {
			fmt.Println("两直线不平行")
		}

		var continueFlag string
		fmt.Print("是否要继续，继续请输入y,退出请输入n:")
		fmt.Scanln(&continueFlag)
		if continueFlag == "y" {
			Whethercontinue = true
		} else {
			Whethercontinue = false
		}

	}

}
