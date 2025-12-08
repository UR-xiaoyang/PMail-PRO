#!/bin/bash

# 配置输出目录
OUTPUT_DIR="../release"
OUTPUT_NAME="pmail"

# 1. 编译前端网站
echo "start building fe"
cd fe 
npm install && npm run build
if [ $? -ne 0 ]; then
    echo "Frontend build failed"
    exit 1
fi
cd ..

# 注意：vite.config.js 已经配置了直接输出到 server/listen/http_server/dist
# 所以不需要再手动复制 dist 文件夹

# 2. 编译 go 语言项目
echo "start building server"
cd server

# 确保输出目录存在
mkdir -p $OUTPUT_DIR

go build -o "$OUTPUT_DIR/$OUTPUT_NAME" main.go
if [ $? -ne 0 ]; then
    echo "Backend build failed"
    exit 1
fi

cd ..
echo "build success! Output file is in release/$OUTPUT_NAME"
