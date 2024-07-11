package BugAssets

import (
	"GatInfo/Modules/BugAssets/fofa"
	"GatInfo/Utils/File"
	"fmt"
	"log"
)

// 定义一个特定的系统指纹，比如畅捷通的某一个版本。该类型特指一个系统漏洞
type Finger struct {
	Type string //指纹类型，fofa,hunter...
	Link string //资产漏洞链接
	Fg   string //具体指纹
	urls []string
}

// 定义一个资产类型，比如畅捷通，Fingers代表多个指纹，因为畅捷通有很多版本类型，并不是只有一种指纹
type Feature struct {
	Name    string //资产通用名，比如畅捷通，金蝶
	Fingers []Finger
}

var featurMap map[string]Feature

func init() {
	featurMap = map[string]Feature{
		"畅捷通": {
			Name: "畅捷通",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203435&highlight=%E7%95%85%E6%8D%B7%E9%80%9A", Fg: "app=\"畅捷通-TPlus\""},
				{Type: "shodan", Link: "https://forum.ywhack.com/viewthread.php?tid=202708&highlight=%E7%95%85%E6%8D%B7%E9%80%9A", Fg: "http.html:\"location='/tplus/'"},
			},
		},
		"金蝶云星空": {
			Name: "金蝶云星空",
			Fingers: []Finger{
				{Type: "hunter", Link: "https://forum.ywhack.com/viewthread.php?tid=203423&highlight=%E9%87%91%E8%9D%B6", Fg: "app.name=\"Kingdee 金蝶云星空\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203233&highlight=%E9%87%91%E8%9D%B6", Fg: "fid=\"rKifsBySfhmVNS3z144vZA==\" || title=\"金蝶云星空 管理中心\" || \"K3Cloud\""},
			},
		},
		"金蝶 EAS": {
			Name: "金蝶 EAS",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203129&highlight=%E9%87%91%E8%9D%B6", Fg: "Body=\"金蝶\" && Title=\"eas系统登录\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203169&highlight=%E9%87%91%E8%9D%B6", Fg: "app=\"Kingdee-EAS\""},
			},
		},
		"用友": {
			Name: "用友",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203548&highlight=%E7%94%A8%E5%8F%8B", Fg: "app=\"用友-UFIDA-NC\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203399&highlight=%E7%94%A8%E5%8F%8B", Fg: "title=\"产品登录界面\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203385&highlight=%E7%94%A8%E5%8F%8B", Fg: "title=\"U9-登录    \""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203335&highlight=%E7%94%A8%E5%8F%8B", Fg: "body=\"nc.sfbase.applet.NCApplet.class\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203284&highlight=%E7%94%A8%E5%8F%8B", Fg: "app=\"用友-U8-Cloud\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203284&highlight=%E7%94%A8%E5%8F%8B", Fg: "body=\"../js/jslib/jquery.blockUI.js\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203330&highlight=%E7%94%A8%E5%8F%8B", Fg: "Body=\"用友GRP-U8行政事业内控管理软件\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/lTalchm3Q7252QlEAVBxNQ", Fg: "body=\"http://uclient.yonyou.com/api/uclient/public/download/windows\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201138&highlight=KSOA", Fg: "product=\"用友-时空KSOA"},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/70Kv46CnlO1FJXeH8e0sag", Fg: "title=\"用友U8CRM\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/70Kv46CnlO1FJXeH8e0sag", Fg: "app=\"用友-移动系统管理\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/CVehl_f1X5cGU4gef_4j0g", Fg: "icon_hash=\"1085941792\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/JDvnGfnLruPrEOeRrBaBAg", Fg: "icon_hash=\"1596996317\" || icon_hash=\"737647198\""},
			},
		},
		"速达软件": {
			Name: "速达软件",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/thread-203338-1-8.html", Fg: "app=\"速达软件-公司产品\""},
			},
		},
		"H3C": {
			Name: "H3C",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/thread-202964-1-19.html", Fg: "app=\"H3C-Ent-Router\""},
				{Type: "hunter", Link: "https://forum.ywhack.com/thread-202964-1-19.html", Fg: "app.name=\"H3C Router Management\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203409&highlight=H3C", Fg: "fid=\"tPmVs5PL6e9m5Xt0J4V2+A==\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203404&highlight=H3C", Fg: "title=\"用户自助服务\""},
				{Type: "hunter", Link: "https://forum.ywhack.com/viewthread.php?tid=202760&highlight=H3C", Fg: "web.body=\"/imc/login.jsf\""},
				{Type: "hunter", Link: "https://forum.ywhack.com/viewthread.php?tid=201961&highlight=H3C", Fg: "web.icon==\"ae2c0f3b4de0023e5e9cf9a09d80fbd1\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201270&highlight=H3C", Fg: "app=\"H3C-ER5200G2\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201182&highlight=H3C", Fg: "app=\"H3C-CVM\" "},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=115708&highlight=H3C", Fg: "app=\"H3C-SecPath-运维审计系统\" && body=\"2018\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201113&highlight=H3C", Fg: "app=\"H3C-CAS\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=115598&highlight=H3C", Fg: "\"/imc/login.jsf\" && body=\"/imc/javax.faces.resource/images/login_help.png.jsf?ln=primefaces-imc-new-webui\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=115494&highlight=H3C", Fg: "app=\"H3C-SecPath-运维审计系统\""},
			},
		},
		"大华": {
			Name: "大华",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203384&highlight=%E5%A4%A7%E5%8D%8E", Fg: "app=\"dahua-智慧园区综合管理平台\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/g9b0aCEkDCPP-dd0cb8_bw", Fg: "app=\"dahua-DSS\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/Ch1AXG3CSXcJkQyYvh9XMw", Fg: "body=\"static/fontshd/font-hd.css\" || body=\"客户端会小于800\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/zIe552z9cklzYyp75SJ3Hg", Fg: "icon_hash=\"-1935899595\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203143&highlight=%E5%A4%A7%E5%8D%8E", Fg: "\"static/qwebchannel.js\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203143&highlight=%E5%A4%A7%E5%8D%8E", Fg: "icon_hash=\"289418320\""},
				{Type: "hunter", Link: "https://forum.ywhack.com/viewthread.php?tid=202874&highlight=%E5%A4%A7%E5%8D%8E", Fg: "web.icon==\"9fa8a9035ce4baa7eee40725b5cfed16\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202874&highlight=%E5%A4%A7%E5%8D%8E", Fg: "icon_hash=\"411052691\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202798&highlight=%E5%A4%A7%E5%8D%8E", Fg: "app=\"dahua-EIMS\""},
			},
		},
		"jeecg": {
			Name: "jeecg",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203515&highlight=jeecg", Fg: "app=\"JEECG\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203026&highlight=jeecg", Fg: "app=\"JeecgBoot-企业级低代码平台\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202924&highlight=jeecg", Fg: "body=\"jeecg-boot\""},
			},
		},
		"时空智友": {
			Name: "时空智友",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202997&highlight=%E6%97%B6%E7%A9%BA", Fg: "body=\"企业流程化管控系统\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202997&highlight=%E6%97%B6%E7%A9%BA", Fg: "body=\"login.jsp?login\" && body=\"登录\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/thread-203612-1-1.html", Fg: "body=\"login.jsp?login=null\""},
			},
		},
		"中成科信": {
			Name: "中成科信",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203501&highlight=%E4%B8%AD%E6%88%90", Fg: "body=\"中成科信\""},
			},
		},
		"泛微": {
			Name: "泛微",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203324&highlight=%E6%B3%9B%E5%BE%AE", Fg: "app=\"泛微-EOffice\" || body=\"/general/login/index.php\" || body=\"eoffice-loading-tip\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203185&highlight=%E6%B3%9B%E5%BE%AE", Fg: "body=\"eoffice10\" && body=\"eoffice_loading_tip\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203246&highlight=%E6%B3%9B%E5%BE%AE", Fg: "title=\"移动管理平台-企业管理\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203005&highlight=%E6%B3%9B%E5%BE%AE", Fg: "icon_hash=\"2062026853\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202532&highlight=%E6%B3%9B%E5%BE%AE", Fg: "app=\"泛微-EOffice\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201180&highlight=%E6%B3%9B%E5%BE%AE", Fg: "app=\"泛微-EMobile\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/4HMIH-cu7i1pts1qItIq9g", Fg: "body=\"/newplugins/js/pnotify/jquery.pnotify.default.css\""},
			},
		},
		"瑞友天翼": {
			Name: "瑞友天翼",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203574&highlight=%E8%87%B4%E8%BF%9C", Fg: "app=\"REALOR-天翼应用虚拟化系统\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202365&highlight=%E7%91%9E%E5%8F%8B%E5%A4%A9%E7%BF%BC", Fg: "body=\"/CasMain.XGI\""},
			},
		},
		"致远": {
			Name: "致远",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203307&highlight=%E8%87%B4%E8%BF%9C", Fg: "body=\"M3-Server 已启动\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203034&highlight=%E8%87%B4%E8%BF%9C", Fg: "app=\"致远互联-OA\""},
			},
		},
		"海康": {
			Name: "海康",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203495&highlight=%E6%B5%B7%E5%BA%B7", Fg: "web.body=\"/portal/skin/ifar/blue/skin.css\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203495&highlight=%E6%B5%B7%E5%BA%B7", Fg: "header=\"X-Content-Type-Options: nosniff\" && body=\"<h1>Welcome to OpenResty!</h1>\" && header=\"X-Xss-Protection: 1; mode=block\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203173&highlight=%E6%B5%B7%E5%BA%B7", Fg: "icon_hash=\"136203464\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202687&highlight=%E6%B5%B7%E5%BA%B7", Fg: "app=\"HIKVISION-iVMS\""},
				{Type: "hunter", Link: "https://forum.ywhack.com/viewthread.php?tid=202746&highlight=%E6%B5%B7%E5%BA%B7", Fg: "body=\"stars1_canvas\" && body=\"/js/plugin/sha256.js\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202725&highlight=%E6%B5%B7%E5%BA%B7", Fg: "app=\"HIKVISION-iSecure-Center\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202609&highlight=%2Fels%2Fstatic", Fg: "product=\"HIKVISION-综合安防管理平台\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202609&highlight=%2Fels%2Fstatic", Fg: "icon_hash=\"-1605849932\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202609&highlight=%2Fels%2Fstatic", Fg: "body=\"dist/jquery.js\" && body=\"home/locationIndex.action\""},
			},
		},
		"电信网关": {
			Name: "电信网关",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203579&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"img/login_bg3.png\" && body=\"系统登录\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/9hn_cA-4QWtfLJhl97uQkA", Fg: "body=\"a:link{text-decoration:none;color:orange;}\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202624&highlight=%E7%94%B5%E4%BF%A1%E7%BD%91%E5%85%B3", Fg: "body=\"<TITLE>系统登录\" && body=\"login.php\" && body=\"img/login_bg3.png\""},
			},
		},
		"和丰多媒体": {
			Name: "和丰多媒体",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203573&highlight=%E4%B8%8A%E4%BC%A0", Fg: "app=\"和丰山海-数字标牌\""},
			},
		},
		"建研信息": {
			Name: "建研信息",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203570&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"/Content/Theme/Standard/webSite/login.css\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203535&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"/Content/Theme/Standard/\""},
			},
		},
		"易宝OA": {
			Name: "易宝OA",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203529&highlight=%E4%B8%8A%E4%BC%A0", Fg: "title=\"欢迎登录易宝OA系统\""},
			},
		},
		"润乾报表": {
			Name: "润乾报表",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203526&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"润乾报表\" || body=\"/raqsoft\""},
			},
		},
		"可视化融合": {
			Name: "可视化融合",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203142&highlight=%E5%8F%AF%E8%A7%86%E5%8C%96%E8%9E%8D%E5%90%88", Fg: "body=\"base/searchInfoWindow_min.css\""},
			},
		},
		"魔方网表": {
			Name: "魔方网表",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203499&highlight=%E4%B8%8A%E4%BC%A0", Fg: "icon_hash=\"694014318\""},
			},
		},
		"云课网校": {
			Name: "云课网校",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203494&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"/static/libs/common/jquery.stickyNavbar.min.js\""},
			},
		},
		"通天星": {
			Name: "通天星",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203483&highlight=%2Fopen%2FwebApi.html", Fg: "body=\"./open/webApi.html\" || body=\"/808gps/\""},
			},
		},
		"飞企互联": {
			Name: "飞企互联",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203456&highlight=%E4%B8%8A%E4%BC%A0", Fg: "body=\"flyrise.stopBackspace.js\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203029&highlight=%09%2B%E9%A3%9E%E4%BC%81%E4%BA%92%E8%81%94", Fg: "icon_hash=\"-391577146\""},
			},
		},
		"明御安全网关": {
			Name: "明御安全网关",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202985&highlight=VPN", Fg: "product=\"安恒信息-明御安全网关\""},
			},
		},
		"MilesightVPN": {
			Name: "MilesightVPN",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202801&highlight=VPN", Fg: "title==\"MilesightVPN\""},
			},
		},
		"奇安信 VPN": {
			Name: "奇安信 VPN",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202618&highlight=VPN", Fg: "icon_hash=\"663709625\""},
			},
		},
		"红海eHR人力资源管理系统": {
			Name: "红海eHR人力资源管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201139&highlight=%E7%BA%A2%E6%B5%B7", Fg: "body=\"RedseaPlatform\" && icon_hash=\"-1015496453\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/3nTtwjAsDLdo7Pqnq7b_EA", Fg: "body=\"/RedseaPlatform/skins/images/favicon.ico\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/3nTtwjAsDLdo7Pqnq7b_EA", Fg: "body=\"/RedseaPlatform/skins/images/favicon.ico\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/3nTtwjAsDLdo7Pqnq7b_EA", Fg: "body=\"/RedseaPlatform/skins/images/favicon.ico\""},
			},
		},
		"中科聚网": {
			Name: "中科聚网",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203457&highlight=%E4%B8%8A%E4%BC%A0", Fg: "title=\"一体化运营平台\" || body=\"thirdparty/ueditor/WordPaster\""},
			},
		},
		"亿赛通": {
			Name: "亿赛通",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203457&highlight=%E4%B8%8A%E4%BC%A0", Fg: "app=\"亿赛通-电子文档安全管理系统\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203457&highlight=%E4%B8%8A%E4%BC%A0", Fg: "title=\"电子文档安全管理系统\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203457&highlight=%E4%B8%8A%E4%BC%A0", Fg: "product=\"亿赛通-DLP\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/asfHtCd_eJfrJaPnHJ8P0A", Fg: "body=\"CDGServer3\" || title=\"电子文档安全管理系统\" || cert=\"esafenet\" ||body=\"/help/getEditionInfo.jsp\"||body=\"/CDGServer3/index.jsp\""},
			},
		},
		"英飞达医学影像存档与通信系统": {
			Name: "英飞达医学影像存档与通信系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/fcupsevB2FytCwiHCYRXXw", Fg: "icon_hash=\"1474455751\" || icon_hash=\"702238928\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/lO9NAWQxlyht2ZEPe8M7Mg", Fg: "\"INFINITT\" && (icon_hash=\"1474455751\" || icon_hash=\"702238928\")"},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/fcupsevB2FytCwiHCYRXXw", Fg: "web.icon=\"0cd46e0cba3abd067cd28e70eb7f2a5f\""},
			},
		},
		"nps": {
			Name: "nps",
			Fingers: []Finger{
				{Type: "fofa", Link: "", Fg: "body=\"nps\" && category=\"黑客工具\""},
			},
		},
		"绿盟": {
			Name: "绿盟",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202943&highlight=%E7%BB%BF%E7%9B%9F", Fg: "body=\"'/needUsbkey.php?username='\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202782&highlight=%E7%BB%BF%E7%9B%9F", Fg: "title=\"NSFOCUS&nbsp;SAS[H]\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201157&highlight=%E7%BB%BF%E7%9B%9F", Fg: "banner=\"PHPSESSID_NF\" || header=\"PHPSESSID_NF\" && body=\"防火墙\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=115769&highlight=%E7%BB%BF%E7%9B%9F", Fg: "body=\"WebApi/encrypt/js-sha1/build/sha1.min.js\""},
			},
		},
		"百傲瑞达": {
			Name: "百傲瑞达",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=115717&highlight=%E7%99%BE%E5%82%B2%E7%91%9E%E8%BE%BE", Fg: "product=\"ZKTECO-百傲瑞达安防管理系统平台\" || icon_hash=\"-1169502834\""},
			},
		},
		"用友智石开PLM": {
			Name: "用友智石开PLM",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/MqmWThb5GuKQeo1s1_J7qg", Fg: "body=\"智石开PLM\""},
			},
		},
		"锐捷 RG-EW1200G": {
			Name: "锐捷 RG-EW1200G",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/9ljtwEVj_Hqog5EkaKFqJQ", Fg: "body=\"static/css/app.2fe6356cdd1ddd0eb8d6317d1a48d379.css\""},
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/X-mRACpw8VAL_0pTORCkiQ", Fg: "body=\"c33367701511b4f6020ec61ded352059\""},
			},
		},
		"东华医疗协同办公系统": {
			Name: "东华医疗协同办公系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=202673&highlight=%E4%B8%9C%E5%8D%8E", Fg: "app=\"DHC-OA\"||product=\"DHC-OA\""},
			},
		},
		"LVS精益管理系统": {
			Name: "LVS精益管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/_Cy6fSXgejKwZBBJ29d81g", Fg: "body=\"/ajax/LVS.Core.Common.STSResult,LVS.Core.Common.ashx\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/_Cy6fSXgejKwZBBJ29d81g", Fg: "body=\"/ajax/LVS.Core.Common.STSResult,LVS.Core.Common.ashx\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/_Cy6fSXgejKwZBBJ29d81g", Fg: "body=\"/ajax/LVS.Core.Common.STSResult,LVS.Core.Common.ashx\""},
			},
		},
		"蓝海卓越计费管理系统": {
			Name: "蓝海卓越计费管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/LYb0S71WJp_0USkXL_mBfw", Fg: "body=\"http://www.natshell.com\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/LYb0S71WJp_0USkXL_mBfw", Fg: "body=\"http://www.natshell.com\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/LYb0S71WJp_0USkXL_mBfw", Fg: "body=\"http://www.natshell.com\""},
			},
		},
		"智慧校园(安校易)管理系统": {
			Name: "智慧校园(安校易)管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/JoDe00HK5XgNOVogDnTLyg", Fg: "body=\"window.external.CallCSharpMethod\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/JoDe00HK5XgNOVogDnTLyg", Fg: "body=\"window.external.CallCSharpMethod\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/JoDe00HK5XgNOVogDnTLyg", Fg: "body=\"window.external.CallCSharpMethod\""},
			},
		},
		"铭飞CMS": {
			Name: "铭飞CMS",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/YtnxlPQGXUOu9Yv9UxmRsQ", Fg: "body=\"static/plugins/ms/1.0.0/ms.js\"||product=\"铭飞/MCMS\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/YtnxlPQGXUOu9Yv9UxmRsQ", Fg: "body=\"static/plugins/ms/1.0.0/ms.js\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/YtnxlPQGXUOu9Yv9UxmRsQ", Fg: "body=\"static/plugins/ms/1.0.0/ms.js\""},
			},
		},
		"宏景HCM": {
			Name: "宏景HCM",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/IRglE9nSomVGzXv7cZCaxg", Fg: "body=\"/images/hcm/themes/default/login/li.png\" || product=\"HJSOFT-HCM\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/IRglE9nSomVGzXv7cZCaxg", Fg: "body=\"/images/hcm/themes/default/login/li.png\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/IRglE9nSomVGzXv7cZCaxg", Fg: "body=\"/images/hcm/themes/default/login/li.png\""},
			},
		},
		"山东聚恒中台": {
			Name: "山东聚恒中台",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/M8w8yZuEn5z9JcieWiTd4Q", Fg: "body=\"sysplat/dataget/data.ashx?type=sendvalidatecode\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/M8w8yZuEn5z9JcieWiTd4Q", Fg: "body=\"sysplat/dataget/data.ashx?type=sendvalidatecode\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/M8w8yZuEn5z9JcieWiTd4Q", Fg: "body=\"sysplat/dataget/data.ashx?type=sendvalidatecode\""},
			},
		},
		"六零导航页": {
			Name: "六零导航页",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/R42AE2fzm3Sg-4_Zop0lIA", Fg: "body=\"六零导航页(LyLme Spage)致力于简洁高效无广告的上网导航和搜索入口\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/R42AE2fzm3Sg-4_Zop0lIA", Fg: "body=\"六零导航页(LyLme Spage)致力于简洁高效无广告的上网导航和搜索入口\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/R42AE2fzm3Sg-4_Zop0lIA", Fg: "body=\"六零导航页(LyLme Spage)致力于简洁高效无广告的上网导航和搜索入口\""},
			},
		},
		"飞鱼星": {
			Name: "飞鱼星",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/7oq0LyoTfJlypcy7NOFGQw", Fg: "body=\"http://www.adslr.com/product/\" || product=\"飞鱼星-路由器\""},
			},
		},
		"MetaCRM": {
			Name: "MetaCRM",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/IozfOEp1Cn5f5F2fDtdNKA", Fg: "body=\"/common/scripts/basic.js\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/IozfOEp1Cn5f5F2fDtdNKA", Fg: "body=\"/common/scripts/basic.js\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/IozfOEp1Cn5f5F2fDtdNKA", Fg: "body=\"/common/scripts/basic.js\""},
			},
		},
		"AJ-Report开源数据大屏": {
			Name: "AJ-Report开源数据大屏",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/mBUUjDjJKPcdCWgGJ2nQAA", Fg: "title=\"AJ-Report\""},
			},
		},
		"万户": {
			Name: "万户",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=201388&highlight=ezEIP", Fg: "app=\"万户网络-CMS\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203388&highlight=%E4%B8%87%E6%88%B7", Fg: "app=\"万户网络-ezOFFICE\""},
			},
		},
		"美特CRM": {
			Name: "美特CRM",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/vDnYCGqsjTL2qdkhn3tgBQ", Fg: "product=\"美特软件-MetaCRM6\""},
			},
		},
		"Rejetto": {
			Name: "Rejetto",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/NgysJFLtf-tGfo6SDzfHxQ", Fg: "app=\"HFS\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/NgysJFLtf-tGfo6SDzfHxQ", Fg: "app.name=\"HTTP File Server\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/NgysJFLtf-tGfo6SDzfHxQ", Fg: "app=\"HFS\""}},
		},
		"PHP-CGI": {
			Name: "PHP-CGI",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/yZeoFAuJF8puEoDcQ_a8bw", Fg: "app=\"XAMPP\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/yZeoFAuJF8puEoDcQ_a8bw", Fg: "app=\"XAMPP\""},
				{Type: "quake", Link: "https://mp.weixin.qq.com/s/yZeoFAuJF8puEoDcQ_a8bw", Fg: "app=\"XAMPP\""}},
		},
		"世邦通信": {
			Name: "世邦通信",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/yZeoFAuJF8puEoDcQ_a8bw", Fg: "icon_hash=\"-1830859634\""},
			},
		},
		"特脸爱云管理系统": {
			Name: "特脸爱云管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/WklRZMaa-mNZtfuCzK93EA", Fg: "title==\"欢迎使用脸爱云 一脸通智慧管理平台\""},
			},
		},
		"Laykefu客服系统": {
			Name: "Laykefu客服系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/xsJHF6lN5z1CudvgnESDUw", Fg: "body=\"static/customer/css/laykefu.css\" || icon_hash=\"-334624619\""},
			},
		},
		"新视窗新一代物业管理系统": {
			Name: "新视窗新一代物业管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/viewthread.php?tid=203524&highlight=%E6%96%B0%E8%A7%86%E7%AA%97", Fg: "body=\"BPMSite/ClientSupport/OCXInstall.aspx\""},
			},
		},
		"YzmCMS": {
			Name: "YzmCMS",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/irBLBaOcX05Ne-NvY5CpFg", Fg: "app=\"yzmcms\""},
			},
		},
		"Mirth Connect": {
			Name: "Mirth Connect",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/ZEN-sS7-Hi9v0IXhBYLrkQ", Fg: "title=\"Mirth Connect Administrator\""},
			},
		},
		"Zyxel NAS": {
			Name: "Zyxel NAS",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/fcTc4JmtMgkwKdS2AFhIsw", Fg: "body=\"/cmd,/ck6fup6/user_grp_cgi/cgi_modify_userinfo\""},
			},
		},
		"悦库企业网盘": {
			Name: "悦库企业网盘",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/QwXWjjJkrEHZCqMtMeQPIg", Fg: "body=\"/user/downloadclient.html\""},
			},
		},
		"契约锁": {
			Name: "契约锁电子签章平台",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/6eUcPQqDxGbvlojw88sZiA", Fg: "app=\"契约锁-电子签署平台\""},
			},
		},
		"极企智能办公路由": {
			Name: "极企智能办公路由",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/RUpJ5sxQClK73rO5UJ4YWQ", Fg: "title=\"极企智能办公路由\""},
			},
		},
		"碧海威科技-L7云路由": {
			Name: "碧海威科技-L7云路由",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/L_nDPKT3BNJ9zsZRk0skLg", Fg: "app=\"碧海威科技-L7云路由\""},
			},
		},
		"思福迪运维安全管理系统": {
			Name: "思福迪运维安全管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://forum.ywhack.com/redirect.php?tid=203508&goto=lastpost#lastpost", Fg: "title=\"Logbase运维安全管理系统\""},
				{Type: "fofa", Link: "https://forum.ywhack.com/redirect.php?tid=203508&goto=lastpost#lastpost", Fg: "app.name=\"Logbase 思福迪 运维安全系统\""},
			},
		},
		"GeoServer": {
			Name: "GeoServer",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/_VyEMatmCfmasCxkmAhHgw", Fg: "app=\"GeoServer\""},
				{Type: "hunter", Link: "https://mp.weixin.qq.com/s/_VyEMatmCfmasCxkmAhHgw", Fg: "app.name=\"GeoServer\""},
			},
		},
		"APP分发签名系统": {
			Name: "APP分发签名系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/aISmcCRE8v9325eMWaGKaw", Fg: "body=\"/api/statistics/service-usage-amount\""},
			},
		},
		"在线录音管理系统": {
			Name: "在线录音管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/YpUbneU-zGQryEdiOLd36g", Fg: "title=\"在线录音管理系统\""},
			},
		},
		"maxView Storage Manager系统": {
			Name: "maxView Storage Manager系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/eRmZgFPD6hmZlpiMKU5zMw", Fg: "title=\"maxView Storage Manager - Login\""},
			},
		},
		"天融信上网行为管理系统": {
			Name: "天融信上网行为管理系统",
			Fingers: []Finger{
				{Type: "fofa", Link: "https://mp.weixin.qq.com/s/N1r3C46t48WXoVZ1qCGbwA", Fg: "body=\"/images/arrow03.png\" && body=\"400-7770777\""},
			},
		},
		"Guns": {
			Name: "Guns",
			Fingers: []Finger{
				{Type: "fofa", Link: "shiro", Fg: "title=\"Guns\" && title=\"登录\""},
			},
		},
	}
}

