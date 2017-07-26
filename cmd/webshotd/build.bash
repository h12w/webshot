#!/bin/bash

# https://github.com/gotk3/gotk3/issues/152

if [[ $(lsb_release -cs) = xenial ]]; then
go install -tags gtk_3_18 github.com/gotk3/gotk3/gtk
go build -v -tags gtk_3_18 -gcflags "-N -l"
docker build . -t h12w/webshotd
docker push h12w/webshotd
#docker run --name webshotd --publish 9191:9191 --detach h12w/webshotd
fi
