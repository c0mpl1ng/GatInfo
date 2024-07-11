package Config

import (
	"GatInfo/Utils/File"
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Email             string
	Fofa_token        string
	Fofa_timeout      string
	Censys_api_id     string
	Censys_api_secret string
}

var instance *Config

func GetConfig() *Config {
	if instance != nil {
		return instance
	}
	instance := &Config{}
	//创建一个结构体变量的反射
	cr := reflect.ValueOf(instance).Elem()
	//打开文件io流
	f, err := os.Open(File.GetCurrentAbPathByExecutable() + "/config.ini")
	if err != nil {
		//log.Fatal(err)
		log.SetPrefix("[-] ")
		log.Println("GetConfig Error:" + err.Error())
		os.Exit(1)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//我们要逐行读取文件内容
	s := bufio.NewScanner(f)
	for s.Scan() {
		//以=分割,前面为key,后面为value
		var str = s.Text()
		var index = strings.Index(str, "=")
		var key = str[0:index]
		var value = str[index+1:]
		//通过反射将字段设置进去
		cr.FieldByName(key).Set(reflect.ValueOf(value))
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return instance
	//创建一个空的结构体,将本地文件读取的信息放入

}

func (this *Config) GetEmail() string {
	return this.Email
}
func (this *Config) GetFofa_token() string {
	return this.Fofa_token
}
func (this *Config) GetCensys_api_id() string {
	return this.Censys_api_id
}
func (this *Config) GetCensys_api_secret() string {
	return this.Censys_api_secret
}
