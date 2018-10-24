package main

import (
	"fmt"
	"jvmgo/classpath"
	"strings"
	"jvmgo/classfile"
)


func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class=="" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}


func startJVM(cmd *Cmd) {
	// 设置jre(java标准) 和 cp(用户路径) 的路径
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)

	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}


func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\b", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlag())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfacecs: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))

	for _, f := range cf.Fields() {
		fmt.Printf("   %s\n", f.Name())
	}
	fmt.Printf("method count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("    %s\n", m.Name())
	}
}