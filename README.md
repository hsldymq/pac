# PAC

**pac.js中默认的代理地址是SOCKS 127.0.0.1:1080, 请根据使用者自己情况修改**

## 使用方式
### 1. 直接使用
将pac.js文件应用到浏览器代理设置,或其他代理工具中指定pac.js文件系统路径

**这个仓库中的pac.js文件是作者自用,有一些作者访问比较慢但是没有被墙的网站也被加入到列表中,使用者请自行fork根据自己的需要覆盖/修改.**

### 2. docker
由于chrome的沙盒环境不再允许在PAC proxy配置中指定文件系统中的PAC文件,而是通过http服务来获取.

所以特意制作这个docker镜像,提供一个pac http服务.

firefox以后会不会这么搞也很难说.

#### docker pull
```bash
docker pull hsldymq/pac
```

#### docker run
```bash
docker run --name pac --restart always --log-driver journald -p 10801:10801 -v path/to/your/pac/file:/var/www/pac/pac.js -d hsldymq/pac
````

#### docker-compose
```docker-compose
version: "3"
services:
    pac:
        image: hsldymq/pac
        container_name: pac
        restart: always
        logging:
            driver: journald 
        volumes:
            - /path/to/your/pac/file:/var/www/pac/pac.js
        ports:
            - 10801:10801
```

在docker目录下,有一个已写好的docker-compose.yml文件, 不想自己写的话可以进入docker目录运行`docker-compose up -d`

#### 使用
在浏览器设置或其他代理工具中指定PAC文件地址为 http://localhost:10801/pac.js

#### 容器端口
容器内部监听10801端口,宿主机的监听端口你可以根据自己的需要进行映射

#### 容器中pac.js路径
容器中pac文件路径强制为/var/www/pac/pac.js, 所以无论宿主机环境中PAC文件是如何命名,都必须映射到容器中的/var/www/pac/pac.js
