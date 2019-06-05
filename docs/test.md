## 테스트


토픽 생성

### Kafka 장애 발생 시, 상태확인 테스트

| Seq | kafka1 | kafka2 | kafka3 | 설명             |
|-----|--------|--------|--------|------------------|
| 1   | up     | up     | up     |                  |
| 2   | up     | up     | down   | kafka3 장애 발생  |
| 3   | up     | down   | down   | kafka2 장애 발생  |
| 4   | up     | up     | down   | kafka2 장애 복구  |
| 5   | up     | up     | up     | kafka3 장애 복구  |

상태확인 명령어

```
$KAFKA_HOME/bin/kafka-topics.sh --bootstrap-server $KAFKA --list
$KAFKA_HOME/bin/kafka-topics.sh --bootstrap-server $KAFKA --describe --topic my-topic
```


* Seq. 1

    ```
    Topic: my-topic Partition: 0    Leader: 3       Replicas: 2,1,3    Isr: 3,2,1    
    Topic: my-topic Partition: 1    Leader: 3       Replicas: 3,2,1    Isr: 3,2,1    
    Topic: my-topic Partition: 2    Leader: 3       Replicas: 1,3,2    Isr: 3,2,1    
    Topic: my-topic Partition: 3    Leader: 3       Replicas: 2,3,1    Isr: 3,2,1    
    Topic: my-topic Partition: 4    Leader: 3       Replicas: 3,1,2    Isr: 3,2,1    
    ```

* Seq. 2

    ```
    Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3    Isr: 2,1      
    Topic: my-topic Partition: 1    Leader: 2       Replicas: 3,2,1    Isr: 2,1      
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 2,1      
    Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1    Isr: 2,1      
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 2,1      
    ```

* Seq. 3

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1   
    ```

* Seq. 4

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1,2   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1,2   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1,2   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1,2   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1,2   
    ```

* Seq. 5
    
    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1,2,3   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1,2,3   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1,2,3   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1,2,3   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1,2,3   
    ```

### Test case #2

| Seq | kafka1 | kafka2 | kafka3 |
|-----|--------|--------|--------|
| 1   | up     | up     | up     |
| 2   | up     | up     | down   |
| 3   | up     | down   | down   |
| 4   | down   | down   | down   |
| 5   | down   | down   | up     |
| 6   | down   | up     | up     |
| 7   | up     | up     | up     |

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
Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3    Isr: 2,1,3
Topic: my-topic Partition: 1    Leader: 3       Replicas: 3,2,1    Isr: 3,2,1
Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1    
Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1    Isr: 2,3,1
Topic: my-topic Partition: 4    Leader: 3       Replicas: 3,1,2    Isr: 3,1,2
```

### Offset 확인
```
토픽 생성
$KAFKA_HOME/bin/kafka-topics.sh --zookeeper $ZOOKEEPER --replication-factor 3 --partitions 5 --topic test-offset --create
ㄱ
토픽 상세
$KAFKA_HOME/bin/kafka-topics.sh --bootstrap-server $KAFKA --describe --topic test-offset


$KAFKA_HOME/bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list $KAFKA -topic test-offset --time -1

Consumer 그룹 보기
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --list

Consume topic
consumer-group -brokers $KAFKA -group testgroup2 -topics test-offset

testgroup 상태보기
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
---
TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                 HOST            CLIENT-ID
test-offset     2          37              60              23              sarama-ff21d16e-8aa6-4388-bce3-da4696e7924b /172.17.0.9     sarama
test-offset     1          42              61              19              sarama-ff21d16e-8aa6-4388-bce3-da4696e7924b /172.17.0.9     sarama
test-offset     0          37              57              20              sarama-ff21d16e-8aa6-4388-bce3-da4696e7924b /172.17.0.9     sarama
test-offset     4          47              64              17              sarama-ff21d16e-8aa6-4388-bce3-da4696e7924b /172.17.0.9     sarama
test-offset     3          37              58              21              sarama-ff21d16e-8aa6-4388-bce3-da4696e7924b /172.17.0.9     sarama


Topic, Consumer 그룹 OFfset 변경
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --topic test-offset --reset-offsets --to-earliest --execute
Consumer group 'testgroup2' has no active members.                                                                         
                                                                                                                           
TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID       
test-offset     2          0               60              60              -               -               -               
test-offset     1          0               61              61              -               -               -               
test-offset     0          0               57              57              -               -               -               
test-offset     4          0               64              64              -               -               -               
test-offset     3          0               58              58              -               -               -               
```

