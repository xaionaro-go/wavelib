package wavelib

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

// analog of https://github.com/rafat/wavelib/wiki/SWT-Example-Code
func ExampleSWT1D() {
	rand.Seed(0)
	obj := WaveInit("bior3.5")
	defer obj.Free()

	input := make([]float64, 256)
	for idx := range input {
		input[idx] = rand.Float64()
	}

	wt := WTInit(obj, "swt", 256, 1)
	defer wt.Free()

	SetWTConv(wt, "direct")
	SWT(wt, input)
	for _, value := range wt.Output() {
		fmt.Printf("%g\n", value)
	}

	output := make([]float64, 256)
	ISWT(wt, output)

	diff := make([]float64, 256)
	for i := 0; i < wt.SigLength(); i++ {
		diff[i] = output[i] - input[i]
	}

	fmt.Printf("\n MAX %g \n", absMax(diff))
	WTSummary(wt)
}

func absMax(s []float64) float64 {
	max := float64(0)
	for _, v := range s {
		abs := math.Abs(v)
		if abs > max {
			max = abs
		}
	}

	return max
}

func TestSWT1D(t *testing.T) {
	rand.Seed(0)
	obj := WaveInit("bior3.5")
	defer obj.Free()

	input := make([]float64, 256)
	for idx := range input {
		input[idx] = rand.Float64()
	}

	wt := WTInit(obj, "swt", 256, 1)
	defer wt.Free()

	SetWTConv(wt, "direct")
	SWT(wt, input)

	output := make([]float64, 256)
	ISWT(wt, output)

	diff := make([]float64, 256)
	for i := 0; i < wt.SigLength(); i++ {
		diff[i] = output[i] - input[i]
	}

	require.Less(t, absMax(diff), 1e-10)
	WTSummary(wt)
}
