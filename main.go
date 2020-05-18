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
loop2:
	input.Scan()
	CalcCount, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("错误,请输入数字类型!!")
		goto loop2
	}
	CalcCount = CalcCount * 1e6
	wait := sync.WaitGroup{}

	begin := time.Now()
	numMap := map[string]int{}
	rand.Seed(time.Now().UnixNano())
	waitCalc := sync.WaitGroup{}
	waitCalc.Add(4)
	splitCount := CalcCount / 4
	locker := sync.Mutex{}
	for step := 0; step < 4; step++ {
		go func() {
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
				locker.Lock()
				numMap[fmt.Sprintf("%v", nums)]++
				locker.Unlock()
			}
			waitCalc.Done()
		}()
	}

	waitCalc.Wait()
	max := ""
	times := 0
	for k, v := range numMap {
		if v > times {
			times = v
			max = k
		}
	}
	pass := time.Now().Sub(begin)
	log.Printf("pass %ds key:%s , times :%d", pass/1e9, max, times)
	wait.Wait()

}
