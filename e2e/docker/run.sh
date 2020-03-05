#!/bin/bash

/usr/bin/mongod --config /etc/mongod.conf &

sudo -u elasticsearch /usr/share/elasticsearch/bin/elasticsearch &

/usr/share/graylog-server/bin/graylog-server &

wait -n
