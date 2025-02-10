package binary_test

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

/*
字节序
通常多字节对象 都被存储为 连续的字节序列
如 int x = 168496141 在内存中16进制存储为 0x0a0b0c0d

大端序 Big-Endian : 将数据低位放在较小的地址. 符合人类阅读习惯
x 排列为 0a0b0c0d

小端序 Little-Endian : 将数据地位放在较大的地址. 符合计算机读取方式. CPU读取时是从低向高读取的
x排列为 0d0c0b0a
*/

// unsafe.Sizeof: 查看int类型 在当前系统所占用内存大小. 类型为 uintptr
// go中 32位4字节 64位8字节
const INT_SIZE = int(unsafe.Sizeof(0))

func BinaryTest() {
	var x int = 168496141
	fmt.Printf("0x%08x\n", x) // 16进制显示为 0x0a0b0c0d

	i := 0x0a0b0c0d

	// 大端序
	BBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(BBytes, uint32(i))
	fmt.Println(BBytes) // [10 11 12 13]

	// 小端序
	LBytes := make([]byte, 4)                        // 创建长度为4的字节数组
	binary.LittleEndian.PutUint32(LBytes, uint32(i)) // 数字转为字节数组
	fmt.Println(LBytes)                              // [13 12 11 10]

	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	fmt.Println(bs) // &[13 12 11 10 0 0 0 0]

	if bs[0] == 0x10 {
		fmt.Println("大端序")
	} else {
		fmt.Println("小端序")
	}
}
