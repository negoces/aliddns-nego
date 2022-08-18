# aliddns-nego

1. 编译：`go build && strip aliddns-nego`
2. 设置相关环境变量

|变量名|描述|
|:-|:-|
|CONN_TEST_API| generate_204 地址 |
|ENABLE_IPV4| IPv4 功能开关 |
|ENABLE_IPV6| IPv4 功能开关 |
|ACCESS_ID| 访问 Token ID |
|ACCESS_SECRET| 访问 Token 密钥 |
|MYIPV4_API|获取当前IPv4的API|
|MYIPV6_API|获取当前IPv6的API|
|DOMAIN_V4|IPv6主域名|
|SUB_DOMAIN_V4|IPv6记录名|
|DOMAIN_V6|IPv6主域名|
|SUB_DOMAIN_V6|IPv6记录名|

3. 运行 `aliddns-nego`

## 设置定时任务

```ini
# /etc/systemd/system/aliddns.service
[Unit]
Description=AliDDNS (Author: negoces)

[Service]
# Environment=xxx=xxx
EnvironmentFile=/opt/aliddns/config
ExecStart=/opt/aliddns/aliddns-nego
```

```ini
# /etc/systemd/system/aliddns.timer
[Unit]
Description=Timer - AliDDNS (Author: negoces)

[Timer]
OnCalendar=*-*-* *:00/10:00
Unit=aliddns.service

[Install]
WantedBy=timers.target
```

- 测试：`sudo systemctl start aliddns`
- 启用：`sudo systemctl enable --now aliddns.timer`
