## Kafka 설정

| Kafka | Broker ID |IP|
|:-------:|:---:|:---:|
|kafka1| 1 |172.17.0.5|
|kafka2| 2 |172.17.0.6|
|kafka3| 3 |172.17.0.7|

### 컨테이너 개별 설정

* 디렉토리 생성

    ```
    mkdir -p /data1 /data2 /data3
    ```

### 컨테이너 개별 설정    

```
vi /kafka/config/server.properties
```

* kafka1

    ```
    broker.id=1
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    listeners=PLAINTEXT://0.0.0.0:9092
    advertised.listeners=PLAINTEXT://kafka1:9092
    ..(생략)
    ```
    
* kafka2

    ```
    broker.id=2
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    listeners=PLAINTEXT://0.0.0.0:9092
    advertised.listeners=PLAINTEXT://kafka2:9092
    ..(생략)
    ```

* kafka3

    ```
    broker.id=3
    log.dirs=/data1,/data2,/data3
    zookeeper.connect=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    listeners=PLAINTEXT://0.0.0.0:9092
    advertised.listeners=PLAINTEXT://kafka3:9092    
    ..(생략)
    ```

* kafka 시작

    ```console
    $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
    ```

