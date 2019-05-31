## 테스트

docker run -v e:/gohome:/gohome -it --name test ubuntu bash

sed -i s#archive\.ubuntu\.com#mirror.kakao.com#g /etc/apt/sources.list
apt update && apt install -y curl git vim
curl https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz | tar xvfz - -C /
vi ~/.profile
------------------------------------------------------------------
export GOROOT=/go
export GOPATH=/gohome
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
------------------------------------------------------------------ 

### KAFKA Start/Stop Test

| Seq | kafka1 | kafka2 | kafka3 |
|-----|--------|--------|--------|
| 1   | on     | on     | on     |
| 2   | on     | on     | off    |
| 3   | on     | off    | off    |
| 4   | on     | on     | off    |
| 5   | on     | on     | on     |

* kafka1-On, kafka2-On, kafka3-On

    ```
    Topic: my-topic Partition: 0    Leader: 3       Replicas: 2,1,3 Isr: 3,2,1    
    Topic: my-topic Partition: 1    Leader: 3       Replicas: 3,2,1 Isr: 3,2,1    
    Topic: my-topic Partition: 2    Leader: 3       Replicas: 1,3,2 Isr: 3,2,1    
    Topic: my-topic Partition: 3    Leader: 3       Replicas: 2,3,1 Isr: 3,2,1    
    Topic: my-topic Partition: 4    Leader: 3       Replicas: 3,1,2 Isr: 3,2,1    
    ```

* kafka1-On, kafka2-On, kafka3-Off

    ```
    Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3 Isr: 2,1      
    Topic: my-topic Partition: 1    Leader: 2       Replicas: 3,2,1 Isr: 2,1      
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2 Isr: 2,1      
    Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1 Isr: 2,1      
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2 Isr: 2,1      
    ```

* kafka1-On, kafka2-Off, kafka3-Off

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3 Isr: 1   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1 Isr: 1   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2 Isr: 1   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1 Isr: 1   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2 Isr: 1   
    ```

* kafka1-On, kafka2-On, kafka3-Off

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3 Isr: 1,2   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1 Isr: 1,2   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2 Isr: 1,2   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1 Isr: 1,2   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2 Isr: 1,2   
    ```

* kafka1-On, kafka2-On, kafka3-On

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3 Isr: 1,2,3   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1 Isr: 1,2,3   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2 Isr: 1,2,3   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1 Isr: 1,2,3   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2 Isr: 1,2,3   
    ```

docker run --name test -it ubuntu bash

### Topics

* Event (name=`event`)
    - date, src_ip, dst_ip, category, filename,  url 
    
* Traffic (name=`traffic`)
    - date, in_bps, out_bps

* Resource (name=`resource`)
    - date, cpu usazge, memory usage, disk usage

### Topic 생성

```shell
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper $ZOOKEEPER --replication-factor 3 --partitions 5 --topic my-topic --create
```

### Topic 목록 조회

```shell
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper $ZOOKEEPER -list
---
my-topic
```

### Topic 상세 조회

```
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper $ZOOKEEPER --describe --topic my-topic 
---
Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3 Isr: 2,1,3
Topic: my-topic Partition: 1    Leader: 3       Replicas: 3,2,1 Isr: 3,2,1
Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2 Isr: 1    
Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1 Isr: 2,3,1
Topic: my-topic Partition: 4    Leader: 3       Replicas: 3,1,2 Isr: 3,1,2
```

```
/kafka/bin/kafka-topics.sh --zookeeper
```

* 토픽 보기

* 토픽 그룹 보기 / p184



토픽생성

보내기, 받기


---
test

docker run --name producer -it ubuntu bash
docker run --name consumer -it ubuntu bash

hosts 설정

# Producer

cd /kafka/bin


* Topic 목록


$KAFKA_HOME/bin/kafka-topics.sh --zookeeper zoo1:2181,zoo2:2181,zoo3:2181 --replication-factor 3 --partitions 5 --topic my-topic --create

$KAFKA_HOME/bin/kafka-console-producer.sh --broker-list $KAFKA --topic my-topic

$KAFKA_HOME/bin/kafka-console-consumer.sh --bootstrap-server $KAFKA --topic my-topic --from-beginning


kafka-console-producer.bat --broker-list localhost:8001,localhost:8002,localhost:8003 --topic my-topic



 * 실시간 처리 위함
 배치 전송
 분산 처리
파티셔닝
