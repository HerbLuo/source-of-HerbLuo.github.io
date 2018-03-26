---
title: Web开发经典问题之跨域
date: 2018-01-13 21:29:17
tags: [cross origin, java, servlet, CORS, jsonp, fetch]
---

前端面试的时候被问的最多的是什么？
就当前的我而言，是跨域。

<!-- more -->

### 简述
很长一段时间，
正纳闷，跨域这事情和前端有毛线关系呢，
其实要强行扯出关系，那也是可以的...
下面主要介绍一下几种常见的跨域的解决方案

### 正文

#### 一般情况（CORS）

如果后端代码在公司或部门掌控之内，
毫无疑问，此时由后端往`Response Headers`
写一个`Access-Control-Allow-Origin`是最方便的，它的值可为`*`，
但**强烈不建议写成`*`<sup><a href="#note1">①</a></sup>**，正确写法是允许`[发布后的域名，localhost，内网]`

- - - - -

#### JSONP
`jsonp` 仅适用于get方法，适用范围并不广，
它主要用于API由第三方提供，且仅存在get方法的情形

- - - - -

#### fetch
`fetch` 同样不是完美的解决方案，它的限制如下：
`fetch option with no-cors`  适用于`post`请求，但只适用于发送请求，
也就是说：服务端可以接收，但浏览器不会给你返回结果

- - - - -

#### 反代
反向代理适用于其它各种原因导致的跨域问题，
同样，他的缺点也很明显：浪费服务器资源。
关于反向代理的搭建方法，有待后续补充，

- - - - -

[原文出处](http://blog.cloudself.cn/2018/01/13/cross-origin/)

### 注
<a name="note1">①</a> 浏览器之所以限制javascript跨域，那肯定是有它的原因，主要是[安全方面](https://www.zhihu.com/question/26379635)的，
为了项目的持续健康发展，请尽量不要使用 `*` 这种不负责任的方式，
这里提供[Java Servlet Filter版的CORS过滤器](http://blog.cloudself.cn/download/cross-origin/CORSFilter.java)作为参考
