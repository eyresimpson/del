# Aine Linux Tools - Delete

> Version 0.0.1 (2024-02-05 Release)
>
> Author: Noah Jones
>

## 简介

Del (Delete) 隶属于 aine-linux-tools 中的一部分，是一个简单、安全的文件/文件夹删除工具，需要和 Rec (Recover) 配合使用，此工具目前计划仅适用于Linux相关发行版，暂无对windows/macos的支持（win/macos此功能已经足够便捷）

类似于rm工具，del工具不会立刻删除，而是将删除的文件移动到用户根目录下存储，在七天后才会真正的删除。在此期间内，用户可以通过Rec工具将文件恢复到原位置，因此非常适合对linux不熟悉的小白用户使用。

## 安装

将对应平台的可执行文件路径添加到环境变量即可，无需安装。

## 使用

执行如下命令即可：

```bash
# 同时删除多个文件
del a.log b.log c.log

# 强制删除，且删除目录时不提示
del -rf /opt/app/jar-lib test1.jar test1.jar
```
