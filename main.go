package main

import "fmt"

func main() {
	//声明
	key := ""
	loop := true

	balance := 10000.0

	money := 0.0

	note := ""

	details := "收支\t账户金额\t收支金额\t说 明"

	for {
		fmt.Println("\n------------家庭收支记账软件---------------")
		fmt.Println("             1 收支明细")
		fmt.Println("             2 登记收入")
		fmt.Println("             3 登记支出")
		fmt.Println("             4 退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("------------当前收支明细记录---------------")
				fmt.Println(details)
				fmt.Println("\t",balance,"\t",money,"\t",note)
			case "2":
				fmt.Println("------------当前登记收入记录---------------")
				fmt.Println("本次收入金额：")
				fmt.Scanln(&money)
				fmt.Println("本次收入说明：")
				fmt.Scanln(&note)
			case "3":
				fmt.Println("------------当前登记支出记录---------------")
			case "4":
				loop = false
			default:
				fmt.Println("请输入正确的选项..")
		}

		if !loop {
			break
		}
	}

	fmt.Println("你退出家庭记账软件的使用")
}
