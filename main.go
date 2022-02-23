package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"
)

var red = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33,
}

var blue = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16,
}

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
	CalcCount *= 1e6
	wait := sync.WaitGroup{}
	pNumber := runtime.GOMAXPROCS(0)
	wait.Add(pNumber)
	numMap := map[string]int{}
	begin := time.Now()
	pool := sync.Pool{}
	pool.New = func() interface{} {
		return append(red)
	}
	localMap := make([]map[string]int, pNumber)
	for i := 0; i < pNumber; i++ {
		go func(times int, index int) {
			localMap[index] = make(map[string]int)
			for i := 0; i < times; i++ {
				//random red ball
				get := pool.Get().([]int)
				rand.Shuffle(len(get), func(i, j int) {
					get[i], get[j] = get[j], get[i]
				})
				sort.Ints(get[:6])
				//random blue ball
				get[6] = blue[rand.Intn(len(blue)-1)]
				result := fmt.Sprintf("%d,%d,%d,%d,%d,%d %d", get[0], get[1], get[2], get[3], get[4], get[5], get[6])
				localMap[index][result] += 1
				pool.Put(get)
			}
			wait.Done()
		}(CalcCount/pNumber, i)
	}
	wait.Wait()
	for _, local := range localMap {
		for key, value := range local {
			numMap[key] += value
		}
	}
	max := ""
	times := 0
	for key, value := range numMap {
		if value > times {
			max = key
			times = value
		}
	}
	pass := time.Now().Sub(begin)
	log.Printf("num %s times :%d\n", max, times)
	log.Printf("pass %ds\n", pass/1e9)
}
