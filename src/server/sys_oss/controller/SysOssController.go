package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_oss/pojo"
	"cwrs_go_server/src/server/sys_oss/service"

	"github.com/gin-gonic/gin"
)

var sysOssServiceImpl = service.SysOssService{}

// @Tags 上传文件【平台&APP】
// @Summary 上传文件
// @Description 上传文件
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.UploadReq true "上传参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=cwrs_utils.GetUploadURLResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysOss/upload [get]
func AddSysOssUpload(c *gin.Context) {
	var req pojo.UploadReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysOssServiceImpl.AddSysOssUpload(c, &req)
}

// @Tags 上传文件【平台&APP】
// @Summary 查询文件夹列表
// @Description 查询文件夹列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysOssListReq true "查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=pojo.SysPostResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysOss/dirList [get]
func GetSysOssDirList(c *gin.Context) {
	var req pojo.GetSysOssListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysOssServiceImpl.GetSysOssDirList(c, &req)
}

//// @Tags 上传文件【平台&APP】
//// @Summary 查询文件列表
//// @Description 查询文件列表
//// @Accept json
//// @Produce json
//// @Param Token header string true "token"
//// @Param req query pojo.GetSysOssListReq true "查询参数"
//// @Success 200 {object} cwrs_res.ResSuccessDataList{data=pojo.SysPostResp} "操作成功"
//// @Failure 422 {object} cwrs_res.ResError "操作失败"
//// @Router /sysOss/fileList [get]
//func GetSysOssList(c *gin.Context) {
//	var req pojo.GetSysOssListReq
//	if err := c.ShouldBindQuery(&req); err != nil {
//		cwrs_res.Parameter(c, err, "参数错误！")
//		return
//	}
//	sysOssServiceImpl.GetSysOssList(c, &req)
//}

// @Tags 上传文件【平台&APP】
// @Summary 删除文件
// @Description 删除文件
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param ObjectName query string true "查询参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysOss/del [delete]
func DeleteSysOss(c *gin.Context) {
	objectName := c.Query("ObjectName")
	err := cwrs_utils.AliyunOssDeleteObject(objectName)
	if err != nil {
		cwrs_res.Waring(c, err, "删除oss文件失败")
		return
	}
	cwrs_res.Success(c, "删除成功")
}
