# Go Debounce

## Example

```go

func TestDebounce(t *testing.T) {
	count := 0
	callTimes := 0

	debouncedFunc := debounce.Debounce(15*time.Millisecond, func() {
		callTimes++
		fmt.Println("Debounced Function: " + strconv.Itoa(callTimes) + " => " + strconv.Itoa(count))
	})

	fmt.Println("Start：Call 300 times.")
	for i := 0; i < 300; i++ {
		count = i
		go debouncedFunc()
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	}
	fmt.Println("Ended：Call 300 times.")
	time.Sleep(3 * time.Second)
	fmt.Println("debouncedFunc called times：" + strconv.Itoa(callTimes))
}

```



