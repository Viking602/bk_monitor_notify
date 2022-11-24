# 蓝鲸监控告警回调通知

## 1. 介绍

本项目用于蓝鲸告警配置回调通知

开发框架:[go-zero](https://go-zero.dev/cn/)

当前支持： 钉钉机器人 、飞书机器人

## 2. 配置方法

### 飞书机器人:

飞书地址:

http://127.0.0.1:8887/api/notify/larkBot?key={飞书机器人key}&secret={签名secret}

添加机器人:

添加机器人后选择签名校验

飞书机器人key为webhook最后的一串uuid
<img alt="img.png" height="60%" src="image/img.png" width="70%"/>

飞书机器人secret为签名校验的secret

<img alt="img_1.png" height="60%" src="image/img_1.png" width="60%"/>

注意：默认会在告警时间加8小时如不需要添加则添加参数
&timeLocal=0

完整链接示意:

默认转换时区:

http://127.0.0.1:8887/api/notify/larkBot?key=06ddbd05-bc87-4c28-8230-6d0406ac3c67&secret=jbfG9PaiYHj02GOp

不转换时区:

http://127.0.0.1:8887/api/notify/larkBot?key=06ddbd05-bc87-4c28-8230-6d0406ac3c67&secret=jbfG9PaiYHj02GOp&timeLocal=0

飞书机器人效果:

<img alt="img_9.png" src="image/img_9.png" width="60%" height="60%"/>
<img alt="img_5.png" height="50%" src="image/img_5.png" width="60%"/>
<img alt="img_10.png" height="50%" src="image/img_10.png" width="60%"/>
<img alt="img_11.png" height="50%" src="image/img_12.png" width="60%"/>

### 钉钉机器人:

添加机器人:

获取机器人webhook地址拿到最后的token
链接示意https://oapi.dingtalk.com/robot/send?access_token=5165f0fb126694c6a15770c8029*********************f129e22
<img alt="img_2.png" height="60%" src="image/img_2.png" width="60%"/>
安全设置选择加签验证
<img alt="img_3.png" height="60%" src="image/img_3.png" width="60%"/>

钉钉地址:

http://127.0.0.1:8887/api/notify/dingTalkBot?key={钉钉机器人Token}&secret={钉钉签名secret}

钉钉机器人Token取webhook地址最后的access_token=后面的一串字符

完整链接示意:

http://127.0.0.1:8887/api/notify/dingTalkBot?key=asdadfqwe123as123&secret=test1234567654asd

转换时区参考上述飞书链接

钉钉展示效果:

<img alt="img_7.png" height="60%" src="image/img_7.png" width="40%"/>
<img alt="img_8.png" height="60%" src="image/img_8.png" width="40%"/>
<img alt="img_4.png" height="60%" src="image/img_4.png" width="40%"/>
<img alt="img_6.png" height="60%" src="image/img_6.png" width="40%"/>

## 3. config文件

```yaml
Name: notify
Host: 0.0.0.0
Port: 8887 # 启动端口
BkUrl:  # 平台链接用于展示
```

## 4. 蓝鲸配置方法

<img alt="screenshot-20220927-114456.png" src="image/screenshot-20220927-114456.png"/>

## 5. 启动方法

```bash
./notify -f etc/notify.yaml
```

## 6. Docker启动
    
```bash
docker run -d -p 8887:8887 --name bk-notify -e BK_URL=http://paas.xxx.com viking602/bk-notify:latest
```