## Go Check Error

- Collecting Error Logs
- Printing error logs
  - Print error log location


### init

```go
import "github.com/XRSec/checkerr"
```

### 0x01

```go
checkerr.Notice("打印通知", "结果")
checkerr.Notice("Print Notice", "Result")
```

### 0x02

```go
func main() {
	var err error
	if err := errors.New(Error Message); err != nil {
		checkerr.Log("main", err) // AppName main
	}
}
```

### Damo

```go
func main() {
	ErrorsData = make(map[string]string)
	for i := 0; i < 2; i++ {
		checkerr.Notice("Notice "+time.Now().Format("15:04:05.00000"), "Result "+strconv.Itoa(i))
		checkerr.Error("Error "+time.Now().Format("15:04:05.00000")+" ", "Result "+strconv.Itoa(i))
		checkerr.Warning("Warning "+time.Now().Format("15:04:05.00000")+" ", "Result "+strconv.Itoa(i))
		fmt.Println("----------------------")
		if i == 0 {
			if err := errors.New(strconv.Itoa(i)); err != nil {
				checkerr.Log("main", err)
				fmt.Println("The ErrorsData = ", ErrorsData)
			}
		} else {
			//ErrorsData = make(map[string]string)
			if err := errors.New(strconv.Itoa(i)); err != nil {
				checkerr.Log("main", err)
				fmt.Println("The ErrorsData = ", ErrorsData)
			}
		}
	}
	fmt.Println("----------------------")
	for i, v := range ErrorsData {
		checkerr.Warning(i, v)
	}
}
```