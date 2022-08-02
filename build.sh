#!/bin/bash

name='kubernetes-rbac-example'
version='1.0.0'
tag="semihbkgr/$name:$version"

docker build --tag $tag .

while true; do
    read -p "Do you wish to upload the image? [y/n] : " yn
    case $yn in
        [Yy]* ) docker push $tag; break;;
        [Nn]* ) exit;;
        * ) echo "[y/n] : ";;
    esac
done
