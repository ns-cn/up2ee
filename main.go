package main

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var Root = &cobra.Command{
	Use:   "up2ee",
	Long:  "upload to gitee",
	Short: "upload to gitee",
	Run: func(cmd *cobra.Command, args []string) {
		isEmpty := func(value string, description string) bool {
			if value == "" {
				fmt.Printf("未配置%s\n", description)
				return true
			}
			return false
		}
		empty := isEmpty(UserName, "用户名") || isEmpty(UserPassword, "用户密码") || isEmpty(ClientId, "ClientId") || isEmpty(ClientSecret, "ClientSecret")
		if empty {
			return
		}
		size := len(args)
		if size == 0 {
			return
		}
		var err error
		accessData, err = access()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Println(accessData)
		if TestAccess {
			return
		}
		var images = make([]string, size)
		for _, arg := range args {
			image, err := readFile(arg)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			images = append(images, image)
		}
		downloadUrls := uploadImages(images)
		for _, url := range downloadUrls {
			if url != "" {
				fmt.Println(url)
			}
		}
	},
}

var (
	//user information required
	UserName     = ""
	UserPassword = ""
	ClientId     = ""
	ClientSecret = ""
	TestAccess   = false
	Repository   = ""
	// runtime data
	accessData AccessData
)

func init() {
	Root.Flags().StringVarP(&UserName, "username", "u", "", "用户名")
	Root.Flags().StringVarP(&UserPassword, "userpassword", "p", "", "用户密码")
	Root.Flags().StringVarP(&ClientId, "clientid", "c", "", "ClientId,通过gitee后台创建生成")
	Root.Flags().StringVarP(&ClientSecret, "clientsecret", "s", "", "ClientSecret,通过gitee后台创建生成")
	Root.Flags().StringVarP(&Repository, "repository", "r", "up2ee-data", "仓库,optional")
	Root.Flags().BoolVarP(&TestAccess, "testaccess", "t", false, "测试认证(仅认证,不做任何操作)")
}

func readFile(uri string) (string, error) {
	file, err := ioutil.ReadFile(uri)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(file), nil
}

func main() {
	//file, _ := ioutil.ReadFile("/Users/tangyujun/workspace/gitee.com/ns-cn/up2ee/iShot.png")
	//imageBase64 := base64.StdEncoding.EncodeToString(file)
	//fmt.Println(imageBase64)
	Root.Execute()
}
