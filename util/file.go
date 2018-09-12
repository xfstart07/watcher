// Author: Xu Fei
// Date: 2018/9/12
package util

import "io/ioutil"

// 获取文件夹下所有文件
func GetFileNames(dirPth string) (files []string, err error) {
	files = make([]string, 0, 100)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}

		files = append(files, dirPth + "/" + fi.Name())
	}

	return files, nil
}