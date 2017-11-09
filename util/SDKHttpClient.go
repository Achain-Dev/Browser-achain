package util

import (
	"net/http"
	"strings"
	"math/rand"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"encoding/base64"
)

type postData struct {
	Method  string
	Id      string
	Jsonrpc string
	Params  []string
}

func Post(url string, key string, method string, params []string) string {
	/*
	构造rpc传递json
	 */
	data := postData{
		Method:  method,
		Id:      strconv.Itoa(getRandomNumber()),
		Jsonrpc: "2.0",
		Params:  params,
	}
	postData := PostDataToString(data)

	/*
	处理json中出现null的问题
	 */
	if len(data.Params) == 0 {
		postData = strings.Replace(postData, "null", "[]", -1)
	}
	fmt.Println(postData)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	rpcAuth := strconv.Itoa(int((random.Float32()*9+1)*100000)) + base64.StdEncoding.EncodeToString([]byte(key))

	client := http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if err != nil {
		panic(err.Error())
		return ""
	}

	req.Header.Set("Content-type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", rpcAuth)

	fmt.Println("【SDKHttpClient】｜POST开始：url=", url)
	resp, err := client.Do(req)
	fmt.Println("【SDKHttpClient】｜POST开始 URL:[{", url, "}][method={", method, "}][jsonArray={", params, "}],响应结果[response={", resp, "}]!")
	if err != nil {
		panic("error get response")
		fmt.Println("【SDKHttpClient】｜POST URL:", url, "响应结果[{}]!", resp.StatusCode)
		return ""
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic("error get response body")
			return ""
		}
		fmt.Println("【SDKHttpClient】｜响应结果：{", resp.StatusCode, "},[{", string(body), "}]")
		resp.Body.Close()
		return string(body)
	} else {
		fmt.Println("【SDKHttpClient】｜POST URL:[{", url, "}],响应结果[{", resp.StatusCode, "}]!")
		return ""
	}
}

func getRandomNumber() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(100)
}

func PostDataToString(data interface{}) string {
	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json error!")
		return ""
	}
	return strings.ToLower(string(result[:]))
}
