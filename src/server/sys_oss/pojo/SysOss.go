package pojo

type UploadReq struct {
	DirName string `json:"dirName" form:"dirName" binding:"required"` //oss文件夹名称 如：image file video photo
	Postfix string `json:"postfix" form:"postfix" binding:"required"` //文件后缀 如：.jpg .png .docx .xlsx
}

type GetSysOssListReq struct {
	DirName string `json:"dirName" form:"dirName" binding:"required"` //oss文件夹名称 如：image file video photo
}
