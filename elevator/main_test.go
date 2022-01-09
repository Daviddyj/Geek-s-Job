package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	elevator := &elevator{}
	elevator.inputRecord(1, 5) //1为默认电梯楼层 5为楼层
	floorOfElevator := elevator.getElevatorFloor()
	if floorOfElevator != 1 {
		t.Fatalf("预期电梯1楼，但得到的是:%d,", floorOfElevator)
	}
}

func TestCase2(t *testing.T) {
	elevator := &elevator{}
	storey := &storey{}
	elevator.inputRecord(1, 5) //1为默认电梯楼层 5为楼层
	storey.press(3)
	elevator.goRun(storey, storey)
	floorOfElevator := elevator.getElevatorFloor()
	if floorOfElevator != 3 {
		t.Fatalf("预期电梯3楼，但得到的是:%d,", floorOfElevator)
	}

}

func TestCase3(t *testing.T) {
	elevator := &elevator{}
	elevator.inputRecord(3, 5) //3为默认电梯楼层 5为楼层
	storey4 := &storey{}
	storey2 := &storey{}

	storey4.press(4)
	//time.Sleep(time.Duration(1) * time.Second)      //电梯运行一层一秒
	storey2.press(2)
	elevator.goRun(storey4, storey4, storey2)

	floorOfElevator := elevator.getElevatorFloor()
	if floorOfElevator != 2 {
		t.Fatalf("预期电梯2楼，但得到的是:%d,", floorOfElevator)
	}

}

func TestCase4(t *testing.T) {
	elevator := &elevator{}
	elevator.inputRecord(3, 5) //3为默认电梯楼层 5为楼层
	storey2 := &storey{}
	storey4 := &storey{}
	storey5 := &storey{}

	storey4.press(4)
	storey5.press(5)
	storey2.press(2)
	//time.Sleep(time.Duration(1) * time.Second) //电梯运行一层一秒
	elevator.goRun(storey4, storey5, storey2)
	floorOfElevator := elevator.getElevatorFloor()
	if floorOfElevator != 2 {
		t.Fatalf("预期电梯2楼，但得到的是:%d,", floorOfElevator)
	}

	//time.Sleep(time.Duration(1) * time.Second) //电梯运行一层一秒
	//floorOfElevator = elevator.getElevatorFloor()
	//if floorOfElevator != 5 {
	//	t.Fatalf("预期电梯4楼，但得到的是:%d,", floorOfElevator)
	//}
	//
	//time.Sleep(time.Duration(3) * time.Second) //电梯运行一层一秒
	//floorOfElevator = elevator.getElevatorFloor()
	//if floorOfElevator != 4 {
	//	t.Fatalf("预期电梯4楼，但得到的是:%d,", floorOfElevator)
	//}

}
