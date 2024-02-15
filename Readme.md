# Aine Linux Tools - Delete

> Version 1.0.0 (2024-02-15 Release)
>
> Author: Noah Jones
>

## 简介

Del (Delete) 隶属于 aine-linux-tools 中的一部分，是一个简单、安全的文件/文件夹删除工具，此工具类似于rm工具，del工具不会立刻删除，而是将删除的文件移动到用户根目录下存储，在七天后才会真正的删除。在此期间内，用户可以通过-r将文件恢复到原位置，因此非常适合对linux不熟悉的小白用户使用。目前计划仅适用于Linux相关发行版，暂无对windows/macos的支持（win/macos此功能已经足够便捷，macos理论也可以使用，但没有维护计划）

## 安装

将对应平台的可执行文件路径添加到环境变量即可

## 使用

执行如下命令即可：

```bash
# 同时删除多个文件
del a.log b.log c.log

# 查看所有已删除的文件
del -l

# 根据文件名筛选已删除的文件
del -l | grep a1

# 根据文件路径筛选
del -l | grep /home/user/a1

# 恢复指定id的文件
del -r [your_id]

# 清空回收站
del -c

# 永久删除指定id的文件
del -f a.log b.log c.log

```
