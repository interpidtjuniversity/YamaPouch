package network

import (
	"fmt"
	"net"
	"testing"
)

func TestAllocate(t *testing.T) {
	_, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	ip, _ := ipAllocator.Allocate(ipnet)
	t.Logf("alloc ip: %v", ip)
}

func TestAllocateAnother(t *testing.T) {
	_, ipnet, _ := net.ParseCIDR("192.168.10.1/24")
	ip, _ := ipAllocator.Allocate(ipnet)
	t.Logf("alloc ip: %v", ip)
}

func TestRelease(t *testing.T) {
	ip, ipnet, _ := net.ParseCIDR("192.168.10.4/24")
	ipAllocator.Release(ipnet, &ip)
}

func TestIPAM_Release(t *testing.T) {
	o, ipnet, _ := net.ParseCIDR("192.168.30.1/24")
	fmt.Print(o)
	fmt.Print("\n")
	fmt.Print(ipnet)
	fmt.Print("\n")
	ip := net.ParseIP("192.168.30.1")
	fmt.Print(ip)
	ipAllocator.Release(ipnet, &ip)
}