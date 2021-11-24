package mo

import (
	xconv "github.com/goclub/conv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math"
	"testing"
)

var testCoord = struct {
	WGS84 WGS84
	GCJ02 GCJ02
	BD09 BD09
}{
	WGS84: WGS84{121.47192098254719,31.22882564130273},
	GCJ02: GCJ02{121.47644966280748,31.22688746269209},
	BD09: BD09{121.48294392041718,31.23285726855796},
}

func TestGeojson(t *testing.T) {
	suite.Run(t, new(TestGeoJsonSuite))
}

type TestGeoJsonSuite struct {
	suite.Suite
}

func (suite TestGeoJsonSuite) TestNewPoint() {
	t := suite.T()
	point := NewPoint(testCoord.WGS84)
	assert.Equal(t,point.WGS84(), testCoord.WGS84)
	assert.Equal(t,point.Coordinates, []float64{testCoord.WGS84.Longitude, testCoord.WGS84.Latitude})
}
func (suite TestGeoJsonSuite) TestWGS84_BD09() {
	t := suite.T()
	if abs := math.Abs(testCoord.WGS84.BD09().Latitude- testCoord.BD09.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.WGS84.BD09().Longitude- testCoord.BD09.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}
func (suite TestGeoJsonSuite) TestWGS84_GCJ02() {
	t := suite.T()
	if abs := math.Abs(testCoord.WGS84.GCJ02().Latitude- testCoord.GCJ02.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.WGS84.GCJ02().Longitude- testCoord.GCJ02.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}
func (suite TestGeoJsonSuite) TestGCJ02_BD09() {
	t := suite.T()
	if abs := math.Abs(testCoord.GCJ02.BD09().Latitude- testCoord.BD09.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.GCJ02.BD09().Longitude- testCoord.BD09.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}
func (suite TestGeoJsonSuite) TestGCJ02_WGS84() {
	t := suite.T()
	if abs := math.Abs(testCoord.GCJ02.WGS84().Latitude- testCoord.WGS84.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.GCJ02.WGS84().Longitude- testCoord.WGS84.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}
func (suite TestGeoJsonSuite) TestGCJ02_LatCommaLngString() {
	t := suite.T()
	assert.Equal(t,testCoord.GCJ02.LatCommaLngString(), xconv.Float64String(testCoord.GCJ02.Latitude) + "," + xconv.Float64String(testCoord.GCJ02.Longitude))
}
func (suite TestGeoJsonSuite) TestBD09_GCJ02() {
	t := suite.T()
	if abs := math.Abs(testCoord.BD09.GCJ02().Latitude- testCoord.GCJ02.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.BD09.GCJ02().Longitude- testCoord.GCJ02.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}
func (suite TestGeoJsonSuite) TestBD09_WGS84() {
	t := suite.T()
	if abs := math.Abs(testCoord.BD09.WGS84().Latitude- testCoord.WGS84.Latitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
	if abs := math.Abs(testCoord.BD09.WGS84().Longitude- testCoord.WGS84.Longitude) ; abs > 0.001 {
		t.Errorf(t.Name() + " " + "fail abs is %#v", abs)
	}
}