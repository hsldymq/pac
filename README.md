# PAC

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

#### 使用
在浏览器设置或其他代理工具中指定PAC文件地址为 http://localhost:10801/pac.js

#### 环境变量
LISTEN_PORT: pac服务监听端口, 默认:`10801`

在docker run中使用 -e 或者在docker-compose中使用environment修改监听端口

#### 容器目录
容器中pac文件路径强制为/var/www/pac/pac.js, 所以无论宿主机环境中PAC文件是如何命名的