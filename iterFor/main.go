package iterFor

func Repeat(inputStr string, times int) (repeatStr string) { //注意这里repeatStr 就已经被初始化，所以函数内就直接赋值就可以了。
	repeatStr = ""
	for i := 0; i < times; i++ {
		repeatStr += inputStr
	}
	return repeatStr
}
