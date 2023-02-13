package main
import (
  "fmt"
)
func myFunc(){
  fmt.Println("Call myFunc")
}
func myFunc2(){
  fmt.Println("Call myFunc2")
}
func main() {
  fmt.Println("hi, 自定义函数必须被main函数调用，才能运行。")
  //自定义函数必须被main函数调用，才能运行
  myFunc2()
}
