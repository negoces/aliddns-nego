package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func readEnvWithDef(env string, def string) string {
	buff := os.Getenv(env)
	if buff != "" {
		return buff
	} else {
		return def
	}
}

func MustReadEnv(env string) (string, error) {
	buff := os.Getenv(env)
	if buff != "" {
		return buff, nil
	} else {
		return "", errors.New("请设置环境变量: " + env)
	}
}

func getEnv() {
	println("===> 读取环境变量")
	err := godotenv.Load()
	if err != nil {
		println("获取 .env 失败，跳过")
	} else {
		println("获取 .env 成功")
	}
	connTestAPI = readEnvWithDef("CONN_TEST_API", "https://connect.rom.miui.com/generate_204")
	flag.v4 = readEnvWithDef("ENABLE_IPV4", "true")
	flag.v6 = readEnvWithDef("ENABLE_IPV6", "true")
	myip.v4 = readEnvWithDef("MYIPV4_API", "https://myip.ipip.net")
	myip.v6 = readEnvWithDef("MYIPV6_API", "https://myip6.ipip.net")
	aliApiToken.accessKeyId, err = MustReadEnv("ACCESS_ID")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	aliApiToken.accessKeySecret, err = MustReadEnv("ACCESS_SECRET")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if flag.v4 == "true" {
		v4.domain, err = MustReadEnv("DOMAIN_V4")
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		v4.subDomain, err = MustReadEnv("SUB_DOMAIN_V4")
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	if flag.v6 == "true" {
		v6.domain, err = MustReadEnv("DOMAIN_V6")
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		v6.subDomain, err = MustReadEnv("SUB_DOMAIN_V6")
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
	println("读取完毕")
	println("网络测试 API:", connTestAPI)
	println("是否启用 IPv4:", flag.v4)
	println("是否启用 IPv6:", flag.v6)
	println("IPv4 获取 API:", myip.v4)
	println("IPv6 获取 API:", myip.v6)
	if flag.v4 == "true" {
		println("将要更新的 IPv4 域名:", v4.subDomain+"."+v4.domain)
	}
	if flag.v6 == "true" {
		println("将要更新的 IPv6 域名:", v6.subDomain+"."+v6.domain)
	}
}
