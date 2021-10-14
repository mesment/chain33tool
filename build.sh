CGOENABLE=0 GOARCH=amd64 GOOS=linux go build  -o chain33tool-linux
CGOENABLE=0 GOARCH=amd64 GOOS=darwin go build  -o chain33tool-macos
CGOENABLE=0 GOARCH=amd64 GOOS=windows go build  -o chain33tool.exe