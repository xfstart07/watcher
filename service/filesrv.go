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

const filechunk = 8192 // 数据块 8KB

func StorePaths() {
	paths := config.Config.WatchPaths
	for idx := range paths {
		StoreFiles(paths[idx])
	}
}

// 计算现有的文件
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

// 计算MD5并存入DB
func Store(filePath string) {
	if !util.PathExist(filePath) {
		zlog.Error("文件不存在", zap.String("path", filePath))
		return
	}

	if exist, _ := Redis.Exists(filePath).Result(); exist > 0 {
		zlog.Info("任务已经存在")
		return
	}

	// 同一个路径，1分半钟只计算一次
	Redis.SetNX(filePath, 1, 1*time.Minute)

	files := strings.Split(filePath, "/")
	fileName := files[len(files)-1]

	if exist, _ := Redis.Exists(fileName).Result(); exist > 0  {
		zlog.Info("文件MD5已经存在")
		return
	}

	// FIXME: 文件的写入会触发多个事件，所以延迟1分钟等文件写入完成后执行计算
	time.AfterFunc(1*time.Minute, func() {
		md5, err := calcFile(filePath)
		if err != nil {
			zlog.Error("计算MD5失败", zap.Error(err))
		}

		// store redis
		storeMD5(fileName, md5)

		zlog.Info("MD5", zap.String("md5", md5))
	})
}

// store redis
// expire 7 days
func storeMD5(fileName, md5 string) {
	expire := 7*24*time.Hour
	Redis.Set(fileName, md5, expire)
}

// calculate File MD5
func calcFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		zlog.Error("文件打开失败", zap.Error(err))
		return "", err
	}
	defer file.Close()

	// calculate the file size
	info, err := file.Stat()
	if err != nil {
		zlog.Error("File Stat Err", zap.Error(err))
		return "", err
	}

	fileSize := info.Size()
	blocks := uint64(math.Ceil(float64(fileSize) / float64(filechunk)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(filechunk, float64(fileSize-int64(i*filechunk))))
		buf := make([]byte, blockSize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}
	cipherText := string(hash.Sum(nil))
	cipherMd5 := fmt.Sprintf("%x", cipherText)

	return strings.ToUpper(cipherMd5), nil
}
