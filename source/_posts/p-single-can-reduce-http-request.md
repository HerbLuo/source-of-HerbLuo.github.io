---
title: 前端 之 使用pSingle减少HTTP请求
date: 2017-12-09 02:13:16
tags: [javascript, async, promise]
---
最近总想搞大事，把简单的东西做复杂
pSingle就是一个这样的模块，它可以处理异步操作，使其“单例”

<!-- more -->

### 简介

pSingle 用于处理异步操作
pSingle 包裹的异步操作在完成前不会再次执行。 
只有当 pSingle 包裹的异步操作执行完毕后，pSingle才会使其可能再次运行

pSingle 和 lodash.throttle的区别是，
lodash.throttle 是固定时间只会执行一次
pSingle 是 promise执行完毕前只执行一次


### 安装
```bash
npm install --save p-single
```

### 使用方法

pSingle 接受一个参数，该参数是一个方法，该方法需返回Promise。  

```javascript 
import pSingle from './p-single.js'

// pSingle需要一个返回 Promise的方法
const doWhenClick = pSingle(function promiseReturnedFunc() {
  return new Promise(/*...*/)
})

// 或者直接使用 async
const getAccessToken = pSingle(async (username, password) => {
  // ...
})
```

如果你使用了 ES Next - decorator，你可以这样使用 pSingle

```javascript 
import {PSingle} from './p-single.js'

class Api {
  /**
   * 刷新 token
   * 该请求直到返回前不会再次发送，
   * 该请求返回结果后，每一个调用者都会收到正确的 Promise信号
   */
  @PSingle()
  oauthByRefreshToken (refreshToken) {
    return fetch(/*...*/)
      // ...
  }
}
export default new Api()
```

### 适用场景
- 单击按钮触发事件后，禁止再次触发事件直到事件执行完毕
- 自动登陆模块的设计

### 小提示
建议将 pSingle用于 api层，
尽管用于业务逻辑层可能会提高性能，但与此同时会增加代码复杂度

### 代码实现

pSingle 的代码实现非常简单，只有三十多行

```javascript 
/**
 * p-single
 * change logs:
 * 2017/11/5 herbluo created
 */
export const pSingle = fn => {
  const suspends = []
  let isRunning = false

  return (...args) => new Promise((resolve, reject) => {
    const success = val => {
      resolve(val)
      suspends.forEach(({resolve}) => resolve(val))
      isRunning = false
    }
    const fail = err => {
      reject(err)
      suspends.forEach(({reject}) => reject(err))
      isRunning = false
    }

    if (!isRunning) {
      isRunning = true
      fn(...args).then(success, fail)
    } else {
      suspends.push({resolve, reject})
    }
  })
}

export default pSingle

export const PSingle = (thisBinding) => {
  return (target, property, descriptor) => {
    descriptor.value = pSingle(
      descriptor.value.bind(thisBinding || target)
    )
  }
}
```

[原文出处](http://blog.cloudself.cn/2017/12/09/p-single-can-reduce-http-request/)
[Github](https://github.com/HerbLuo/p-single)