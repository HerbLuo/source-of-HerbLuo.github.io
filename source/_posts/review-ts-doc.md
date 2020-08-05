---
title: typescript 官方文档再读
date: 2019-05-01 00:00:00
tags: [typescript]
---
对于从其他静态强类型语言转过来的，或者熟悉多种语言的开发者来说，可能入门`typescript`的时候只是走马观花式地看了一遍官方文档，从而漏过很多重要内容。  
随着了解的深入，越发觉得`typescript`的类型系统无比强大，是时候再精读一遍文档了。  
本文结合实际开发体验，对一些重点常用内容进行摘录。

<!-- more -->

1. `ts 2.8`: `infer` 关键字
```typescript
type ReturnType<T> = T extends (...args: any[]) => infer R ? R : any;
```
`infer` 一般与 `extends` 一起使用（至少目前为止，我没有见到他独立存在）。用于提取内部类型。    
再来一个手写的例子。
```typescript
type PromiseType<T> = T extends Promise<infer U> ? U : never;
```
