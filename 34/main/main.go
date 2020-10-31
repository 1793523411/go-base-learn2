package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	net2 "github.com/shirou/gopsutil/net"
)

//! psutil是一个跨平台进程和系统监控的Python库，而gopsutil是其Go语言版本的实现
// !Go语言部署简单、性能好的特点非常适合做一些诸如采集系统信息和监控的服务

//!采集CPU相关信息。
// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

// mem info
func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

// host info
func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// disk info
func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

// netIO
func getNetInfo() {
	info, _ := net2.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	fmt.Println(addrs)
	// [127.0.0.1/8 192.168.191.128/24 ::1/128 fe80::f0c8:6579:c4c:e77e/64]
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		fmt.Println(ipAddr)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// // Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	return localAddr.IP.String()
}

func main() {
	// getCpuInfo()
	// getCpuLoad()
	// getMemInfo()
	// getHostInfo()
	// fmt.Println("----------------------")
	// getDiskInfo()

	// getNetInfo()
	res, err := GetLocalIP()
	if err == nil {
		fmt.Println(res)
	}
	// res := GetOutboundIP()
	// fmt.Println(res)
}
