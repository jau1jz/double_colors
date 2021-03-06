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

	log.Printf("你需要计算几组双色球号码:\n")

	input := bufio.NewScanner(os.Stdin)
loop:
	input.Scan()
	NumCount, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("错误,请输入数字类型!!")
		goto loop
	}
	log.Printf("需要计算多少次（百万）:\n")
loop2:
	input.Scan()
	CalcCount, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("错误,请输入数字类型!!")
		goto loop2
	}
	wait := sync.WaitGroup{}
	wait.Add(NumCount)
	for i := 0; i < NumCount; i++ {
		log.Printf("开始计算第%d组。", i+1)
		go func() {
			begin := time.Now()
			numMap := map[string]int{}
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < CalcCount*1e6; i++ {
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
				numMap[fmt.Sprintf("%v", nums)]++
			}
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
			wait.Done()
		}()
	}
	wait.Wait()
}
