package rtda

/**
	这里的槽适用于描述： 局部标亮表 和 操作栈
 */

type Slot struct {
	num			int32
	ref 		*Object
}
