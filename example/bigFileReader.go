package example

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

func processChunk(buf []byte) {
	buffer := bytes.NewBuffer(buf)
	fmt.Println(buffer.String())
}

func bytesCombine(buffer bytes.Buffer, bytes ...[]byte) bytes.Buffer {
	for index := 0; index < len(bytes); index++ {
		buffer.Write(bytes[index])
	}
	return buffer
}

func ExecBigFileReader() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	// sync pools to reuse the memory and decrease the preassure on //Garbage Collector
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 3)
		return lines
	}}

	// 缓存起来的 byte
	// var buffer bytes.Buffer
	reader := bufio.NewReader(file)
	for {
		buf := linesPool.Get().([]byte) // the chunk size
		n, err := reader.Read(buf)      // loading chunk into buffer
		buf = buf[:n]
		if n == 0 {
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			return
		}

		// 读取整行，返回缓冲里的包含已读取的数据和 delim 字节的字节数组
		nextUntillNewline, err := reader.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}
		// 使用协程处理数据
		wg.Add(1)
		go func() {
			// process each chunk concurrently
			processChunk(buf)
			wg.Done()
		}()

		// 合并 buffer，(通常使用 bufio.Writer 即可保存到文件，无需自己实现)
		// buffer = bytesCombine(buffer, buf)
	}
	// fmt.Printf("%s", buffer.Bytes())

	wg.Wait()
}

/*
	读取一个 16GB 大小，上百万行内容的txt或者log文件，如果直接读取会导致机器负载过高该怎么办？
	思路：
	1. 一次性读取整个文件到内存中进行处理，这将显著的增加内存使用，但是会节省大量的时间。
	2. 逐行读取，这能帮助我们减少内存的使用，但会耗费大量的时间在IO上。
	3. 分段读取数据，既对占用内存可控，也能节省大量的时间。（推荐）

	优化：
	1. sync.Pool可以减轻GC的压力，我们可以重复使用已分配的内存，减少内存消耗，加快处理速度。
	2. goroutine帮助我们并发处理多个切块，显著加快处理速度。
*/
