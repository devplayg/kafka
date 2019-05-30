#### [목록](../README.md)

## 테스트 케이스

### 환경설정

```
vi ~/.profile
---
ZOOKEEPER=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
```

### Topics

* Event (name=``)
    - date, src_ip, dst_ip, category, filename,  url 
    
* Traffic (name=`traffic)
    - date, in_bps, out_bps

* Resource (name=`resource)
    - date, cpu usazge, memory usage, disk usage

## 

* Topic 목록 조회

```shell
/kafka/bin/kafka-topics.sh --zookeeper $ZOOKEEPER -list
```

* Topic 상세 조회

```shell
/kafka/bin/kafka-topics.sh --zookeeper $ZOOKEEPER --topic my-topic --describe
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


# Topic 생성
./kafka-topics.sh --zookeeper kafka:2181,kafka2:2181,kafka3:2181 --replication-factor 1 --partitions 1 --topic my-topic --create

 ./kafka-console-producer.sh --broker-list kafka1:9092,kafka2:9092,kafka3:9092 --topic my-topic
 ./kafka-console-consumer.sh --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092 --topic my-topic --from-beginning



 * 실시간 처리 위함
 배치 전송
 분산 처리
파티셔닝


토픽은 미
