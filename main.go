package main

import "os"

func main() {
	getEnv()
	println("===> 检查网络连接")
	ms, err := verifyNetConn(connTestAPI)
	if err != nil {
		println("网络连接失败", err.Error())
		os.Exit(1)
	}
	println("网络连接成功，延迟(ms):", ms)
	if flag.v4 == "true" {
		println("===> IPv4: 查询当前 IP")
		getIP(myip.v4, &v4.currentIP, Ipv4Reg)
		if err != nil {
			println("查询当前 IP 失败:", err.Error())
			os.Exit(1)
		}
		println("当前 IP:", v4.currentIP)
		println("===> IPv4: 查询当前域名记录")
		err = queryRecord(v4.subDomain+"."+v4.domain, "A", &v4.record)
		if err != nil {
			println("查询记录失败:", err.Error())
			os.Exit(1)
		}
		if v4.record.ID == "" {
			println("未查询到相应 A 记录, 即将添加")
			println("===> IPv4: 添加新 A 记录")
			err = addRecord(v4.domain, v4.subDomain, "A", v4.currentIP)
			if err != nil {
				println("添加记录失败:", err.Error())
				os.Exit(1)
			}
			println("添加记录成功:", "A", v4.currentIP)
		} else if v4.record.IP != v4.currentIP {
			println("当前记录:", v4.record.IP, ",与当前 IP 不一致, 即将更新")
			println("===> IPv4: 更新 A 记录")
			err = updateRecord(v4.record.ID, v4.subDomain, "A", v4.currentIP)
			if err != nil {
				println("更新记录失败:", err.Error())
				os.Exit(1)
			}
			println("更新记录成功:", "A", v4.currentIP)
		} else {
			println("当前记录:", v4.record.IP, ",与当前 IP 一致, 无需更新")
		}
	}
	if flag.v6 == "true" {
		println("===> IPv6: 查询当前 IP")
		getIP(myip.v6, &v6.currentIP, Ipv6Reg)
		if err != nil {
			println("查询当前 IP 失败:", err.Error())
			os.Exit(1)
		}
		println("当前 IP:", v6.currentIP)
		println("===> IPv6: 查询当前域名记录")
		err = queryRecord(v6.subDomain+"."+v6.domain, "AAAA", &v6.record)
		if err != nil {
			println("查询记录失败:", err.Error())
			os.Exit(1)
		}
		if v6.record.ID == "" {
			println("未查询到相应 AAAA 记录, 即将添加")
			println("===> IPv6: 添加新 AAAA 记录")
			err = addRecord(v6.domain, v6.subDomain, "AAAA", v6.currentIP)
			if err != nil {
				println("添加记录失败:", err.Error())
				os.Exit(1)
			}
			println("添加记录成功:", "AAAA", v6.currentIP)
		} else if v6.record.IP != v6.currentIP {
			println("当前记录:", v6.record.IP, ",与当前 IP 不一致, 即将更新")
			println("===> IPv6: 更新 AAAA 记录")
			err = updateRecord(v6.record.ID, v6.subDomain, "AAAA", v6.currentIP)
			if err != nil {
				println("更新记录失败:", err.Error())
				os.Exit(1)
			}
			println("更新记录成功:", "AAAA", v6.currentIP)
		} else {
			println("当前记录:", v6.record.IP, ",与当前 IP 一致, 无需更新")
		}
	}
}
