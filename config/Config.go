package config

import (
	"errors"
	"fmt"
	"httpFileServer/utils/datautils"
	"httpFileServer/utils/ioutils"
	"strings"

	"github.com/unknwon/goconfig"
)

type ConfigData struct {
	Port          int
	FilePort      int
	ServedFolders []string
	NameSpaces    []string
	UploadPath    string
	UseDirectLink bool
	UseAuth       bool
	UserName      string
	Password      string
	LogAccess     bool
	LogFilePath   string
}

var _configData ConfigData
var CONFIGFILE string = "./http.conf"

var HTTP_SERVER_PORT string = "http.server.port" //http展示页端口
var FILE_SERVER_PORT string = "file.server.port" //文件下载的端口
var SERVED_FOLDERS string = "served.folders"     //
var NAME_SPACE string = "namespaces"             //
var UPLOAD_PATH string = "upload.path"           //上传路径
var USE_DIRECT_LINK string = "use.directlink"    //是否使用直链
var USE_AUTH string = " use.auth"                //是否使用登录验证方式访问
var USERNAME string = "username"                 //用户名
var PASSWORD string = "password"                 //密码
var LOG_ACCESS string = "log.access"             //打印访问历史
//LOG_FILE_PATH 打印文件路径
var LOG_FILE_PATH string = "log.file.path"

// GetConfig get all config copied
func GetConfig() (configData *ConfigData) {
	configData = &ConfigData{}
	datautils.DeepCopy(configData, _configData)
	return configData
}

//LoadConfig load config when init...
func LoadConfig(args []string) (configData *ConfigData) {
	//default
	defaultConfig := GetDefaultConfig()
	_configData = defaultConfig
	//http.conf
	fileConfig, _ := ParseConfigFromFile(CONFIGFILE)

	//-c,not use http.conf

	//Print config
}

func ParseConfigFromFile(path string) (defaultConfig *ConfigData, err error) {
	exists, _ := ioutils.PathExists(path)
	if !exists {
		return nil, errors.New(" ERROR -c file Not exists")
	}
	//default config
	defaultConfig := GetDefaultConfig()
	c, err := goconfig.LoadConfigFile(path)
	if err != nil {
		fmt.Println("ERROR can not open ", path)
		return nil, errors.New(" ERROR can not open  " + path + err)
	}
	//http 端口
	value, err := c.Int("", HTTP_SERVER_PORT, 8080)
	if err != nil {
		defaultConfig.Port = value
	} else {
		fmt.Println("WARN", "HTTP_SERVER_PORT must be  int value and not null,used default 8080")
	}

	value, err := c.Int("", FILE_SERVER_PORT, 8081)
	if err != nil {
		defaultConfig.FilePort = value
	} else {
		fmt.Println("WARN", "FILE_SERVER_PORT must be int value and not null,used default 8081 ")
	}

	value := c.MustValueArray("", SERVED_FOLDERS, ",")
	if len(value) > 0 {
		defaultConfig.ServedFolders = value
	} else {
		fmt.Println("WARN", "SERVED_FOLDERS used default value ")
	}

	value := c.MustValueArray("", NAME_SPACE, ",")
	defaultConfig.NameSpaces = make([]string, 0, len(defaultConfig.ServedFolders))
	for index, v := range defaultConfig.NameSpaces {
		if index > len(value)-1 { //超过自定义的Namepspace
			defaultConfig.NameSpaces[index] = defaultConfig.ServedFolders[0]
			continue
		} else {
			defaultConfig.NameSpaces = value[0]
		}
	}

	value, err := c.GetValue("", UPLOAD_PATH)
	if err != nil {
		defaultConfig.UploadPath = value
	} else {
		fmt.Println("ERROR", "UPLOAD_PATH must be not be NULL ")
	}

	value, err := c.Bool("", USE_DIRECT_LINK)
	if err != nil {
		defaultConfig.UseDirectLink = value
	} else {
		fmt.Println("ERROR", "USE_DIRECT_LINK must be  bool value and not null ")
	}

	value, err := c.Bool("", USE_AUTH)
	if err != nil {
		defaultConfig.UseAuth = value
	} else {
		fmt.Println("ERROR", "USE_AUTH must be  true or false value and not null ")
	}

	value, err := c.GetValue("", USERNAME)
	if err != nil {
		defaultConfig.UserName = value
	}

	value, err := c.GetValue("", PASSWORD)
	if err != nil {
		defaultConfig.Password = value
	}

	value, err := c.Bool("", LOG_ACCESS)
	if err != nil {
		defaultConfig.LogAccess = value
	}

	value, err := c.GetValue("", LOG_FILE_PATH)
	if err != nil {
		defaultConfig.LogFilePath = value
	}

	return defaultConfig, nil

}

func PrintlnHelp() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func ParseConfigFromCmd(args []string) (defaultConfig *ConfigData, isConfigFile bool, isSaved bool, err error) {
	if len(args) < 2 {
		return nil, false, false, nil
	}
	length := len(args)
	//search for -c
	var ConfigData defaultConfig
	var error err
	for index, arg := range args {
		if strings.Compare("-c", args) == 0 { //config File
			if index+1 >= length {
				return nil, false, nil, errors.New(" ERROR -c has no enough params")
			}
			filePath := args[index+1]
			defaultConfig, err = ParseConfigFromFile(path)
			break
		}
	}
	if err == nil {
		defaultConfig = new(ConfigData)
	}
	isCmdEnd := false
	isSaved := false
	for index, arg := range args {
		switch {

		case strings.Compare("-p", args) == 0: //port
			break
		case strings.Compare("-f", args) == 0: //served folder
			break
		case strings.Compare("-U", args) == 0: //userName
			break
		case strings.Compare("-P", args) == 0: //password
			break
		case strings.Compare("-l", args) == 0: //log path
			break
		case strings.Compare("-s", args) == 0: //log path
			isSaved = true
		}
	}

}

func GetDefaultConfig() (defaultConfig *ConfigData, err error) {
	//default
	defaultConfig := new(ConfigData)
	defaultConfig.Port = 8080
	defaultConfig.ServedFolders = []string{"./sharedFolder"}
	defaultConfig.NameSpaces = []string{""}
	defaultConfig.UploadPath = "./sharedFolder/upload"
	defaultConfig.UseDirectLink = true //启用直接连接
	defaultConfig.UseAuth = false
	defaultConfig.UserName = ""
	defaultConfig.Password = ""
	defaultConfig.LogAccess = false
	defaultConfig.LogFilePath = "./logs/"
	return defaultConfig, nil

}
