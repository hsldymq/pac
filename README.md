# PAC

**pac.js中默认的代理地址是SOCKS 127.0.0.1:1080, 你根据自己情况修改**

## 直接使用
将pac.js文件应用到浏览器代理设置,或其他代理工具中指定pac.js文件系统路径

## docker
由于chrome的沙盒环境不再允许在PAC proxy配置中指定文件系统中的PAC文件,firefox以后会不会这么搞也很难说.所以特意制作这个docker镜像

#### docker pull
```bash
docker pull hsldymq/pac
```

#### docker run
```bash
docker run --name pac --log-driver journald -p 10801:10801 -v path/to/your/pac/file:/var/www/pac/pac.js -d hsldymq/pac
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

**不建议使用这个仓库中的pac文件**

#### 容器端口
容器内部监听10801端口,宿主机的监听端口你可以根据自己的需要进行映射

#### 容器中pac.js路径
容器中pac文件路径强制为/var/www/pac/pac.js, 所以无论宿主机环境中PAC文件是如何命名的