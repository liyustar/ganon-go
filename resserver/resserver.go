// resserver 是表示Oauth下的资源服务器
package resserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Config struct {
	ClientId string
	ClientSecret string
	RedirectUrl string
}

type Context struct {
	Conf Config
}

func (c *Context) Hello(w http.ResponseWriter, _ *http.Request) {
	// 解析制定文件生成模板对象
	var temp *template.Template
	var err error
	if temp, err = template.ParseFiles("views/hello.html"); err != nil {
		fmt.Println("读取文件失败，错误信息为：", err)
		return
	}

	// 利用给定数据渲染模板（html页面），并将结果写入w，返回给前端
	if err = temp.Execute(w, c.Conf); err != nil {
		fmt.Println("读取渲染html页面失败，错误信息为：", err)
		return
	}
}

func (c *Context) Oauth(w http.ResponseWriter, r *http.Request) {
	var err error
	var code = r.URL.Query().Get("code") // 获取code
	fmt.Println(err, code)

	var tokenAuthUrl = c.GetTokenAuthUrl(code)
	var token *Token
	if token, err = GetToken(tokenAuthUrl); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("token: %+v", token)

	var userInfo map[string]interface{}
	if userInfo, err = GetUserInfo(token); err != nil {
		fmt.Println("获取用户信息失败，错误信息为：", err)
		return
	}

	var userInfoBytes []byte
	if userInfoBytes, err = json.Marshal(userInfo); err != nil {
		fmt.Println("在将用户信息（map）转为用户信息（[]byte）时发生错误，错误信息为：", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(userInfoBytes); err != nil {
		fmt.Println("在将用户信息([]byte)返回前端时发生错误，错误信息为：", err)
		return
	}
}

func (c *Context) GetTokenAuthUrl(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		c.Conf.ClientId, c.Conf.ClientSecret, code)
}

func GetToken(url string) (*Token, error) {
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err!= nil {
		return nil, err
	}

	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
}

func GetUserInfo(token *Token) (map[string]interface{}, error) {
	var userInfoUrl = "https://api.github.com/user"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

