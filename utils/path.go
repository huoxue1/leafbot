package utils

import (
	"os"
)

// PathExists
/**
* @Description:
  判断文件或文件夹是否存在
  如果返回的错误为nil,说明文件或文件夹存在
  如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
  如果返回的错误为其它类型,则不确定是否在存在
* @param path
* @return bool
* @return error
* example
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteToFile(name string, data []byte) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err1 := file.Close()
		if err1 != nil {
			if err1 != nil {
				err = err1
			}
		}
	}(file)
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return err
}
