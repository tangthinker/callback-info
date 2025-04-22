# Callback Info

一个简单高效的Web应用，用于测试和收集回调请求，提供清晰的Web界面来展示收集到的数据。

## 功能特点

- 接收并存储回调请求
- 记录请求IP、参数和时间戳
- 返回200 OK响应
- Web界面展示收集的数据
- 支持分页和排序
- 最多存储10000条最新记录
- JSON参数格式化和高亮显示
- 可配置根目录和服务器端口
- 自动清理旧记录
- 实时数据显示

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
- 根目录：`.`（当前目录）
- 服务器端口：`8080`
- 数据库文件：`callback.db`（位于根目录）

### 自定义配置
```bash
# 自定义根目录
go run main.go -root /path/to/your/root

# 自定义服务器端口
go run main.go -port 9090

# 同时自定义根目录和端口
go run main.go -root /path/to/your/root -port 9090
```

### 测试回调
发送POST请求来测试回调：
```bash
# 基本测试
curl -X POST -H "Content-Type: application/json" -d '{"test":"data"}' http://localhost:8080/api/callback

# 复杂JSON测试
curl -X POST -H "Content-Type: application/json" -d '{
  "name": "test",
  "value": 123,
  "items": ["item1", "item2"],
  "nested": {
    "key": "value"
  }
}' http://localhost:8080/api/callback
```

## Web界面

访问 `http://localhost:8080` 可以：
- 查看所有回调记录
- 查看带语法高亮的格式化JSON参数
- 分页浏览
- 按时间戳排序（最新的在前）
- 查看请求详情，包括：
  - 时间戳
  - IP地址
  - 请求参数

## 数据存储

- 使用SQLite数据库进行高效存储
- 数据库文件存储在根目录下
- 最多存储10000条最新记录
- 达到限制时自动删除最旧的记录
- 记录包含：
  - 请求时间戳
  - IP地址
  - 请求参数
- 自动清理数据以保持性能

## 开发

### 构建
```bash
go build -o callback-info
```

### 运行测试
```bash
go test ./...
```

## 贡献

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m '添加了一些很棒的特性'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启一个 Pull Request

## 许可证

MIT许可证

版权所有 (c) 2024 Callback Info

特此免费授予任何获得本软件副本和相关文档文件（以下简称"软件"）的人不受限制地处理本软件的权利，包括但不限于使用、复制、修改、合并、发布、分发、再许可和/或出售本软件副本的权利，以及允许获得本软件的人这样做的权利，但须符合以下条件：

上述版权声明和本许可声明应包含在本软件的所有副本或主要部分中。

本软件按"原样"提供，不提供任何形式的明示或暗示的保证，包括但不限于对适销性、特定用途的适用性和非侵权性的保证。在任何情况下，作者或版权持有人均不对任何索赔、损害或其他责任负责，无论是在合同诉讼、侵权诉讼或其他诉讼中，由软件或软件的使用或其他交易引起、由软件引起或与之相关的任何索赔、损害或其他责任。 