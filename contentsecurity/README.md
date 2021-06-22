> 阿里云内容安全golang-sdk


### 现有接口

- [文本审核](https://help.aliyun.com/document_detail/70439.html?spm=a2c4g.11186623.6.701.24554850Tz6viw)
- 图片审核
    - [同步检测](https://help.aliyun.com/document_detail/70292.html?spm=a2c4g.11186623.6.628.6f3c3860Jp6Yxn)
    - [异步检测](https://help.aliyun.com/document_detail/70430.html?spm=a2c4g.11186623.6.629.21114cacMpmTvK)
    - [异步检测结果查询](https://help.aliyun.com/document_detail/70430.html?spm=a2c4g.11186623.6.629.105f4cac4rKyID#title-4tb-bxu-pxg)
- 视频审核
    - [同步检测](https://help.aliyun.com/document_detail/87391.html?spm=a2c4g.11186623.6.688.45503698sWHa5N)
    - [异步检测](https://help.aliyun.com/document_detail/70436.html?spm=a2c4g.11186623.6.689.62804cacEtu1Vp)
    - [异步检测结果查询](https://help.aliyun.com/document_detail/70436.html?spm=a2c4g.11186623.6.689.41acaba5vQq6rf#title-4w9-nwq-fyn)


### 安装

```bash
go get github.com/wksw/aliyun-golang-sdk/contentsecurity
```

### 使用（文本内容检测）

```golang
    import "github.com/wksw/aliyun-golang-sdk/contentsecurity"

    func main() {
        client := contentsecurity.New("endpoint", "ak", "sk")
        resp, err := client.ScanText(&ScanTextReq{
            ScanCommonReq: ScanCommonReq{
                Scenes: []string{"antispam"},
            },
            Tasks: []ScanTextTask{
                {
                    ClientInfo: ClientInfo{
                        UserId:   "abc",
                        UserType: UserOther,
                    },
                    DataId:  "123",
                    Content: "本校小额贷款，安全、快捷、方便、无抵押，随机随贷，当天放款，上门服务。",
                },
            },
        })
        if err != nil {
            log.Fatal(err.Error())
        }
        // 根据业务逻辑自行处理返回逻辑
    }
```