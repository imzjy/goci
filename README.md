# goci

简单的自动化集成工具，支持GitHub和BitBucket

# 安装

编译安装

```shell
$export GOPATH=/path/to/you/gopath
$go get -u github.com/imzjy/goci
$./bin/goci
```

直接下载

- [Linux 64](https://github.com/imzjy/goci/releases)

# 配置Webhooks

在GitHub或者BitBucket的仓库中设置Webhooks，目前只支持push event。

- GitHub指向:    `http://you-host.com:8080/github`
- BitBucket指向: `http://you-host.com:8080/bitbucket`

# 配置文件语法

标准的JOSN格式的文件

```text
{
	"port":"8080",    //监听端口
	"triggers":       //触发设定
		[
			{
				"Issuer" :     "bitbucket",        // 监听bitbucket的push事件
				"Repository" : "imzjy/wschat",     // push的仓库，只有push到这个仓库才会触发命令
				"Branch" :     "master",           // push的分支，只有push到这个分支才会触发命令
				"Type" :       "local",            // 执行本地命令
				"Cmd" :        "ps aux"            // push事件通知是需要执行的命令
			},
			{
				"Issuer" :     "github",
				"Repository" : "imzjy/s3cl",
				"Branch" :     "online",
				"Type" :       "ssh",              // 执行ssh远程命令
				"SshUser" :    "zjy",              // 远程用户机器名
				"SshHost" :    "www.imzjy.com",    // 远程主机名
				"SshKey" :     "~/.ssh/id_rsa",    // 登陆远程主机用的私钥
				"Cmd" :        "ps aux"            // 执行的远程命令
			}
		]
}
```