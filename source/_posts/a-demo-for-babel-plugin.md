---
title: 编写babel-plugin 模拟kotlin中的also等扩展方法
date: 2018-02-21 01:56:48
tags: [javascript, babel plugin, kotlin]
---
写js的时候，会经常怀念kotlin中的各种奇淫技巧，
特别是also，let之类的扩展方法 以及Elvis操作符等。

显然，为Object添加属性是实现扩展方法的可选方案，
然而，那样做不好。

于是就想到了，借助babel插件，模拟这些kotlin扩展方法

<!-- more -->

### 使用方式
一：安装
```bash
npm install --save-dev babel-plugin-kotlish-also
```
二：在.babelrc添加如下plugin
```json
{
  "plugins": [
    "kotlish-also"
  ]
}
```

### 使用场景
来看下面这段代码
```javascript
function foo(o) {
  return o.a().b.c
}
```
如果我们调试时想要输出 `o` 或 `o.a()` 等，
通常情况下，我们不得不做出如下改变
```javascript 
// if we want to print o or o.a etc
function foo(o) {
  console.log(o)
  const t = o.a()
  console.log(t)
  return t.b.c
}
```
当调试完毕，又要改回原有代码，
也算是相当麻烦，
当使用了 `kotlish-also` 后，可以写成这样
```javascript 
function fooFull(o) {
  return o.also(i => console.log(i))
    .a().also(i => console.log(i))
    .b.c
}
```
或者使用`'it'`省略匿名函数也可以
```javascript
function foo(o) {
  return o.also(console.log(it))
    .a().also(console.log(it))
    .b.c
}
```
插件最终会将代码转换成如下形式
```javascript
function foo(o) {
  return function (_o) {
    (function (it) {
      return console.log(it);
    })(_o);

    return _o;
  }(function (_o) {
    (function (it) {
      return console.log(it);
    })(_o);

    return _o;
  }(o).a()).b.c;
}
```

甚至，插件为了方便输出，增加了alsoPrint方法
```javascript
const a = obj.alsoPrint().a
```
等价于
```javascript
const a = obj.also(function (_o) {
  console.log(_o)
}).a
```

除了also，
kotlish-also 同时提供的方法还有let,takeIf,takeUnless等，
它们都是仿照kotlin编写的，其具体语法可参考 [kotlin stdlib functions](http://kotlinlang.org/api/latest/jvm/stdlib/kotlin/index.html#functions)

[原文出处](http://blog.cloudself.cn/2018/02/21/a-demo-for-babel-plugin/)
[Github](https://github.com/HerbLuo/babel-plugin-kotlish-also)
