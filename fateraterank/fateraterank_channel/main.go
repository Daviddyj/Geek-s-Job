package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type People struct {
	name    string
	fatRate float64
	rank    int
}

func main() {
	for {
		peopleCh := make(chan People, 1000)
		rankCh := make(chan []People, 1000)
		peopleSlice := []People{}
		peopleCounter := 1000
		wg := sync.WaitGroup{}
		wg.Add(peopleCounter)
		rand.Seed(time.Now().Unix())

		for i := 0; i < peopleCounter; i++ {
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				var people = People{
					name:    fmt.Sprintf("Stu%d", i),
					fatRate: randFloats(0.1, 0.4, 1),
				}
				//将各个goroutine中传入的结构体放入peopleCh
				peopleCh <- people
				//获取排名
				getRand(people.name, rankCh)
			}(i, &wg)
		}

		//取出peopleCh中的结构体放入peopleSlice切片
		finishedFileCount := 0
		for people := range peopleCh {
			finishedFileCount++
			peopleSlice = append(peopleSlice, people)
			if finishedFileCount == peopleCounter {
				close(peopleCh)
			}
		}

		//将对切片peopleSlice中的结构体根据fatRate大小进行升序排列
		sort.Slice(peopleSlice, func(i, j int) bool {
			return peopleSlice[i].fatRate < peopleSlice[j].fatRate
		})
		//修改切片中People结构rank体字段的值
		for i, _ := range peopleSlice {
			peopleSlice[i].rank = i + 1
		}
		//将排好序的peopleSlice传入rankCh通道，根据用户数量传入相同数量的peopleSlice，确保每个用户都能拿到完整的一份排好顺序的peopleSlice
		for i := 0; i < peopleCounter; i++ {
			rankCh <- peopleSlice
		}

		wg.Wait()
	}
}

//从rankCh中拿出peopleSlice切片，根据用户的名字取出对应的排名
func getRand(name string, rank chan []People) {
	peopleSlice := <-rank
	for _, people := range peopleSlice {
		if people.name == name {
			fmt.Printf("%s的体脂排名为%d\n", name, people.rank)
		}
	}
}

//随机生成体脂率
func randFloats(min, max float64, n int) float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	result := res[0]
	return result
}
