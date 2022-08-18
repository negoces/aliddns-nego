package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type IPInfo struct {
	IP string `json:"ip"`
}

type Record struct {
	IP string
	ID string
}

var (
	Test_URL        string = "https://connect.rom.miui.com/generate_204"
	IPv4_API_URL    string = "https://api.ipify.org?format=json"
	IPv6_API_URL    string = "https://api6.ipify.org?format=json"
	accessKeyId     string
	accessKeySecret string
	v6_disable_flag string = "0"
	IPv4            string
	IPv6            string
	domain          string
	domain_v6       string
	subDomain       string
	subDomain_v6    string
	v4_record       Record
	v6_record       Record
)

func main() {
	log.Println("正在从环境变量获取相关设置")
	err := getEnv()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("获取完成")
	log.Println("将查询与修改的 IPv4 域名为:", subDomain+"."+domain)
	if v6_disable_flag == "0" {
		log.Println("将查询与修改的 IPv6 域名为:", subDomain_v6+"."+domain_v6)
	}
	log.Println("将使用", IPv4_API_URL, "获取 IPv4")
	if v6_disable_flag == "0" {
		log.Println("将使用", IPv6_API_URL, "获取 IPv6")
	}
	log.Println("正在检查网络连接")
	log.Println("正在请求", Test_URL)
	err = verifyNetConn()
	if err != nil {
		log.Println("网络检查失败:", err.Error())
		os.Exit(1)
	}
	log.Println("网络连接正常")
	log.Println("正在获取IP")
	err = getIP(IPv4_API_URL, &IPv4)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("IPv4:", IPv4)
	if v6_disable_flag == "0" {
		err = getIP(IPv6_API_URL, &IPv6)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Println("IPv6:", IPv6)
	}
	log.Println("正在查询记录")
	err = queryRecord(subDomain+"."+domain, "A", &v4_record)
	if err != nil {
		log.Println("查询失败:", err.Error())
		os.Exit(1)
	}
	if v6_disable_flag == "0" {
		err = queryRecord(subDomain_v6+"."+domain_v6, "AAAA", &v6_record)
		if err != nil {
			log.Println("查询失败: ", err.Error())
			os.Exit(1)
		}
	}
	if v4_record.ID == "" {
		log.Println("未查询到相应 A 记录，添加记录:", IPv4)
		err = addRecord(domain, subDomain, "A", IPv4)
		if err != nil {
			log.Println("添加失败: ", err.Error())
			os.Exit(1)
		}
		log.Println("A 记录添加成功")
	} else if v4_record.IP != IPv4 {
		log.Println("A 记录与当前 IP 不匹配，更新记录:", IPv4)
		err = updateRecord(v4_record.ID, subDomain, "A", IPv4)
		if err != nil {
			log.Println("更新失败: ", err.Error())
			os.Exit(1)
		}
		log.Println("A 记录更新成功")
	} else {
		log.Println("当前 A 记录:", v4_record.IP, "无需更新")
	}
	if v6_disable_flag == "0" {
		if v6_record.ID == "" {
			log.Println("未查询到相应 AAAA 记录，添加记录:", IPv6)
			err = addRecord(domain_v6, subDomain_v6, "AAAA", IPv6)
			if err != nil {
				log.Println("添加失败: ", err.Error())
				os.Exit(1)
			}
			log.Println("AAAA 记录添加成功")
		} else if v6_record.IP != IPv6 {
			log.Println("AAAA 记录与当前 IP 不匹配，更新记录:", IPv6)
			err = updateRecord(v6_record.ID, subDomain_v6, "AAAA", IPv6)
			if err != nil {
				log.Println("更新失败: ", err.Error())
				os.Exit(1)
			}
			log.Println("AAAA 记录更新成功")
		} else {
			log.Println("当前 AAAA 记录:", v6_record.IP, "无需更新")
		}
	}
}

func getEnv() error {
	accessKeyId = os.Getenv("ACCESS_ID")
	if accessKeyId == "" {
		return errors.New("无法获取 AccessID, 请设置 ACCESS_ID 环境变量")
	}
	accessKeySecret = os.Getenv("ACCESS_SECRET")
	if accessKeySecret == "" {
		return errors.New("无法获取 AccessSecret, 请设置 ACCESS_SECRET 环境变量")
	}
	buff_v6_disable_flag := os.Getenv("DISABLE_V6")
	if buff_v6_disable_flag != "" {
		v6_disable_flag = buff_v6_disable_flag
	}
	domain = os.Getenv("DOMAIN")
	if domain == "" {
		return errors.New("无法获取 Domain, 请设置 DOMAIN 环境变量")
	}
	subDomain = os.Getenv("SUB_DOMAIN")
	if subDomain == "" {
		return errors.New("无法获取 SubDomain, 请设置 SUB_DOMAIN 环境变量")
	}
	if v6_disable_flag == "0" {
		domain_v6 = os.Getenv("DOMAIN_V6")
		if domain_v6 == "" {
			return errors.New("无法获取 Domain v6, 请设置 DOMAIN_V6 环境变量")
		}
		subDomain_v6 = os.Getenv("SUB_DOMAIN_V6")
		if subDomain_v6 == "" {
			return errors.New("无法获取 SubDomain v6, 请设置 ACCESS_SECRET 环境变量")
		}
	}
	buff_v4api := os.Getenv("IPv4_API_URL")
	if buff_v4api != "" {
		IPv4_API_URL = buff_v4api
	}
	buff_v6api := os.Getenv("IPv6_API_URL")
	if buff_v4api != "" {
		IPv6_API_URL = buff_v6api
	}
	return nil
}

func verifyNetConn() error {
	_, err := http.Get(Test_URL)
	if err != nil {
		return err
	}
	return nil
}

func getIP(api_url string, var_ip *string) error {
	res, err := http.Get(api_url)
	if err != nil {
		defer res.Body.Close()
		return err
	}
	ip, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	var resp IPInfo
	err = json.Unmarshal(ip, &resp)
	if err != nil {
		return err
	}
	*var_ip = resp.IP
	return nil
}
