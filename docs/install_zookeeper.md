#### [목록](../README.md)

## Zookeeper 설정
    
### 컨테이너 공통 설정

* Zookeeper 설정

    ```
    vi /kafka/config/zookeeper.properties
    ---
    dataDir=/data
    tickTime=2000
    initLimit=5
    syncLimit=2
    server.1=zoo1:2888:3888
    server.2=zoo2:2888:3888
    server.3=zoo3:2888:3888
    ```

### 컨테이너 개별 설정

* zoo1

    ```
    mkdir -p /data && echo 1 > /data/myid
    ```
    
* zoo2

    ```
    mkdir -p /data && echo 2 > /data/myid
    ```
    
* zoo3

    ```
    mkdir -p /data && echo 3 > /data/myid
    ```

### Zookeeper 컨테이너 시작

```
/kafka/bin/zookeeper-server-start.sh /kafka/config/zookeeper.properties
```
