package certificates

import (
	"GatInfo/Config"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"time"
)

func getToken() string {
	C := Config.GetConfig()
	C.GetEmail()
	api := Config.GetConfig().GetCensys_api_id()
	secret := Config.GetConfig().GetCensys_api_secret()
	token := base64.StdEncoding.EncodeToString([]byte(api + ":" + secret))
	return token
}
func GetCertByCensys(target string) map[string]interface{} {
	host := "https://search.censys.io/api"
	CertResult := make(map[string]interface{})
	CertList := []string{}
	jsonData := map[string]interface{}{
		"q":        "parsed.subject_dn:" + target,
		"per_page": 100,
	}

	log.SetPrefix("[*] ")
	log.Println("Begin GetCertByCensys:" + target)
	client := resty.New()
	client.SetTimeout(30 * time.Second)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	//client.SetProxy("http://127.0.0.1:7777")
	resp, err := client.R().
		SetHeader("Authorization", "Basic "+getToken()).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0").
		SetBody(jsonData).
		Post(host + "/v2/certificates/search")
	if err != nil {
		log.SetPrefix("[-] ")
		log.Println(err)
		return GetCertByCensys(target)
	}

	var result map[string]interface{}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.SetPrefix("[-] ")
		log.Println(err)
		log.Println(resp.Body())
		return CertResult
	}
	if data, ok := result["result"].(map[string]interface{}); ok {
		if hits, ok := data["hits"].([]interface{}); ok {
			for _, hit := range hits {
				if hitMap, ok := hit.(map[string]interface{}); ok {
					if certList, ok := hitMap["names"].([]interface{}); ok {
						for _, cert := range certList {
							CertList = append(CertList, cert.(string))
						}
					}
				}

			}
		}

	}
	log.SetPrefix("[*] ")
	log.Println("End GetCertByCensys.Size:" + strconv.Itoa(len(CertList)))
	CertResult[target] = CertList
	//延迟两秒，避免被封IP
	time.Sleep(2 * time.Second)
	return CertResult
}
