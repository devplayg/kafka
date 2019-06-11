## 테스트 시나리오

### Kafka 상태확인 명령어

    $KAFKA_HOME/bin/kafka-topics.sh --bootstrap-server $KAFKA --list
    $KAFKA_HOME/bin/kafka-topics.sh --bootstrap-server $KAFKA --describe --topic my-topic
    
### Kafka 장애 발생 시, 상태확인 테스트

순차적 장애 발생 후, 순차적 복구

| Seq | kafka1 | kafka2   | kafka3   | 설명             |
|-----|--------|----------|----------|------------------|
| 1   | up     | up       | up       |                  |
| 2   | up     | up       | ~~down~~ | kafka3 장애 발생  |
| 3   | up     | ~~down~~ | ~~down~~ | kafka2 장애 발생  |
| 4   | up     | up       | ~~down~~ | kafka2 장애 복구  |
| 5   | up     | up       | up       | kafka3 장애 복구  |


* Seq 1. 정상

    ```
    Topic: my-topic Partition: 0    Leader: 3       Replicas: 2,1,3    Isr: 3,2,1    
    Topic: my-topic Partition: 1    Leader: 3       Replicas: 3,2,1    Isr: 3,2,1    
    Topic: my-topic Partition: 2    Leader: 3       Replicas: 1,3,2    Isr: 3,2,1    
    Topic: my-topic Partition: 3    Leader: 3       Replicas: 2,3,1    Isr: 3,2,1    
    Topic: my-topic Partition: 4    Leader: 3       Replicas: 3,1,2    Isr: 3,2,1    
    ```

* Seq 2. kafka3 장애 발생

    ```
    Topic: my-topic Partition: 0    Leader: 2       Replicas: 2,1,3    Isr: 2,1      
    Topic: my-topic Partition: 1    Leader: 2       Replicas: 3,2,1    Isr: 2,1      
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 2,1      
    Topic: my-topic Partition: 3    Leader: 2       Replicas: 2,3,1    Isr: 2,1      
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 2,1      
    ```

* Seq 3. kafka2 장애 발생

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1   
    ```

* Seq 4. kafka2 장애 복구

    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1,2   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1,2   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1,2   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1,2   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1,2   
    ```

* Seq 5. 정상
    
    ```
    Topic: my-topic Partition: 0    Leader: 1       Replicas: 2,1,3    Isr: 1,2,3   
    Topic: my-topic Partition: 1    Leader: 1       Replicas: 3,2,1    Isr: 1,2,3   
    Topic: my-topic Partition: 2    Leader: 1       Replicas: 1,3,2    Isr: 1,2,3   
    Topic: my-topic Partition: 3    Leader: 1       Replicas: 2,3,1    Isr: 1,2,3   
    Topic: my-topic Partition: 4    Leader: 1       Replicas: 3,1,2    Isr: 1,2,3   
    ```

### 그룹 "testgroup2"의 Offset을 0으로 초기화

변경

```
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --topic test-offset --reset-offsets --to-earliest --execute
TOPIC                          PARTITION  NEW-OFFSET
test-offset                    2          0
test-offset                    1          0
test-offset                    0          0
test-offset                    4          0
test-offset                    3          0
```

확인

```
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
---
Consumer group 'testgroup2' has no active members.                                                                         
                                                                                                                           
TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID       
test-offset     2          0               60              60              -               -               -               
test-offset     1          0               61              61              -               -               -               
test-offset     0          0               57              57              -               -               -               
test-offset     4          0               64              64              -               -               -               
test-offset     3          0               58              58              -               -               -               
```


### 그룹 "testgroup2"의 Offset을 40으로 변경

변경

```
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --topic test-offset --reset-offsets --to-offset 40  --execute
---
TOPIC                          PARTITION  NEW-OFFSET
test-offset                    2          40
test-offset                    1          40
test-offset                    0          40
test-offset                    4          40
test-offset                    3          40
```

확인

```
$KAFKA_HOME/bin/kafka-consumer-groups.sh --bootstrap-server $KAFKA --group testgroup2 --describe
Consumer group 'testgroup2' has no active members.

TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID     HOST            CLIENT-ID
test-offset     2          40              60              20              -               -               -
test-offset     1          40              61              21              -               -               -
test-offset     0          40              57              17              -               -               -
test-offset     4          40              64              24              -               -               -
test-offset     3          40              58              18              -               -               -
```
