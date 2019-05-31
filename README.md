# Apache Kafka
 
A distributed streaming platform. https://kafka.apache.org/downloads

## 버전

* `OS`: Docker Ubuntu 18.04
* `Kafka`: 2.2
* `Zookeeper`: 3.4.13
* `Java`: OpenJDK 11

## Docker 컨테이너 구성

|구분       | Hostname | IP         |
|-----------|----------|------------|
|Zookeeper  | zoo1     | 172.17.0.2 |
|           | zoo2     | 172.17.0.3 |
|           | zoo3     | 172.17.0.4 |
|Kafka      | kafka1   | 172.17.0.5 |
|           | kafka2   | 172.17.0.6 |
|           | kafka3   | 172.17.0.7 |

## Docker OS 이미지 다운로드

```
docker pull ubuntu
```
    
## Docker 컨테이너 생성

* Zookeeper 컨테이너 생성

    ```
    docker run --name zoo1 -it ubuntu bash
    docker run --name zoo2 -it ubuntu bash
    docker run --name zoo3 -it ubuntu bash
    ```

* Kafka 컨테이너 생성
  
    ```  
    docker run -p 8001:9092 --name kafka1 -it ubuntu bash
    docker run -p 8002:9092 --name kafka2 -it ubuntu bash
    docker run -p 8003:9092 --name kafka3 -it ubuntu bash
    ```

* 컨테이너 IP 확인

    ```
    docker inspect -f "{{ .NetworkSettings.IPAddress }}" zoo1 zoo2 zoo3
    --- 
    172.17.0.2
    172.17.0.3
    172.17.0.4
    ```
    
    ```
    docker inspect -f "{{ .NetworkSettings.IPAddress }}" kafka1 kafka2 kafka3
    ---
    172.17.0.5
    172.17.0.6
    172.17.0.7
    ```
    
## 호스트 파일 설정

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

## 컨테이너 개별 설정

* 필수 패키지 다운로드
    
    ```
    sed -i s#archive\.ubuntu\.com#mirror.kakao.com#g /etc/apt/sources.list
    apt update && apt install -y vim net-tools inetutils-ping telnet curl openssh-server openjdk-11-jdk
    echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
    echo "root:1234" | chpasswd
    service ssh start
    ```

* Java 버전 확인

    ```
    java -version
    ---
    openjdk version "11.0.3" 2019-04-16
    OpenJDK Runtime Environment (build 11.0.3+7-Ubuntu-1ubuntu218.04.1)
    OpenJDK 64-Bit Server VM (build 11.0.3+7-Ubuntu-1ubuntu218.04.1, mixed mode, sharing)
    ```

* Kafka 설치

    ```
    curl -sL http://apache.mirror.cdnetworks.com/kafka/2.2.0/kafka_2.12-2.2.0.tgz | tar xvz -C /
    mv /kafka_2.12-2.2.0/ /kafka
    ```

## Tip

    ```
    vi ~/.profile && . ~/.profile
    ---
    export ZOOKEEPER=zoo1:2181,zoo2:2181,zoo3:2181/my-kafka
    export KAFKA=kafka1:9092,kafka2:9092,kafka3:9092
    export KAFKA_HOME=/kafka
    
    alias .pro='vi ~/.profile'
    alias pro='. ~/.profile'
    alias kconf='vi $KAFKA_HOME/config/server.properties'
    alias zconf='vi $KAFKA_HOME/config/zookeeper.properties'
    alias zstart='$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties'
    alias zstop='$KAFKA_HOME/bin/zookeeper-server-stop.sh'
    alias kstart='$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties'
    alias kstop='$KAFKA_HOME/bin/kafka-server-stop.sh'
    ```
    
## 포트 정보

* `Kafka`: 9092
* `Zookeeper`: 2181, 2888:3888

#### [Zookeeper 설정](docs/install_zookeeper.md)
#### [Kafka 설정](docs/install_kafka.md)
#### [Kafka Test](docs/test.md)


