package checkerr

import (
	"github.com/fatih/color"
	"path"
	"runtime"
	"strconv"
	"time"
)

var (
	ErrorsData map[string]string
)

func init() {
	ErrorsData = make(map[string]string)
}

func Log(appName string, err error) {
	if err != nil {
		pc, file, line, _ := runtime.Caller(1)
		timeTmp := " " + time.Now().Format("15:04:05.00000")
		functionTmp := " " + path.Ext(runtime.FuncForPC(pc).Name())[1:]
		codelineTmp := " " + path.Base(file) + ":" + strconv.Itoa(line)
		color.New(color.FgRed).PrintfFunc()(" [%s]", appName)
		color.New(color.FgMagenta).PrintfFunc()("%v", timeTmp)
		color.New(color.FgGreen).PrintfFunc()("%s%s", codelineTmp, functionTmp)
		color.Yellow(" %v", err)
		ErrorsData["["+appName+"]"+timeTmp+codelineTmp+functionTmp] = err.Error()
	}
	return
}

func Notice(noticeTXT, resultTXT string) {
	color.New(color.FgBlue).PrintfFunc()(" %v", noticeTXT)
	color.New(color.FgGreen).PrintfFunc()("%v\n", resultTXT)
	ErrorsData[noticeTXT] = resultTXT
}

func Error(noticeTXT, resultTXT string) {
	color.New(color.FgYellow).PrintfFunc()(" %v", noticeTXT)
	color.New(color.FgGreen).PrintfFunc()("%v\n", resultTXT)
	ErrorsData[noticeTXT] = resultTXT
}

func Warning(noticeTXT, resultTXT string) {
	color.New(color.FgRed).PrintfFunc()(" %v", noticeTXT)
	color.New(color.FgGreen).PrintfFunc()("%v\n", resultTXT)
	ErrorsData[noticeTXT] = resultTXT
}
