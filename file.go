package automod

import (
	"io/ioutil"
	"log"
	"os"
)

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建文件
func createFile(filename string) (bool, error) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return true,nil
	}

	return false,nil
}

func createModelFile(path string, filename string, ormType string, content string) bool {
	//判断文件夹是否存在
	exist,err := pathExists(path)

	if err != nil {
		log.Fatal(err)
	}

	if !exist {
		//创建文件夹
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	filename = path + "/" + filename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(content))

	if err != nil {
		return false
	}

	return true

}

func readAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

//将首字母大写，下划线_去掉改为大写，例子ab_cd -> AbCd
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
