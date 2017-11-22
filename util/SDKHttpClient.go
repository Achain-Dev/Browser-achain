package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type postData struct {
	Method  string
	Id      string
	JsonRpc string
	Params  []string
}

func Post(url string, key string, method string, params []string) string {
	// Construct RPC to deliver json
	data := postData{
		Method:  method,
		Id:      strconv.Itoa(getRandomNumber()),
		JsonRpc: "2.0",
		Params:  params,
	}
	postData := PostDataToString(data)

	// Handle the problem of null in json
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

	fmt.Println("【SDKHttpClient】｜POST begin：url=", url)
	resp, err := client.Do(req)
	fmt.Println("【SDKHttpClient】｜POST begin URL:[{", url, "}][method={", method, "}][jsonArray={", params, "}],reponse result[response={", resp, "}]!")
	if err != nil {
		panic("error get response")
		fmt.Println("【SDKHttpClient】｜POST URL:", url, "reponse result[{}]!", resp.StatusCode)
		return ""
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic("error get response body")
			return ""
		}
		fmt.Println("【SDKHttpClient】｜reponse result：{", resp.StatusCode, "},[{", string(body), "}]")
		resp.Body.Close()
		return string(body)
	} else {
		fmt.Println("【SDKHttpClient】｜POST URL:[{", url, "}],reponse result[{", resp.StatusCode, "}]!")
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
	res := strings.Replace(string(result[:]), "Method", "method", 1)
	res = strings.Replace(res, "Id", "id", 1)
	res = strings.Replace(res, "JsonRpc", "jsonrpc", 1)
	res = strings.Replace(res, "Params", "params", 1)
	return res
}
