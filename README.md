[![GoDoc](https://godoc.org/github.com/xaionaro-go/wavelib?status.svg)](https://pkg.go.dev/github.com/xaionaro-go/wavelib?tab=doc)
# Example
```go
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
```