```
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
Consumer group 'testgroup2' has no active members.

TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
test-offset     2          60              60              0               -               -               -
test-topic      4          106             106             0               -               -               -
test-topic      2          122             122             0               -               -               -
test-offset     1          61              61              0               -               -               -
test-offset     0          57              57              0               -               -               -
test-topic      3          113             113             0               -               -               -
test-offset     4          64              64              0               -               -               -
test-topic      1          112             112             0               -               -               -
test-topic      0          107             107             0               -               -               -
test-offset     3          58              58              0               -               -               -
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --topic test-offset --reset-offsets --to-offset 40  --execute

TOPIC                          PARTITION  NEW-OFFSET
test-offset                    2          40
test-offset                    1          40
test-offset                    0          40
test-offset                    4          40
test-offset                    3          40
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
Consumer group 'testgroup2' has no active members.

TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
test-offset     2          40              60              20              -               -               -
test-topic      4          106             106             0               -               -               -
test-topic      2          122             122             0               -               -               -
test-offset     1          40              61              21              -               -               -
test-offset     0          40              57              17              -               -               -
test-topic      3          113             113             0               -               -               -
test-offset     4          40              64              24              -               -               -
test-topic      1          112             112             0               -               -               -
test-topic      0          107             107             0               -               -               -
test-offset     3          40              58              18              -               -               -
```

```go
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
Consumer group 'testgroup2' has no active members.

TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
test-offset     2          60              60              0               -               -               -
test-topic      4          106             106             0               -               -               -
test-topic      2          122             122             0               -               -               -
test-offset     1          61              61              0               -               -               -
test-offset     0          57              57              0               -               -               -
test-topic      3          113             113             0               -               -               -
test-offset     4          64              64              0               -               -               -
test-topic      1          112             112             0               -               -               -
test-topic      0          107             107             0               -               -               -
test-offset     3          58              58              0               -               -               -
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --topic test-offset:0,1 --reset-offsets --to-offset 40  --execute

TOPIC                          PARTITION  NEW-OFFSET
test-offset                    0          40
test-offset                    1          40
root@992071e194fb:/gohome/src/github.com/devplayg/kafka/produce_rand_consume_rand# $KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
Consumer group 'testgroup2' has no active members.

TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
test-offset     2          60              60              0               -               -               -
test-topic      4          106             106             0               -               -               -
test-topic      2          122             122             0               -               -               -
test-offset     1          40              61              21              -               -               -
test-offset     0          40              57              17              -               -               -
test-topic      3          113             113             0               -               -               -
test-offset     4          64              64              0               -               -               -
test-topic      1          112             112             0               -               -               -
test-topic      0          107             107             0               -               -               -
test-offset     3          58              58              0               -               -               -
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






### 환경구성

```
docker run -v e:/gohome:/gohome -it --name test ubuntu bash
docker run -v e:/gohome:/gohome -it --name consumer ubuntu bash
```

```
sed -i s#archive\.ubuntu\.com#mirror.kakao.com#g /etc/apt/sources.list
apt update && apt install -y curl git vim
curl https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz | tar xvfz - -C /
vi ~/.profile
------------------------------------------------------------------
export GOROOT=/go
export GOPATH=/gohome
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
------------------------------------------------------------------
```


|No|Producer |Consumer |
|---|---|---|
|1  |랜덤전송| 수신|
|2  |파티션별 전송| 수신|

