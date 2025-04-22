# Callback Info

一个用于测试和收集回调请求的简单Web应用，提供清晰的Web界面来展示收集到的数据。

## 功能特点

- 接收并存储回调请求
- 记录请求IP、参数和时间戳
- 返回200 OK响应
- Web界面展示收集的数据
- 支持分页和排序
- 最多存储10000条最新记录
- JSON参数格式化和高亮显示
- 可配置数据库位置和服务器端口

## 安装

1. 克隆仓库：
```bash
git clone https://github.com/tangthinker/callback-info.git
cd callback-info
```

2. 安装依赖：
```bash
go mod tidy
```

## 使用方法

### 基本使用
```bash
go run main.go
```
这将使用默认设置启动服务器：
- 数据库文件：`./callback.db`
- 服务器端口：`8080`

### 自定义配置
```bash
# 自定义数据库位置
go run main.go -db /path/to/your/callback.db

# 自定义服务器端口
go run main.go -port 9090

# 同时自定义数据库位置和端口
go run main.go -db /path/to/your/callback.db -port 9090
```

### 测试回调
发送POST请求来测试回调：
```bash
curl -X POST -H "Content-Type: application/json" -d '{"test":"data"}' http://localhost:8080/api/callback
```

## Web界面

访问 `http://localhost:8080` 可以：
- 查看所有回调记录
- 查看格式化的JSON参数
- 分页浏览
- 按时间戳排序

## 数据存储

- 使用SQLite数据库
- 最多存储10000条最新记录
- 达到限制时自动删除最旧的记录
- 记录包含：
  - 请求时间戳
  - IP地址
  - 请求参数

## 许可证

MIT许可证 