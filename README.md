- 本小程序主要用来实现Geneva的以下四条规则，还可以自定义端口、需要修改的window size的值。
```
"[TCP:flags:SA]-tamper{TCP:window:replace:0}-|"
"[TCP:flags:A]-tamper{TCP:window:replace:0}-|"
"[TCP:flags:PA]-tamper{TCP:window:replace:0}-|"
"[TCP:flags:FA]-tamper{TCP:window:replace:0}-|"
```
- 具体用法：https://www.444.run