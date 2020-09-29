package util

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"strings"
)

// 把字符串转为point结构体
func GeoPointFromString(location string) (point *elastic.GeoPoint, err error){
	locationSlice := strings.SplitN(location, ",", 2)
	if len(locationSlice) != 2{
		return nil, fmt.Errorf("%s is not a valid geo point string", location)
	}
	lon,lat := locationSlice[0], locationSlice[1]
	if point, err = elastic.GeoPointFromString(lat+","+lon); err != nil {
		return
	}
	return
}