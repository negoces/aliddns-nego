# aliddns-nego

1. 编译：`go build && strip aliddns-nego`
2. 设置相关环境变量

|变量名|描述|
|:-|:-|
|ACCESS_ID| 访问 Token ID |
|ACCESS_SECRET| 访问 Token 密钥 |
|DOMAIN|主域名|
|SUB_DOMAIN|记录名|
|DISABLE_V6|禁用 IPv6 `0`或`1` |
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
