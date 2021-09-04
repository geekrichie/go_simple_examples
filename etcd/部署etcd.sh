#!/bin/bash

etcdpath=$(pwd)
nohup ${etcdpath}/etcd --name infra1 \
--listen-client-urls http://192.168.33.13:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.33.13:2379 \
--listen-peer-urls http://192.168.33.13:12380 \
--initial-advertise-peer-urls http://192.168.33.13:12380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://192.168.33.13:12380,infra2=http://192.168.33.13:22380,infra3=http://192.168.33.13:32380' \
--initial-cluster-state new &

nohup ${etcdpath}/etcd --name infra2 \
--listen-client-urls http://192.168.33.13:12379,http://127.0.0.1:12379 \
--advertise-client-urls http://192.168.33.13:12379 \
--listen-peer-urls http://192.168.33.13:22380 \
--initial-advertise-peer-urls http://192.168.33.13:22380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://192.168.33.13:12380,infra2=http://192.168.33.13:22380,infra3=http://192.168.33.13:32380' \
--initial-cluster-state new &

nohup ${etcdpath}/etcd --name infra3 \
--listen-client-urls http://192.168.33.13:22379,http://127.0.0.1:22379 \
--advertise-client-urls http://192.168.33.13:22379 \
--listen-peer-urls http://192.168.33.13:32380 \
--initial-advertise-peer-urls http://192.168.33.13:32380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://192.168.33.13:12380,infra2=http://192.168.33.13:22380,infra3=http://192.168.33.13:32380' \
--initial-cluster-state new &