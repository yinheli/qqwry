package qqwry

import (
	"log"
	"testing"
)

func TestQQwry(t *testing.T) {
	q := NewQQwry("qqwry.dat")
	q.Find("183.224.52.133")
	log.Printf("ip:%v, country:%v, city:%v", q.Ip, q.Country, q.City)
}
