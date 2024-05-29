package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func SysInfo() {
	// 获取网络IO信息
	netIO, err := net.IOCounters(false)
	if err != nil {
		fmt.Println("无法获取网络IO信息:", err)
	} else {
		fmt.Println("网络IO信息:")
		for _, io := range netIO {
			fmt.Printf("  接口: %s\n", io.Name)
			fmt.Printf("    接收数据: %d bytes\n", io.BytesRecv)
			fmt.Printf("    发送数据: %d bytes\n", io.BytesSent)
		}
	}

	// 获取CPU信息
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("无法获取CPU信息:", err)
	} else {
		fmt.Println("CPU信息:")
		for i, percent := range cpuPercent {
			fmt.Printf("  核心 %d: %.2f%% 使用率\n", i, percent)
		}
	}

	// 获取内存信息
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("无法获取内存信息:", err)
	} else {
		fmt.Println("内存信息:")
		fmt.Printf("  总内存: %d bytes\n", memInfo.Total)
		fmt.Printf("  可用内存: %d bytes\n", memInfo.Available)
		fmt.Printf("  使用率: %.2f%%\n", memInfo.UsedPercent)
	}
}
