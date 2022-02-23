package meta

//FileMeta：文件元信息结构
type FileMeta struct {
	FileSha1 string //文件唯一标识
	FileName string
	FileSize int64
	Location string //文件路径
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)

}

//UpdateFileMeta: 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta：通过 sha1 值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//GetLastFileMetas:获取批量的文件元信息列表
//func GetLastFileMetas(count int)[]FileMeta{
//	fMetaArray:=make([]FileMeta,len(fileMetas))
//	for _,v:=range fileMetas{
//		fMetaArray=append(fMetaArray,v)
//	}
//	sort.Sort(ByUploadTime(fMetaArray))
//	return fMetaArray[0:count]
//}
//RemoveFileMeta:删除元信息
func RemoveFileMeta(filesha1 string) {
	delete(fileMetas, filesha1)

}
