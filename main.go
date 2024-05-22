package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 定义参数
	mode := flag.String("mode", "succeed", "模拟运行的结果")
	seconds := flag.Int("seconds", 10, "模拟运行时长")
	// 解析命令行参数
	flag.Parse()
	//获取的参数
	fmt.Println("mode==", *mode)
	fmt.Println("seconds==", *seconds)

	switch *mode {
	case "cpu":
		go cpuRun()
	case "memory":
		go memoryRun()
	}

	for i := 0; i < *seconds; i++ {
		fmt.Println("已经运行：", i, "秒")
		time.Sleep(time.Second)
	}
	switch *mode {
	case "succeed":
		fmt.Println("运行成功！")
	case "failed":
		panic("运行失败！")
	case "oom":
		memory := make([]int, 1e10)
		fmt.Println(memory)
	case "loop":
		fmt.Println("开始卡死！")
		for {

		}
	}
}

// CPU占用
func cpuRun() {
	var wg sync.WaitGroup
	maxGoroutines := runtime.NumCPU() - 1 // // 设置为CPU的核心数

	for i := 1; i <= maxGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("启动协程 #%d...\n", id)
			for {
				// 可以在这里添加一些轻量级操作，但注意即使是空循环也会消耗CPU资源
			}
		}(i) // 注意这里直接传入i的当前值，避免goroutine共享同一个变量的问题
	}

	// 等待所有goroutine启动完毕
	wg.Wait()
	fmt.Println("所有协程都已启动完毕！")
}

// 内存
func memoryRun() {
	// 开辟1GB的内存，这里使用byte类型，因为每个byte占1字节
	size := 1 << 30 // 1GB = 1 * 2^30 bytes

	// 创建一个大小为1GB的byte切片，并初始化为0
	slice := make([]byte, size)

	// 填充切片，将所有元素设置为1（在byte中，1表示其最大值255）
	for i := range slice {
		slice[i] = 255 // 设置为255等效于设置为1，因为在byte范围内，255代表最大的值，通常表示"全1"
	}
}
