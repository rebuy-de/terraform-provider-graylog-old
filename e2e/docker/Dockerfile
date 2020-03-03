# all-in-one image for e2e testing purposes. We use the Ubuntu base image
# instead of a smaller one, because the docs [1] has a compact guide how to
# install it there, which should reduce the required maintenance. We want an
# all-in-one image so we do not need to care about setting up networking on CI
# systems.
#
# This image is not suitable for production use!
#
# [1]: http://docs.graylog.org/en/3.2/pages/installation/os/ubuntu.html
FROM ubuntu:18.04

# Prerequisites

RUN apt-get -y update && apt-get -y upgrade
RUN apt-get -y install apt-transport-https openjdk-8-jre-headless uuid-runtime pwgen wget sudo

# MongoDB

RUN apt-get -y install gnupg2
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 9DA31620334BD75D9DCB49F368818C72E52529D4
RUN echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.0 multiverse" > /etc/apt/sources.list.d/mongodb-org-4.0.list
RUN apt-get -y update
RUN apt-get -y install mongodb-org

# Elasticsearch

RUN wget -q https://artifacts.elastic.co/GPG-KEY-elasticsearch -O myKey
RUN apt-key add myKey
RUN echo "deb https://artifacts.elastic.co/packages/oss-6.x/apt stable main" >> /etc/apt/sources.list.d/elastic-6.x.list
RUN apt-get -y update && apt-get -y install elasticsearch-oss

# Graylog

RUN wget https://packages.graylog2.org/repo/packages/graylog-3.2-repository_latest.deb
RUN dpkg -i graylog-3.2-repository_latest.deb
RUN apt-get -y update && apt-get -y install graylog-server

# Configure

RUN echo "password_secret = The minimum length for password_secret is 16 characters." >> /etc/graylog/server/server.conf
RUN echo -n "root_password_sha2 = " >> /etc/graylog/server/server.conf
RUN echo -n "admin" | sha256sum | cut -d" " -f1 >> /etc/graylog/server/server.conf
RUN echo "http_bind_address = 0.0.0.0:9000" >> /etc/graylog/server/server.conf

# Other

COPY run.sh /run.sh
ENTRYPOINT ["/bin/bash"]
CMD ["/run.sh"]
