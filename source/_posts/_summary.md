
### CSS3

### 前端性能优化

前端性能优化的首要任务是什么？应该是处理网络io

### 执行顺序
process.nextTick和Promise的回调函数，追加在本轮循环
(Promise的构造函数的执行时间为立即执行)

setTimeout()和setInterval()

io

轮询

setImmediate()

### 其它
1. for in遍历的是对象（包括数组）的索引（即键名），而for of遍历的是数组元素值

2. for in，Object.keys，不包括像原生方法 Array.prototype.sort

1. service worker, 仅在https下才能生效，用途：离线、通知。

1. localStorage想存储更多数据：结合二级域名，使用IndexedDB

1. 对抽象语法树的理解，我的理解就是一种中间产物，目的是为了方便对代码进行检查，转换，压缩等一系列操作。

1. 浏览器渲染机制

1. 事件捕获，冒泡




