# ufo-adapter
 继承 beam-adapter
 
 https://github.com/Assetsadapter/beam-adapter
 
ufo暂时无法适配openwallet钱包体系，无法实现离线生产地址，离线交易签名。
因为ufo的Mimblewimble协议需要钱包一直保持在线，才能正常地接收和发送资产。
为了实现企业级别钱包解决方案，我们划分两个应用场景。
- 用户托管钱包场景。就是企业托管了注册用户的钱包，为每一个用户分配接收地址，这个钱包不对外开放。在一台装有beam钱包的服务器上，运行openw-beam钱包服务，监听用户的充值，并定时进行汇总。
- 财务系统提币钱包场景。企业不会直接从用户托管钱包中做转账，而是在一个独立的热钱包中驻留日常业务的需要的资产。财务系统集成beam-adapter，可发送指令向openw-beam钱包服务创建新地址，可查询用户托管钱包的新充值记录。

## 官方资料

### 官网

https://ufo.link

### 接口文档

#### Wallet API

https://github.com/BeamMW/beam/wiki/Beam-wallet-protocol-API

#### Explorer API

https://github.com/BeamMW/beam/wiki/Beam-Node-Explorer-API

### 浏览器

https://explorer.ufo.link



