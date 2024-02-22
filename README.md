# tiktok-shop-auth-service

> ⚠️ 注意，该服务需部署后，需确保 TikTok 能访问到。

设置环境变量。

```shell
>> export TTS_APP_KEY=xxxx TTS_APP_SECRET=xxxx
```

启动服务。

```shell
>> chmod +x /path/to/tts-auth-service
# 启动服务
>> /path/to/tts-auth-service

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

在 TikTok Partner Center 对商家请求授权，观察该服务日志，若输出 AccessToken 则符合预期。
