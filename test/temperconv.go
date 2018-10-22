package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64			// 摄氏温度
type Fahrenheit float64			// 华氏温度

func (self Celsius) String() string {
	return fmt.Sprint("%g°C", self)
}

func (self Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", self)
}

const (
	AbsoluteZeroC	Celsius	=	-273.15	// 绝对零度
	FreezingC		Celsius	=	0		// 结冰点温度
	BoilingC		Celsius	=	100		// 沸水温度
)



func CToF(c Celsius) Fahrenheit {
	// todo
	return Fahrenheit(c * 9 / 5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	// todo
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	tempC := Celsius(12)
	tempC += 1
	fmt.Println(tempC)
	fmt.Println("this is about Celsius...")

	tempF := Fahrenheit(12)
	tempF += 1
	fmt.Println(tempF)
	fmt.Println("this is about Fahrenheit...")

	fmt.Println(tempF == 0)
	fmt.Println(tempF >= 0)
	//fmt.Println(tempC == tempF) 这种情况是出现编译error，在进行基础操作时，如果操作数的类型是相同或则类型是继承关系，则可以操作。
	fmt.Println(tempC == Celsius(12))

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, FToC(f), c, CToF(c))
	}
}

