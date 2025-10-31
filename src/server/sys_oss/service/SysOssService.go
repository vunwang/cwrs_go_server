package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_oss/pojo"
	"github.com/gin-gonic/gin"
	"time"
)

type SysOssService struct{}

// AddSysOssUpload 上传文件
func (*SysOssService) AddSysOssUpload(c *gin.Context, req *pojo.UploadReq) {
	//上传文件目录名称 images：图片、files：excel、word、pdf等文件、videos：监控视频/音频(需要设置保留时间的)、photos：监控照片(需要设置保留时间的)
	if req.DirName != "images" && req.DirName != "files" && req.DirName != "videos" && req.DirName != "photos" {
		cwrs_res.Waring(c, nil, "上传目录错误，只能是images、files、videos、photos")
		return
	}
	dirName := req.DirName + "/" + time.Now().Format("200601")
	err, resp := cwrs_utils.AliyunOssGetUploadUrl(dirName, req.Postfix)
	if err != nil {
		cwrs_res.Waring(c, err, "上传失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", resp)
}

// GetSysOssDirList 查询对应目录下的所有子目录
func (*SysOssService) GetSysOssDirList(c *gin.Context, req *pojo.GetSysOssListReq) {
	err, files := cwrs_utils.AliyunOssGetSubDirs(req.DirName)
	if err != nil {
		cwrs_res.Waring(c, err, "查询文件列表失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", files)
}

// GetSysOssList 查询对应目录下所有文件列表
func (*SysOssService) GetSysOssList(c *gin.Context, req *pojo.GetSysOssListReq) {
	err, files := cwrs_utils.AliyunOssGetFileList(req.DirName)
	if err != nil {
		cwrs_res.Waring(c, err, "查询文件列表失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", files)
}