func Entry(target string, tp string) {
	count := 0
	excelFile := File.CreateExcel()
	File.DeleteSheet(excelFile, "Sheet1")
	//写入domain->证书信息
	File.CreateSheet(excelFile, "FingerToAssets")
	File.WriteExcel(excelFile, "FingerToAssets", "A1", "Name")
	File.WriteExcel(excelFile, "FingerToAssets", "B1", "Reference")
	File.WriteExcel(excelFile, "FingerToAssets", "C1", "Finger")
	File.WriteExcel(excelFile, "FingerToAssets", "D1", "Urls")

	for _, value := range featurMap {
		for i, _ := range value.Fingers {
			if value.Fingers[i].Type == "fofa" {
				if tp == "City" {
					value.Fingers[i].urls = fofa.GetUrlsByCity("(" + value.Fingers[i].Fg + ") && city=\"" + target + "\"")
				} else if tp == "Region" {
					value.Fingers[i].urls = fofa.GetUrlsByCity("(" + value.Fingers[i].Fg + ") && region=\"" + target + "\"")
				}
			}
		}
	}

	for _, value := range featurMap {
		for _, finger := range value.Fingers {
			for _, url := range finger.urls {
				File.WriteExcel(excelFile, "FingerToAssets", fmt.Sprintf("A%d", count+2), value.Name)
				File.WriteExcel(excelFile, "FingerToAssets", fmt.Sprintf("B%d", count+2), finger.Link)
				File.WriteExcel(excelFile, "FingerToAssets", fmt.Sprintf("C%d", count+2), finger.Fg)
				File.WriteExcel(excelFile, "FingerToAssets", fmt.Sprintf("D%d", count+2), url)
				fmt.Println(url)
				count++
			}
		}
	}
	dir := File.GetCurrentAbPath()
	File.SaveExcel(excelFile, dir+"/"+target+"BugAssets.xlsx")
	log.SetPrefix("[+] ")
	log.Println("BugAssets 已写入:" + dir + "/" + target + "BugAssets.xlsx")

}
