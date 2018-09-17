package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"go.uber.org/zap"
	"github.com/xfstart07/watcher/util"
	"time"
	"github.com/xfstart07/watcher/config"
)

var (
	zlog = util.NewLog()
)

const filechunk = 8192 // we settle for 8KB


func StorePaths() {
	paths := config.Config.WatchPaths
	for idx := range paths {
		StoreFiles(paths[idx])
	}
}

func StoreFiles(path string) {

	if !util.PathExist(path) {
		return
	}

	files, err := util.GetFileNames(path)
	if err != nil {
		zlog.Sugar().Error("获取文件夹中文件名失败", zap.Error(err))
		return
	}

	for idx := range files {
		Store(files[idx])
	}

}

func Store(filepath string) {
	if !util.PathExist(filepath) {
		zlog.Error("文件不存在", zap.String("path", filepath))
		return
	}

	if exist, _ := Redis.Exists(filepath).Result(); exist > 0 {
		zlog.Info("任务已经存在")
		return
	}

	// 同一个路径，1分半钟只计算一次
	Redis.SetNX(filepath, 1, time.Duration(90))

	files := strings.Split(filepath, "/")
	fileName := files[len(files)-1]

	//if exist, _ := Redis.Exists(fileName).Result(); exist > 0  {
	//	zlog.Info("文件MD5已经存在")
	//	return
	//}

	md5, err := sumFile(filepath)
	if err != nil {
		zlog.Error("计算MD5失败", zap.Error(err))
	}

	// store redis
	storeFile(fileName, md5)

	zlog.Info("MD5", zap.String("md5", md5))
}

// store redis
func storeFile(fileName, md5 string) {
	// FIXME: 永久保存, 应该有过期时间
	//expire := 7*24*time.Hour
	expire := time.Duration(0)
	Redis.Set(fileName, md5, expire)
}

func sumFile(filepath string) (string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		zlog.Error("文件打开失败", zap.Error(err))
		return "", err
	}
	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()
	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}
	cipherText := string(hash.Sum(nil))
	cipherMd5 := fmt.Sprintf("%x", cipherText)

	return strings.ToUpper(cipherMd5), nil
}
