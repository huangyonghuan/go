package main
import (
	"awesomeProject"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//点用其他文件或项目
	fmt.Println("%v\n", awesomeProject.CeShi())
	fmt.Println("%v\n", awesomeProject.Dd())
	//输入输出
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println("你输入的是：", input.Text())

}