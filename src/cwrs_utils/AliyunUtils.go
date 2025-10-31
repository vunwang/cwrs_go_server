package cwrs_utils

import (
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	_ "path"
	"strings"
	"time"
)

// 声明全局变量（不赋值）
var (
	EndPoint        string
	RegionId        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	Host            string
)

var Credentials *sts.Credentials
var Expiration time.Time

type GetUploadURLResp struct {
	//FileName string `json:"fileName"`
	UploadUrl   string `json:"uploadUrl"`
	DownloadUrl string `json:"downloadUrl"`
}

// 初始化oss
func AliyunOssInit() {
	EndPoint = cwrs_viper.GlobalViper.GetString("aliYun.endPoint")
	RegionId = cwrs_viper.GlobalViper.GetString("aliYun.regionId")
	AccessKeyId = cwrs_viper.GlobalViper.GetString("aliYun.accessKeyId")
	AccessKeySecret = cwrs_viper.GlobalViper.GetString("aliYun.accessKeySecret")
	BucketName = cwrs_viper.GlobalViper.GetString("aliYun.bucket")
	Host = "https://" + BucketName + "." + EndPoint + "/" //plcsh.oss-cn-shanghai.aliyuncs.com

	// client *sts.Client
	client, err := sts.NewClientWithAccessKey(RegionId, AccessKeyId, AccessKeySecret)
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	// RAM访问控制->身份管理->角色
	request.RoleArn = "acs:ram::1770556036511338:role/tianyu"
	request.RoleSessionName = "tianyu"

	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Printf("client.AssumeRole:%v\n", err)
		return
	}

	Credentials = &response.Credentials
	Expiration, _ = time.Parse(time.RFC3339, Credentials.Expiration)
}

var ossClient *oss.Client = nil
var ossBucket *oss.Bucket = nil

// 初始化客户端和bucket
func GetClientAndBucket() (*oss.Client, *oss.Bucket, error) {
	var err error

	// 过期
	if time.Now().Add(10 * time.Second).After(Expiration) {
		AliyunOssInit()
		ossClient = nil
		ossBucket = nil
	}

	if ossBucket == nil {
		ossClient, err = oss.New(EndPoint,
			Credentials.AccessKeyId,
			Credentials.AccessKeySecret,
			oss.SecurityToken(Credentials.SecurityToken))
		if err != nil {
			fmt.Printf("oss.New:%v\n", err.Error())
			return nil, nil, err
		}
		ossBucket, err = ossClient.Bucket(BucketName)
		if err != nil {
			fmt.Printf("ossClient.Bucket:%v\n", err.Error())
			return nil, nil, err
		}
		return ossClient, ossBucket, nil
	} else {
		return ossClient, ossBucket, nil
	}
}

/**
 * 上传文件-获取上传签名（供前端上传）
 * dir: bucket中保存文件的文件夹
 * postfix: 文件后缀
 */
func AliyunOssGetUploadUrl(dirName string, postfix string) (error, GetUploadURLResp) {
	var resp GetUploadURLResp
	_, bucket, err := GetClientAndBucket()
	if err != nil {
		return err, resp
	}

	if strings.Contains(postfix, ".") {
		postfix = postfix[1:]
	}

	objectName := fmt.Sprintf("%s/%s%s.%s", dirName, time.Now().Format("02150405"), RandString(6), postfix)
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPPut, 240)
	if err != nil {
		return err, resp
	} else {
		resp.UploadUrl = signedUrl
		resp.DownloadUrl = Host + objectName
		//resp.FileName = objectName
		return err, resp
	}
}

/**
 * 删除文件
 * objectName: 文件名
 */
func AliyunOssDeleteObject(objectName string) error {
	if objectName == "" {
		return nil
	}
	objectName = objectName[len(Host):]
	_, bucket, err := GetClientAndBucket()
	if err != nil {
		return err
	}

	return bucket.DeleteObject(objectName)
}

/**
 * 设置文件生命周期
 * days: 天数
 */
func AliyunOssSetLifecycleRules(days int) error {
	//获取对应文件夹下的子文件夹列表
	_, dirs := getSubDirs()
	var rules []oss.LifecycleRule
	for _, dir := range dirs {
		rule := oss.LifecycleRule{
			ID:         "lifecycle-" + strings.Trim(dir, "/"), // 确保 ID 唯一
			Prefix:     dir,
			Status:     "Enabled",
			Expiration: &oss.LifecycleExpiration{Days: days},
		}
		rules = append(rules, rule)
	}
	client, _, err := GetClientAndBucket()
	if err != nil {
		return err
	}
	err = client.SetBucketLifecycle(BucketName, rules)
	return err
}

func getSubDirs() (error, []string) {
	_, bucket, err := GetClientAndBucket()
	if err != nil {
		return err, nil
	}
	dirs := cwrs_viper.GlobalViper.GetStringSlice("oss.dir")

	var dirRes []string
	for _, dir := range dirs {
		// 确保 dir 以 "/" 结尾，否则可能匹配到其他路径
		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
		result, err1 := bucket.ListObjects(
			oss.Prefix(dir),
			oss.Delimiter("/"),
		)
		if err1 != nil {
			return err1, nil
		}
		for _, prefix := range result.CommonPrefixes {
			dirRes = append(dirRes, prefix)
		}
	}
	return nil, dirRes
}

/**
 * 获取 OSS 中指定目录下的子文件夹列表（非递归）
 * dir: 例如 "videos、photos"
 */
func AliyunOssGetSubDirs(dir string) (error, []map[string]interface{}) {
	_, bucket, err := GetClientAndBucket()
	if err != nil {
		return err, nil
	}
	var dirRes []map[string]interface{}
	// 确保 dir 以 "/" 结尾，否则可能匹配到其他路径
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	result, err := bucket.ListObjects(
		oss.Prefix(dir),
		oss.Delimiter("/"),
	)
	if err != nil {
		return err, nil
	}
	for _, prefix := range result.CommonPrefixes {
		dirRes = append(dirRes, map[string]interface{}{
			"dirName": prefix,
		})
	}
	return nil, dirRes
}

/**
 * 获取文件列表
 * dir: bucket中保存文件的文件夹
 */
func AliyunOssGetFileList(dir string) (error, []map[string]interface{}) {
	_, bucket, err := GetClientAndBucket()
	var files []map[string]interface{}
	if err != nil {
		return err, files
	}

	result, err := bucket.ListObjects(oss.Prefix(dir))
	if err != nil {
		return err, files
	}

	for _, obj := range result.Objects {
		files = append(files, map[string]interface{}{
			"objectKey":   obj.Key,
			"size":        obj.Size,
			"lastMod":     obj.LastModified,
			"downloadUrl": fmt.Sprintf("%s%s", Host, obj.Key),
		})
	}
	return err, files
}
