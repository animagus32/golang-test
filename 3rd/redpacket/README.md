# 实现一个简易的口令红包功能。

## 功能：

1.  用户可以发"口令红包"。发的红包不从余额内扣除。（假设用户都是通过银行卡发的红包）
2.  用户可以抢"口令红包"。每个用户对于每一个红包仅能抢一次。抢到的红包计算在余额内。
3.  "口令红包"可以指定发多少钱，以及红包的个数。每个红包的金额则是随机的。只要知道口令，无论谁都可以抢。
4.  口令由系统自动生成，仅包含大小写字母和数字，8位
5.  未抢完的红包在（发红包的）24小时后自动退回。
6.  用户可以查看自己的余额。
7.  用户可以查看自己抢到的红包列表。

## 完成功能
目前功能基本都已完成，为了方便测试红包收回设为30秒
不足和改进：
由于时间关系，代码中未加配置文件、日志、异常处理等功能；对于抢红包等方面可以加入锁以应对高并发；红包算法可以优化
## 接口
1. 注册 POST: api/v1/user/register  
    参数：
     * username 用户名
     * password 密码(最好md5加密)
     
     ```  
     返回值：
     {
           "code": 0,
           "msg": "",
           "data": {
               "id": 3,
               "username": "test7",
               "balance": 0
           }
     }
     ```
     
2. 登录 POST: api/v1/user/login  
    参数：
     * username 用户名
     * password 密码
     ```
     {
         "code": 0,
         "msg": "",
         "token": "U47zCGFTJm0LELBoyFJwoQ==",
         "data": {
             "id": 1,
             "username": "test9",
             "balance": 0
         }
     }

    ```

3. 获取余额 GET: api/v1/user/1/balance
   参数:
    * header: 需要传 uid 和 token 作为验证登录
     ```angular2html
    {
        "code": 0,
        "msg": "",
        "Data": {
            "balance": 32.36
        }
    }
    ```
4. 创建红包 POST: api/v1/redpacket/dispatch
    参数：
    * uid 用户id
    * amount 金额
    * num 红包数
   ```
   {
       "code": 0,
       "msg": "",
       "Data": {
           "list": [
               {
                   "id": 292,
                   "redpacket_id": 38,
                   "user_id": 0,
                   "amount": 0.87
               },
               {
                   "id": 293,
                   "redpacket_id": 38,
                   "user_id": 0,
                   "amount": 1
               },
               {
                   "id": 294,
                   "redpacket_id": 38,
                   "user_id": 0,
                   "amount": 8.13
               }
           ],
           "secret": "aoMATro6"
       }
   }
   ```
   
5. 抓红包  POST: api/v1/redpacket/grab
    参数：
     * secret 红包密码
     
  ```
{
    "code": 0,
    "msg": "",
    "data": {
        "id": 295,
        "redpacket_id": 39,
        "user_id": 1,
        "amount": 1.6
    }
}
```

6. 获取红包  GET: api/v1/redpacket/list?uid=1
   参数：
   * uid
   
```angular2html
{
    "code": 0,
    "msg": "",
    "data": [
        {
            "id": 1,
            "redpacket_id": 1,
            "user_id": 1,
            "amount": 12.14
        },
        {
            "id": 295,
            "redpacket_id": 39,
            "user_id": 1,
            "amount": 1.6
        }
    ]
}
```


## 测试
因时间有限，目前完成几个主流程测试
```angular2html
=== RUN   TestLogin
--- PASS: TestLogin (0.00s)
=== RUN   TestLoginWrongPasswd
--- PASS: TestLoginWrongPasswd (0.00s)
=== RUN   TestDispatch
--- PASS: TestDispatch (0.01s)
=== RUN   TestGrab
--- PASS: TestGrab (0.01s)
=== RUN   TestGrabFail
--- PASS: TestGrabFail (0.00s)
PASS
ok      redpacket/test  0.046s

```