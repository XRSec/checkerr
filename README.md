## Demo

### Eazy

```go
```



```go
func main() {
	ErrorsData = make(map[string]string)
	for i := 0; i < 2; i++ {
		Notice("Notice "+time.Now().Format("15:04:05.00000")+" ", "Result "+strconv.Itoa(i))
		Error("Error "+time.Now().Format("15:04:05.00000")+" ", "Result "+strconv.Itoa(i))
		Warning("Warning "+time.Now().Format("15:04:05.00000")+" ", "Result "+strconv.Itoa(i))
		if i == 0 {
			if err := errors.New(strconv.Itoa(i)); err != nil {
				CheckErr("main", err)
				fmt.Println("----------------------")
				fmt.Println("The ErrorsData = ", ErrorsData)
			}
		} else {
			//ErrorsData = make(map[string]string)
			fmt.Println("----------------------")
			if err := errors.New(strconv.Itoa(i)); err != nil {
				CheckErr("main", err)
				fmt.Println("----------------------")
				fmt.Println("The ErrorsData = ", ErrorsData)
			}
		}
	}
	fmt.Println("----------------------")
	for i, v := range ErrorsData {
		fmt.Println(i, v)
	}
}
```