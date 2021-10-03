package netdisk

type Folder interface {
	GetName() string                     // 获取文件夹名称
	GetParent() Folder                   // 获取上一级文件夹
	GetFileList() []File                 // 获取文件列表
	GetFolderList() []Folder             // 获取文件夹列表
	DeleteFolder(Folder) error           // 删除文件夹
	DeleteFile(Folder) error             // 删除文件
	CreateFolder(string) (Folder, error) // 创建文件夹
	CreateFile(string) (File, error)     // 创建文件
	OpenFile(string) File                // 根据文件名称获取文件
	OpenFolder(string) Folder            // 根据文件夹名获取文件夹
}
