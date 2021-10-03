package netdisk

type File interface {
	GetName() string                      // 获取文件名
	GetSize() int64                       // 获取文件大小
	GetCtime() int                        // 获取文件创建时间
	Seek(int64, int) (int64, error)       // 设置文件读写偏移
	Read([]byte) (int64, error)           // 按照偏移读取文件
	ReadAt([]byte, int64) (int64, error)  // 偏移读取文件
	Write([]byte) (int64, error)          // 写入文件
	WriteAt([]byte, int64) (int64, error) // 偏移写入文件
}
