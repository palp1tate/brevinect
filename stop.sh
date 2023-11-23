#!/bin/bash

set -e  # 启用错误模式

# 指定服务端口列表
port_list=("8080" "8081" "8082" "8083" "8084")  # 替换成你实际的服务端口列表

# 遍历端口列表，关闭对应的服务进程
for port in "${port_list[@]}"; do
    # 使用lsof命令查找端口对应的进程PID，并使用kill命令关闭进程
    pid=$(lsof -t -i:"$port")
    if [ -n "$pid" ]; then
        if kill -9 "$pid"; then
            echo "Closed service running on port $port"
        else
            echo "Failed to close service on port $port"
        fi
    else
        echo "No process found on port $port"
    fi
done

echo "Services closed."
