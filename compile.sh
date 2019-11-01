#!/bin/bash

SRC_DIR=`pwd`
DST_DIR=`pwd`
rm $DST_DIR/dpt.pb.go
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/dpt.proto
