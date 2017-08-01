FROM ubuntu:16.04
MAINTAINER Wáng Hǎiliàng <w@h12.me>

# webkit
RUN apt-get update && apt-get install --yes \
    xvfb \
    libwebkit2gtk-4.0-37

# font
RUN apt-get install --yes fonts-noto

ADD webshotd /bin
ADD run.bash /bin

ENV DISPLAY=:0
EXPOSE 9191
ENTRYPOINT  /bin/run.bash
