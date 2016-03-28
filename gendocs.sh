#!/bin/bash

#删除之前的
rm "api/docs.go"
swagger -apiPackage="app-server/api/v1" -mainApiFile="app-server/api/routers.go" -output="api/"
