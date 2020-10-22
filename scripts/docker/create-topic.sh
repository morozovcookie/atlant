#!/bin/sh

check_docker_existence() {
  if ! command -v docker > /dev/null
  then
      echo "error: docker could not be found"
      exit 1
  fi

  echo "found: docker"

  DOCKER_BIN_PATH=$(command -v docker)
}

check_container() {
  CONTAINER_NAME=$1

  if [ -z "${CONTAINER_NAME}" ]
  then
    echo "error: empty container name"
    exit 2
  fi

  if [ -z "$(${DOCKER_BIN_PATH} ps --filter name="${CONTAINER_NAME}" --format "{{.ID}}")" ]
  then
    echo "error: container $(CONTAINER_NAME) was not running"
    exit 3
  fi

  echo "found: ${CONTAINER_NAME} container"
}

create_topic() {
  KAFKA_CONTAINER_NAME=$1
  KAFKA_TOPICS_BIN=$2
  ZOOKEEPER_ADDRESS=$3
  TOPIC_NAME=$4
  TOPIC_PARTITIONS=$5
  TOPIC_REPLICATION_FACTOR=$6
  MIN_INSYNC_REPLICAS=$7


  ${DOCKER_BIN_PATH} exec \
    "${KAFKA_CONTAINER_NAME}" \
    "${KAFKA_TOPICS_BIN}" \
    --zookeeper "${ZOOKEEPER_ADDRESS}" \
    --create \
    --topic "${TOPIC_NAME}" \
    --partitions "${TOPIC_PARTITIONS}" \
    --replication-factor "${TOPIC_REPLICATION_FACTOR}" \
    --config min.insync.replicas="${MIN_INSYNC_REPLICAS}"
}

check_docker_existence

# verify if atlant-zookeeper is running
ZOOKEEPER_CONTAINER_NAME=atlant-zookeeper

check_container "$ZOOKEEPER_CONTAINER_NAME"

# verify if atlant-kafka-1 is running
KAFKA_CONTAINER_NAME=atlant-kafka-1

check_container "$KAFKA_CONTAINER_NAME"

# execute kafka-topics.sh outside docker container
SCALA_VERSION=2.13
KAFKA_VERSION=2.6.0
KAFKA_TOPICS_BIN=/opt/kafka_${SCALA_VERSION}-${KAFKA_VERSION}/bin/kafka-topics.sh

ZOOKEEPER_ADDRESS=zookeeper:2181

PRODUCTS_TOPIC_NAME=docker.atlant.cdc.products.0
PRODUCTS_TOPIC_PARTITIONS=3
PRODUCTS_TOPIC_REPLICATION_FACTOR=3
PRODUCTS_MIN_INSYNC_REPLICAS=2

create_topic "$KAFKA_CONTAINER_NAME" \
  "$KAFKA_TOPICS_BIN" \
  "$ZOOKEEPER_ADDRESS" \
  "$PRODUCTS_TOPIC_NAME" \
  "$PRODUCTS_TOPIC_PARTITIONS" \
  "$PRODUCTS_TOPIC_REPLICATION_FACTOR" \
  "$PRODUCTS_MIN_INSYNC_REPLICAS"
