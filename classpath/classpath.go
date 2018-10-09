package classpath

import (
	"path/filepath"
	"os"
)


type Classpath struct {
	bootClasspath 	Entry
	extClasspath	Entry
	userClasspath	Entry
}


func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}


/**
找到jre的路径
1、输入的路径是否正确
2、如果1出现错误，则使用java_home的路径里的jre
 */
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can't find jre folder!")
}


func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


/**
分别加载
jre/lib里的文件
jre/lib/ext里的文件
 */
func(self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}


/**
加载userClassPath路径文件
 */
func(self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newWildcardEntry(cpOption)
}


/**
根据class的完整命名，加载class的二进制数据
from: 启动路径、扩展路径、用户路径
 */
func(self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	return self.userClasspath.readClass(className)
}

func(self *Classpath) String() string {
	return self.userClasspath.String()
}