package ph

import (
	"github.com/fogleman/gg"
	"testing" //nolint:gci
)

func Testph(T *testing.T) {
	context := gg.NewContext(200, 100)
	context.DrawRectangle(0, 0, float64(200), float64(100))
}
