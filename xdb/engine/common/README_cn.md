# common

common 提供 engine 模块中的通用功能，包含挑战、健康度、文件迁移和节点管理等。

## 模块划分
- challenge: 包含挑战生成和存储相关功能，如文件上传、文件续期、文件迁移的场景下，挑战内容的更新；
- health: 包含节点、文件、文件系统健康度等功能，如获取系统中健康的存储节点、计算文件的健康度和系统健康度等；
- migrate: 包含文件迁移相关功能，如文件的拉取及恢复、文件重新加密、文件扩容等功能；
- nodes: 包含节点相关功能，主要是统计存储节点心跳数，用来衡量节点健康度。
