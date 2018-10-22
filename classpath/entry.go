package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

/**
该接口的实现主要有4种：
	Dir 和 Zip 为具体的文件，差异点为 readClass 方法
	Composite 和 wild 为组合对象，差异点为 newXXX 方法
 */
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			return newZipEntry(path)
	}

	return newDirEntry(path)
}
