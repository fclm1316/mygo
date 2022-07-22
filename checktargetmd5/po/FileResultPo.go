package po

// type FileResult struct {
// 	Id            int64 `json:"id"`
// 	FilResultInfo FilResultInfo
// }

type FilResultInfo struct {
	Filename string `json:"filename"`
	Filemd5  string `json:"filemd5"`
	Fileip   string `json:"fileip"`
}
