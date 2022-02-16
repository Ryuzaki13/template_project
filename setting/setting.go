package setting

import (
	"encoding/json"
	"fmt"
	"os"
)

type Setting struct {
	Address string
	Port    string
}

var cfg Setting

func Load(filename string) *Setting {
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println(e.Error())
		panic("не удалось прочитать файл параметров")
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic("не удалось прочитать информацию о файле параметров")
	}

	bytes := make([]byte, stat.Size())

	n, e := file.Read(bytes)
	if e != nil || n != int(stat.Size()) {
		fmt.Println(e.Error())
		panic("не удалось прочитать файл")
	}

	e = json.Unmarshal(bytes, &cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("не удалось загрузить параметры")
	}

	return &cfg
}
