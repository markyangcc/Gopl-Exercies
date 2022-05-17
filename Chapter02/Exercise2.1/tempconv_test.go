// run test with "go test -v"
package tempconv

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	const (
		c Celsius    = 100
		f Fahrenheit = 212
		k Kelvin     = 373.15
	)

	const accuracy = 0.0000001

	if math.Abs(float64(CToF(c)-f)) > accuracy {t.Fatalf("CToF() error")}
	if math.Abs(float64(CToK(c)-k)) > accuracy {t.Fatalf("CToK() error")}
	if math.Abs(float64(FToC(f)-c)) > accuracy {t.Fatalf("FToC() error")}
	if math.Abs(float64(FToK(f)-k)) > accuracy {t.Fatalf("FToK() error")}
	if math.Abs(float64(KToC(k)-c)) > accuracy {t.Fatalf("KToC() error")}
	if math.Abs(float64(KToF(k)-f)) > accuracy {t.Fatalf("KToF() error")}
}
