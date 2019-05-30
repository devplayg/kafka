#### [목록](../README.md)

## Kafka 설정

| Kafka | Broker ID |IP|
|:-------:|:---:|:---:|
|kafka1| 1 |172.17.0.5|
|kafka2| 2 |172.17.0.6|
|kafka3| 3 |172.17.0.7|

### 컨테이너 공통 설정

* 호스트 파일 설정

    ```
    vi /etc/hosts
    ---
    172.17.0.2      zoo1
    172.17.0.3      zoo2
    172.17.0.4      zoo3
    172.17.0.5      kafka1
    172.17.0.6      kafka2
    172.17.0.7      kafka3
    ```

* 디렉토리 생성

    ```
    mkdir -p /data1 /data2 /data3
    ```

### 컨테이너 개별 설정    

    vi /kafka/config/server.properties

* kafka1
    ```
    broker.id=1
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    ..(생략)
    ```
    
* kafka2
    ```
    broker.id=2
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    ..(생략)
    ```

* kafka3

    ```
    broker.id=3
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    ..(생략)
    ```