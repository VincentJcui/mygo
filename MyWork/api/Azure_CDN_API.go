package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/widuu/gojson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

//定义一个结构体
type BodyFiles struct {
	Files       []string
	Directories []string
}

var utc_time = time.Now().UTC().Unix()

//var bodyfiles = make(map[string]*BodyFiles)
var requestTimeStamp = time.Now().UTC().Format("2006-01-02 15:04:05")

//针对Azure的https请求ssl重连的问题,需要对http的参数做调整
var tr = &http.Transport{
	//TLSClientConfig: &tls.Config{InsecureSkipVerify: true,Renegotiation: tls.RenegotiateFreelyAsClient,},
	TLSClientConfig: &tls.Config{Renegotiation: tls.RenegotiateFreelyAsClient},
}
var client = &http.Client{Transport: tr}

func calculateAuthorizationHeader(requestURL, request_timestamp, keyID, keyValue, httpMethod string) string {
	u, err := url.Parse(requestURL)
	if err != nil {
		panic(err)
	}

	var path = u.Path
	m, _ := url.ParseQuery(u.RawQuery)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var orderedQueries []string
	for _, k := range keys {
		orderedQueries = append(orderedQueries, fmt.Sprintf("%s:%s", k, m[k][0]))
	}

	var queries = strings.Join(orderedQueries, ", ")
	content := fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s", path, queries, requestTimeStamp, httpMethod)
	hash := hmac.New(sha256.New, []byte(keyValue))
	hash.Write([]byte(content))
	digest := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	return fmt.Sprintf("AzureCDN %s:%s", keyID, digest)
}

func postapi(date string, Authorization string, uri string, body []byte) {
	//fmt.Println("-------------------------------------------")
	//fmt.Println(date)
	//fmt.Println(Authorization)
	//fmt.Println(uri)
	//fmt.Println(body)
	//fmt.Println(string(body))
	//fmt.Println(bytes.NewBuffer(body))
	//fmt.Println(strings.NewReader(string(body)))
	//fmt.Println(bytes.NewReader(body))
	//fmt.Println("-------------------------------------------")
	//req, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	//req, err := http.NewRequest("POST", uri, strings.NewReader(string(body)))       //普通的post请求
	req, err := http.NewRequest("POST", uri, bytes.NewReader(body)) //body全部二进制数据流进行post
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-azurecdn-request-date", date)
	req.Header.Set("Authorization", Authorization)
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}
	if resp.StatusCode != 202 {
		bodys, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bodys))
	} else {
		fmt.Println("Succeeded")
	}
}

func EndpointID(date string, Authorization string) string {
	uri := "https://restapi.cdn.azure.cn/subscriptions/ee57ca15-7f92-4a93-ac70-f0888a2fb0e3/endpoints?apiVersion=1.0"
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-azurecdn-request-date", date)
	req.Header.Set("Authorization", Authorization)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	x := gojson.Json(string(body))
	//fmt.Println(x)
	_, endpointid := x.ToArray()
	fmt.Println()
	if endpointid[0] == "c766f846-71ea-11e7-8259-0017fa00a611" {
		return endpointid[0]
	} else {
		return endpointid[1]
	}
}

func main() {
	//定义变量
	keyId := "0d1ff2a1-285c-45ea-8328-1623ed865cac"
	keyValue := "ZmNmNDExNDgtMWY4Yi00MmEyLWE0YWQtMzBiYjM2MTZhN2E5"

	//首先先获取endpoinitd,使用GET方法
	geteid_uri := "https://restapi.cdn.azure.cn/subscriptions/ee57ca15-7f92-4a93-ac70-f0888a2fb0e3/endpoints?apiVersion=1.0"
	geteid_Authorization := calculateAuthorizationHeader(geteid_uri, string(utc_time), keyId, keyValue, "GET")
	endpointid := EndpointID(requestTimeStamp, geteid_Authorization)
	//fmt.Println(endpointid)
	//获取endpoinitd之后重新获取Authorization,使用POST方法
	refresh_uri := "https://restapi.cdn.azure.cn/subscriptions/ee57ca15-7f92-4a93-ac70-f0888a2fb0e3/endpoints/" + endpointid + "/purges?apiVersion=1.0"
	refresh_Authorization := calculateAuthorizationHeader(refresh_uri, requestTimeStamp, keyId, keyValue, "POST")
	bb := []string{}
	if len(os.Args) < 1 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "-f":
		files := os.Args[2]
		f := append(bb, files)
		bodyfiles := BodyFiles{
			Files: f,
		}
		Body, err := json.Marshal(bodyfiles)
		if err != nil {
			log.Fatal(err)
		}
		postapi(requestTimeStamp, refresh_Authorization, refresh_uri, Body)
	case "-a":
		files := os.Args[2]
		dirs := os.Args[3]
		f := append(bb, files)
		d := append(bb, dirs)
		bodyfiles := BodyFiles{
			Files:       f,
			Directories: d,
		}
		Body, err := json.Marshal(bodyfiles)
		if err != nil {
			log.Fatal(err)
		}
		postapi(requestTimeStamp, refresh_Authorization, refresh_uri, Body)
	case "-d":
		dirs := os.Args[2]
		d := append(bb, dirs)
		bodyfiles := BodyFiles{
			Directories: d,
		}
		Body, err := json.Marshal(bodyfiles)
		if err != nil {
			log.Fatal(err)
		}
		postapi(requestTimeStamp, refresh_Authorization, refresh_uri, Body)
	default:
		fmt.Printf(`Help:
.%s -f http://fps.ms.kukuplay.com/common/4399/bg8.jpg
.%s -d http://fps.ms.kukuplay.com/res4399
.%s -a http://fps.ms.kukuplay.com/common/4399/bg8.jpg http://fps.ms.kukuplay.com/res4399
		`, os.Args[0], os.Args[0], os.Args[0])
	}
}
