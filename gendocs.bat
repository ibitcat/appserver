
@echo off

del "api\docs.go"
swagger -apiPackage="app-server/api/v1" -mainApiFile="app-server/api/routers.go" -output="api/"

pause