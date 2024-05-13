package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"math"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "CalculateApp",
		Usage: "a simple CLI calculate application",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"}, //命令缩写
				Usage:   "加法",
				Action: func(c *cli.Context) error {
					//	加法要获得两个参数
					addNum1 := c.Args().First()
					addNum2 := c.Args().Get(1)

					firstAddF, err := strconv.ParseFloat(addNum1, 64)
					if err != nil {
						firstAdd, err := strconv.Atoi(addNum1)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
						} else {
							firstAddF = float64(firstAdd)

						}

					}
					secondAddf, err := strconv.ParseFloat(addNum2, 64)
					if err != nil {
						secondAdd, err := strconv.Atoi(addNum2)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
							return nil
						}
						secondAddf = float64(secondAdd)

					}
					fmt.Printf("%f + %f = %f", firstAddF, secondAddf, firstAddF+secondAddf)
					return nil
				},
			},
			{
				Name:    "subtract",
				Aliases: []string{"s"}, //命令缩写
				Usage:   "减法",
				Action: func(c *cli.Context) error {
					//	加法要获得两个参数
					subtractNum1 := c.Args().First()
					subtractNum2 := c.Args().Get(1)

					firstSF, err := strconv.ParseFloat(subtractNum1, 64)
					if err != nil {
						firstS, err := strconv.Atoi(subtractNum1)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
						} else {
							firstSF = float64(firstS)

						}

					}
					secondSF, err := strconv.ParseFloat(subtractNum2, 64)
					if err != nil {
						secondS, err := strconv.Atoi(subtractNum2)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
							return nil
						}
						secondSF = float64(secondS)
					}
					fmt.Printf("%f - %f = %f", firstSF, secondSF, firstSF-secondSF)
					return nil
				},
			},
			{
				Name:    "divide",
				Aliases: []string{"d"}, //命令缩写
				Usage:   "除法",
				Action: func(c *cli.Context) error {
					//	加法要获得两个参数
					Num1 := c.Args().First()
					Num2 := c.Args().Get(1)

					firstDF, err := strconv.ParseFloat(Num1, 64)
					if err != nil {
						firstD, err := strconv.Atoi(Num1)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
						} else {
							firstDF = float64(firstD)

						}

					}
					secondDF, err := strconv.ParseFloat(Num2, 64)
					if err != nil {
						secondD, err := strconv.Atoi(Num2)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
							return nil
						}
						secondDF = float64(secondD)
					}
					if math.Abs(secondDF) < 1e-12 {
						fmt.Println("除数不能等于0")
					}
					fmt.Printf("%f / %f = %f", firstDF, secondDF, firstDF/secondDF)
					return nil
				},
			},
			{
				Name:    "multiply",
				Aliases: []string{"m"}, //命令缩写
				Usage:   "乘法",
				Action: func(c *cli.Context) error {
					//	加法要获得两个参数
					Num1 := c.Args().First()
					Num2 := c.Args().Get(1)

					firstMF, err := strconv.ParseFloat(Num1, 64)
					if err != nil {
						firstM, err := strconv.Atoi(Num1)
						if err != nil {
							fmt.Println("加数格式错误，请输入数字！")
						} else {
							firstMF = float64(firstM)

						}

					}
					secondMF, err := strconv.ParseFloat(Num2, 64)
					if err != nil {
						secondM, err := strconv.Atoi(Num2)
						if err != nil {
							fmt.Println("参数格式错误，请输入数字！")
							return nil
						}
						secondMF = float64(secondM)
					}
					fmt.Printf("%f * %f = %f", firstMF, secondMF, firstMF*secondMF)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
