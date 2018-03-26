---
title: duu
date: 2018-03-21 23:06:45
tags: [duu, javascript]
---

实习中实现的一种配置式的后台管理系统的开发方案，
在此做个回忆，梳理和记录，
有机会再整理，重写并优化。

<!-- more -->

### 一：概述

这种配置式的开发方案仅可用于前端，且通用性暂时不大，
实现部分十分基础，并没有什么高深的东西，
本文仅节选部分配置属性，从`[数据，配置，UI层，配置处理层]`进行说明。

### 二：数据（以人员信息为例，其中所属部门为下拉框）

#### 1: 服务端返回与接收的数据格式
人员信息（Data）
```json
{
  "name": "小明",
  "department": 2 // 所属部门编号
}
```
部门信息（也就是下拉框中的数据，DetailData）
```json
[{
  "id": 2,
  "departmentName": "技术部"
}]

```

#### 2: UI层接受的数据格式为
```json
{
  "name": "小明",
  "department": 2, 
  "_department": { // 并非直接显示在视图上的数据
    "list": [{
      "_value": 2, // 返回服务器的
      "_label": "技术部" // 展示给用户的
    }],
    "_label": "", // 当前的
    "_value": "" // 当前的
  }
}
```

### 三：配置

#### 1. 所有的数据处理（例如，由[服务端数据](#1-服务端返回与接收的数据格式)转换为[UI接受的数据](#2-UI层接受的数据格式为)）都在配置中完成

#### 2. 配置大致如下
```javascript 
const mapping = { // 对服务端返回数据的一种解释
  name: {
    label: '姓名',
    tip: '不超过5个字',
    nullable: false,
    validate: [
      {required: true, message: '姓名不能为空'},
      {max: 5, message: '姓名不超过5个字'},
      {pattern: pattern.zhCN, message: '只允许输入中文'}
    ]
  },
  department: {
    label: '部门',
    list: { // list会最终转换为 initData, onDataLoaded, fetchDetail 这种东西
      key: {
        dataLabel: '',
        dataValue: 'department',
        detailLabel: 'departmentName',
        detailValue: 'id'
      },
      fetch: api.getDepartment,
      default: 0
    }
    // initData (data, payload) {}, // 各种处理数据的回调，他们的执行时期不同
    // onDataLoaded (data, paylaod) {},
    // fetchDetail (data, payload) {},
    // onDetailLoaded (data, list, payload) {},
    // beforeDataIsPost (data, payload) {},
    // onSelected (value, data, payload) {}
  }
}

const config = {
  create: {}, // 创建界面的相关配置
  modify: {}, // 修改界面的相关配置
  
  mapping,
  get: requestBody => fetch(),
  patch: data => fetch()
}

registerConfig('pageName', config) // 注册该配置
```

### 四：UI层

#### 1. UI层使用模板语言处理并展示`[label, tip, nullable]`等基本配置

#### 2. UI层不对数据处理，仅依次调用 `[initData, onDataLoaded, fetchDetail, onDetailLoaded]`等配置中的方法来完成数据处理
1. 优点：使配置拥有完整的操作数据的能力，这一套方案，可以处理`[下拉框，时间范围选择框，远程搜索框，树组件下拉框]`等
2. 缺点：可能导致配置冗长

### 五：配置文件的处理与格式化

为了避免配置文件的冗长，
如[三.2](#2-配置大致如下) 中的 `mapping.department`，
提供`list`这种高糖代替`各种处理数据的回调`，`list配置`最终会被转换成`默认的处理数据的回调`，
这种处理会在注册配置时调用。
