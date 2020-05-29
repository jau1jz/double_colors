package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	log.Printf("欢迎使用！")

	log.Printf("双色球一等奖号码计算系统。\n")

	input := bufio.NewScanner(os.Stdin)
	log.Printf("需要计算多少次（百万）:\n")
loop1:
	input.Scan()
	CalcCount, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("错误,请输入数字类型!!")
		goto loop1
	}
	log.Printf("需要开启多少个协程:\n")
	CalcCount = CalcCount * 1e6
loop2:
	input.Scan()
	GoCount, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("错误,请输入数字类型!!")
		goto loop2
	}
	wait := sync.WaitGroup{}
	begin := time.Now()
	numMap := make([]map[string]int, GoCount)
	rand.Seed(time.Now().UnixNano())
	waitCalc := sync.WaitGroup{}
	waitCalc.Add(GoCount)
	splitCount := CalcCount / GoCount
	for step := 0; step < GoCount; step++ {
		go func(index int) {
			numMap[index] = make(map[string]int)
			for k := 0; k < splitCount; k++ {
				nums := []int{}
				find := func(num int) bool {
					for _, v := range nums {
						if v == num {
							return true
						}
					}
					return false
				}

				for i := 0; i < 6; i++ {
					for {
						redNum := rand.Intn(33) + 1
						if find(redNum) {
							continue
						}
						nums = append(nums, redNum)
						break
					}
				}
				blueNum := rand.Intn(16)
				nums = append(nums, blueNum+1)
				sort.Ints(nums[:6])
				numMap[index][fmt.Sprintf("%v", nums)]++
			}
			waitCalc.Done()
		}(step)
	}

	waitCalc.Wait()
	pass := time.Now().Sub(begin)
	log.Printf("pass %ds ", pass/1e9)
	begin = time.Now()
	max := ""
	times := 0
	for _, v := range numMap {
		for k, v1 := range v {
			if v1 > times {
				times = v1
				max = k
			}
		}

	}
	pass = time.Now().Sub(begin)
	log.Printf("pass %ds key:%s , times :%d", pass/1e9, max, times)
	wait.Wait()

}
