FROM ubuntu:16.04
MAINTAINER Wáng Hǎiliàng <w@h12.me>

RUN apt-get update && apt-get install --yes \
    xvfb \
    libwebkit2gtk-4.0-37 \
    fonts-wqy-microhei

ADD webshot /bin
ADD run.bash /bin/run.bash

ENV DISPLAY=:0
ENTRYPOINT /bin/run.bash
