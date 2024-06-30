package main

import (
	"fmt"
	geoip2db "github.com/cc14514/go-geoip2-db"

	"net"
)

func main() {
	db, _ := geoip2db.NewGeoipDbByStatik()
	defer db.Close()
	record, _ := db.City(net.ParseIP("113.246.158.106"))

	province := record.Subdivisions[0].Names["zh-CN"]
	city := record.City.Names["zh-CN"]
	fmt.Println(province, city)

}
