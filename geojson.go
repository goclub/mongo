package mo

import (
	xconv "github.com/goclub/conv"
	"github.com/goclub/mongo/internal/coord"
)
// geojson https://zhuanlan.zhihu.com/p/141554586
// NewPoint(mo.WGS84{121.48294,31.2328}) // WGS84{经度,纬度}
type Point struct {
	Type        pointType    `json:"type" bson:"type"`
	// []float64{longitude, latitude} []float64{经度, 纬度}
	// 可能所有人都至少一次踩过这个坑：地理坐标点用字符串形式表示时是纬度在前，经度在后（ "latitude,longitude" ），
	// 而数组形式表示时是经度在前，纬度在后（ [longitude,latitude] ）—顺序刚好相反。
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`

}
// 用内部类型来强制调用者使用 NewPoint() 来创造 Point
type pointType *string

func NewPoint(data WGS84) Point {
	typevalue := "Point"
	return Point{
		&typevalue,
		[]float64{data.Longitude, data.Latitude},
	}
}

func (p Point) WGS84() WGS84 {
	return WGS84{
		Longitude: p.Coordinates[0],
		Latitude: p.Coordinates[1],
	}
}

type WGS84 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude float64 `json:"latitude" note:"纬度"`
}
func (data WGS84) GCJ02() GCJ02 {
	lng, lat := coordtransform.WGS84toGCJ02(data.Longitude, data.Latitude)
	return GCJ02{
		Longitude: lng,
		Latitude: lat,
	}
}
func (data WGS84) BD09() BD09 {
	lng, lat := coordtransform.WGS84toBD09(data.Longitude, data.Latitude)
	return BD09{
		Longitude: lng,
		Latitude: lat,
	}
}
type GCJ02 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude float64 `json:"latitude" note:"纬度"`
}
// 返回 "纬度,经度" 格式字符串
// 可能所有人都至少一次踩过这个坑：地理坐标点用字符串形式表示时是纬度在前，经度在后（ "latitude,longitude" ），
// 而数组形式表示时是经度在前，纬度在后（ [longitude,latitude] ）—顺序刚好相反。
func (data GCJ02) LatCommaLngString() (latCommaLng string) {
	return xconv.Float64String(data.Latitude) + "," + xconv.Float64String(data.Longitude)
}
func (data GCJ02) WGS84() WGS84 {
	lng, lat := coordtransform.GCJ02toWGS84(data.Longitude, data.Latitude)
	return WGS84{
		Longitude: lng,
		Latitude: lat,
	}
}
func (data GCJ02) BD09() BD09 {
	lng, lat := coordtransform.GCJ02toBD09(data.Longitude, data.Latitude)
	return BD09{
		Longitude: lng,
		Latitude: lat,
	}
}

type BD09 struct {
	Longitude float64 `json:"longitude" note:"经度"`
	Latitude float64 `json:"latitude" note:"纬度"`
}
func (data BD09) WGS84() WGS84 {
	lng, lat := coordtransform.BD09toWGS84(data.Longitude, data.Latitude)
	return WGS84{
		Longitude: lng,
		Latitude: lat,
	}
}
func (data BD09) GCJ02() GCJ02 {
	lng, lat := coordtransform.BD09toGCJ02(data.Longitude, data.Latitude)
	return GCJ02{
		Longitude: lng,
		Latitude: lat,
	}
}