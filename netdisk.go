package netdisk

type NetDisk interface {
	GetFolder() Folder    // 获取根目录
	GetFreeBytes() int64  // 获取未用字节数
	GetUsedBytes() int64  // 获取已用字节数
	GetTotalBytes() int64 // 获取总共可用字节数
}
