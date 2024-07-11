package fofa

import (
	"GatInfo/Config"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"strings"
)

type AutoGenerated struct {
	Mode    string     `json:"mode"`
	Error   bool       `json:"error"`
	Query   string     `json:"query"`
	Page    int        `json:"page"`
	Size    int        `json:"size"`
	Results [][]string `json:"results"`
}

func fofa_api(keyword string, email string, key string, page int, size int) string {
	input := []byte(keyword)
	encodeString := base64.StdEncoding.EncodeToString(input)
	api_request := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&page=%d&size=%d&key=%s&qbase64=%s&fields=ip,host,title,port,protocol", strings.Trim(email, " "), page, size, strings.Trim(key, " "), encodeString)
	return api_request
}

func GetUrlsByCert(cert string) map[string]interface{} {
	result := make(map[string]interface{})
	var urls []string
	log.SetPrefix("[*] ")
	log.Println("Begin GetUrlsByCert:" + cert)
	requestUrl := fofa_api("cert=\""+cert+"\" && cert.is_valid=true && country=\"CN\"", Config.GetConfig().GetEmail(), Config.GetConfig().GetFofa_token(), 1, 10000)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetHeader("Accept", "*/*;q=0.8").
		SetHeader("Connection", "close").
		SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2227.0 Safari/537.36").
		Get(requestUrl)
	if err != nil {
		log.SetPrefix("[-] ")
		log.Println(err)
		return nil
	}
	var res AutoGenerated
	json.Unmarshal(resp.Body(), &res)
	if len(res.Results) > 0 {
		for _, value := range res.Results {
			if strings.Contains(value[1], "http") {
				urls = append(urls, value[1])
			} else if strings.Contains(value[1], "https") {
				urls = append(urls, "https://"+value[1])
			} else {
				urls = append(urls, "http://"+value[1])
			}

		}
	}
	log.SetPrefix("[*] ")
	log.Println("End GetUrlsByCert.Size:" + strconv.Itoa(len(res.Results)))
	result[cert] = urls
	return result
}