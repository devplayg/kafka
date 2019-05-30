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

## 컨테이너 개별 설정

* 필수 패키지 다운로드
    
    ```
    sed -i s#archive\.ubuntu\.com#mirror.kakao.com#g /etc/apt/sources.list
    apt update && apt install -y vim net-tools openjdk-11-jdk inetutils-ping telnet curl
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
    
## 포트 정보

* `Kafka`: 9092
* `Zookeeper`: 2181, 2888:3888

## Tip

    vi ~/.profile
    ---
    alias .pro='vi ~/.profile'
    alias pro='. ~/.profile'
    alias kconf='vi /kafka/config/server.properties'
    alias zconf='vi /kafka/config/zookeeper.properties'
    alias zstart='/kafka/bin/zookeeper-server-start.sh /kafka/config/zookeeper.properties'
    alias kstart='/kafka/bin/kafka-server-start.sh /kafka/config/server.properties'

#### [Zookeeper 설정](docs/install_zookeeper.md)
#### [Kafka 설정](docs/install_kafka.md)


