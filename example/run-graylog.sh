#!/bin/bash

function start_containers {
        mongo_id=$(set -x; docker run \
            -d mongo:3)

        elasticsearch_id=$(set -x; docker run \
            -d elasticsearch:2 \
            elasticsearch -Des.cluster.name="graylog")

        graylog_id=$(set -x; docker run \
            --link ${mongo_id}:mongo \
            --link ${elasticsearch_id}:elasticsearch \
            -d graylog/graylog)

        graylog_ip=$(docker inspect \
            ${graylog_id} \
            -f "{{.NetworkSettings.IPAddress}}")
}

function stop_containers {
    (
        set -x
        docker stop ${mongo_id:-} ${elasticsearch_id:-} ${graylog_id:-}
        docker rm ${mongo_id:-} ${elasticsearch_id:-} ${graylog_id:-}
    )
}

set -euo pipefail

trap stop_containers ERR EXIT INT TERM
start_containers

echo "Graylog ready"
echo "export GRAYLOG_SERVER_URL=http://${graylog_ip}:9000/"
echo "export GRAYLOG_USERNAME=admin"
echo "export GRAYLOG_PASSWORD=admin"

docker wait ${mongo_id} ${elasticsearch_id} ${graylog_id}
