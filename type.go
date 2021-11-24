package mo

import "github.com/goclub/mongo/internal/coord"

type Point struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

func NewPoint(data WGS84) Point {
	return Point{
		"Point",
		[]float64{data.Longitude, data.Latitude},
	}
}

func (p Point) BD09() BD09 {
	return BD09{p.Coordinates[0], p.Coordinates[1]}
}

type WGS84 struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
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
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
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
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
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