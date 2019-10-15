package main

import "fmt"

//DivedeError 定义一个DivideError结构
type DivedeError struct {
	divedee int
	divider int
}

//实现"error"接口
func (de *DivedeError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
`
	return fmt.Sprintf(strFormat, de.divedee)
}

//Divide 除法
func Divide(varDividee int, varDivider int) (result int, errMsg string) {
	if varDivider == 0 {
		dData := DivedeError{
			divedee: varDividee,
			divider: varDivider,
		}
		errMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func main() {
	//正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10=", result)
	}

	//当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is:", errorMsg)
	}
}
