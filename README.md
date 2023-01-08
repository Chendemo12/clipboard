# sample

## Usage

### 运行项目

- 项目基于`go 1.19`, 采用`go mod`;

1. 配置环境变量:

- `HTTP_API_HOST`: 默认值:`0.0.0.0`;
- `HTTP_API_PORT`: 本机运行端口, 默认值:`7077`;
- `REMOTE_HTTP_HOST`: 对端地址;
- `REMOTE_HTTP_PORT`: 对端服务端口号;
- `REMOTE_SYNC_INTERVAL`; 同步间隔, 区间`[1,10]`, 默认`2`;

2. 或通过`--env` 来指定环境变量文件;
3. 支持的运行参数:

```bash
❯ go run main.go --help
2023/01/08 12:40:32 maxprocs: Leaving GOMAXPROCS=8: CPU quota undefined
Usage of /tmp/go-build343061439/b001/exe/main:
  -dev
        run with DevMode.
  -env string
        specifying a environment file.
  -v    show version.
  -version
        show version and description.
```

### 打包项目

1. 运行`packer.sh`脚本;

- 打包为`DEBIAN`软件包;