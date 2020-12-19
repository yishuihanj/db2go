# pgtogo
一个使 pgsql 数据库表自动生成 go struct 的工具

### 快速使用

进入项目根目录下进行

```shell
go build
```

对于不同的操作系统

`windows`

```powershell
pgtogo.exe -host=127.0.0.1 -port=5432 -user=postgres -pwd=postgres  -gorm=true
```

`linux`

```shell
$ ./pgtogo -host=127.0.0.1 -port=5432 -user=postgres -pwd=postgres  -gorm=true
```

### 命令行提示

执行

```powershell
Usage of pgtogo.exe:
  -dbname string
        数据库名称，必填，否则会报错
  -driver string
        必填，需要连接的数据库，现在只支持mysql、pgsql 例如 -driver=mysql，-driver=pgsql
  -gorm
        是否添加 gorm tag，true添加，false不添加，默认不添加
  -host string
        数据库ip，默认为localhost (default "localhost")
  -outdir string
        .go 文件输出路径，不设置的话会输出到当前程序所在路径 (default "./go_output")
  -port int
        必填，数据库端口
  -pwd string
        必填，数据库密码
  -table string
        需要导出的数据库表名称，如果不设置的话会将该数据库所有的表导出
  -user string
        必填，数据库用户名

```

