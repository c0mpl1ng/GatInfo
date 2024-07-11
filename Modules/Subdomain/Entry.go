package Subdomain

import (
	"GatInfo/Modules/Subdomain/certificates"
	"GatInfo/Modules/Subdomain/fofa"
)

type FullUrlInfo struct {
	CertInfo []interface{}
	FofaInfo []interface{}
}

func removeDuplicates(strs []string) []string {
	unique := make(map[string]struct{})
	result := []string{}
	for _, str := range strs {
		if _, ok := unique[str]; !ok {
			result = append(result, str)
			unique[str] = struct{}{}
		}
	}
	return result
}

func GetSubdomain(targets []string) FullUrlInfo {
	var certList []string
	var certMapList []interface{}
	var urlsMapList []interface{}
	var Result FullUrlInfo
	//遍历target调用GetCertByCensys获取所有包含target的证书，返回map[string]interface格式，将这些map放到一个数组里
	for _, tar := range targets {
		certMapList = append(certMapList, certificates.GetCertByCensys(tar))
	}
	//将map数组中所有证书名提取到数组中然后去重
	for _, cert := range certMapList {
		for _, value := range cert.(map[string]interface{}) {
			value := value.([]string)
			certList = append(certList, value...)
		}

	}
	//去重
	certList = removeDuplicates(certList)
	//循环调用fofa,cert=上一步获取的证书名
	for _, cert := range certList {
		urlsMapList = append(urlsMapList, fofa.GetUrlsByCert(cert))
	}
	Result.CertInfo = certMapList
	Result.FofaInfo = urlsMapList
	return Result

}

// 通过手动指定的证书信息获取资产
func GetSubdomainBySpecifyCert(TargetcertList []string) FullUrlInfo {
	var certList []string
	var certMapList []interface{}
	var urlsMapList []interface{}
	var Result FullUrlInfo

	//去重
	certList = removeDuplicates(TargetcertList)
	//循环调用fofa
	for _, cert := range certList {
		urlsMapList = append(urlsMapList, fofa.GetUrlsByCert(cert))
	}
	Result.CertInfo = certMapList
	Result.FofaInfo = urlsMapList
	return Result

}
