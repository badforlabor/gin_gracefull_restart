### 注意
- 必须在linux上才支持热重启。
- 使用的是[endless](https://github.com/fvbock/endless)。也可使用facebook的[grace](https://github.com/facebookgo/grace)

### 验证流程
- 首先编译bin：go build -o gracefull_restart
- 然后运行 ./gracefull_restart
- 测试是否启动成功： curl http://127.0.0.1:4242/hello
  - 控制台中显示了： main world
- （重点来了）重新编译bin，这次用测试代码：go build -o gracefull_restart simple1/simple_main.go
- ps aux |grep gracefull_restart
- 假定进程id是：28051
- 执行 kill -1 28051
- 再次测试：curl http://127.0.0.1:4242/hello 
  - 这次，控制台中显示：simple world