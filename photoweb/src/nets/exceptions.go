package nets

//用于统一捕获50x服务端错误
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
