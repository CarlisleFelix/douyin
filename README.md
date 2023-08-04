# 各个目录和文件的功能

## config

分类别存放各个模块的config参数，config.go中的结构与config.yaml一一对应，config包中其余文件为每个模块的具体config设置。详情见server示例。

## controller

## database

负责数据库层面的操作，如增删改查以及建表删表等。

## doc

负责放置代码文档。

## init

负责项目的初始化操作，如读取配置等。

## log

负责放置各个模块的日志，注意每个模块应该有自己的日志，不可将日志全部输出到同一个文件中。

## middleware

## model

## router

## service

## utils

## config.yaml

项目的初始化设置，供viper进行读取。

## main.go

项目主文件。