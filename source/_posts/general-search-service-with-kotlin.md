---
title: 后端 之 构建一个通用的搜索服务
date: 2017-12-20 01:03:40
tags: [kotlin, spring data, search service]
---

本文讲述的是 Kotlin + Spring ORM with JPA 实现的一个通用的搜索服务
如果您是Java开发者且还未使用过Kotlin，建议您试试Kotlin，相比起Java，Kotlin更灵活，且顺手。

<!-- more -->

__注：本文所需的代码文件见最后一小节__

### 简介

kotlin实现的搜索服务，可配置允许搜索的字段。

如果您使用的不是Spring MVC + Spring ORM with JPA（_**注1**_），那么您需要额外修改源代码

如果您使用的不是Kotlin，但用的IDE是IDEA，您只需点击 菜单 -> Tools -> Kotlin -> Configure Kotlin in Project，这样，就可以混写Java与Kotlin

- - -
注1：
实质上，该搜索服务基于EntityManager，

您只需要在Spring application context中创建一个EntityManagerFactory，该服务便可正常使用，通常，创建一个它方法，是配置一个Bean:`org.springframework.orm.jpa.LocalContainerEntityManagerFactoryBean`配置方法可参考附件（spring-orm-jpa.xml）和（Demo1）
   
如果您使用Spring Boot，直接导入 `spring-boot-starter-data-jpa`，并配置好用户名密码，即可（可参考Demo2）
- - -

### 示例

直接搜索

```
// 名字中包含g
http://localhost:8080/search-service/user/username/g?size=5&page=0
```

外键搜索

```
// 搜索店铺名包含"flagship"的店铺中的所有商品
.../search-service/item/shop.name/flagship?size=5&page=0
```

范围搜索
```
// 价格10到99http://localhost:8080
.../search-service/item/price/[~10~99]?size=5&page=0
```

同时搜索多个
```
// 用户所在组为 1或2或6 的
.../search-service/user/group/[,1,2,6]?size=5&page=0
```

### 使用方法

1. 导入Maven依赖

    如果使用Spring Boot，Maven依赖如下
    
    ```xml 
    <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-data-jpa</artifactId>
    </dependency>
    
    <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    
    <dependency>
      <groupId>mysql</groupId>
      <artifactId>mysql-connector-java</artifactId>
      <scope>runtime</scope>
    </dependency>
    ```
    否则，Maven依赖如下
    ```xml 
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-webmvc</artifactId>
        <version>4.3.13.RELEASE</version>
    </dependency>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-aspects</artifactId>
        <version>4.3.13.RELEASE</version>
    </dependency>
    <dependency>
        <groupId>org.springframework</groupId>
        <artifactId>spring-orm</artifactId>
        <version>4.3.13.RELEASE</version>
    </dependency>

    <dependency>
        <groupId>org.hibernate</groupId>
        <artifactId>hibernate-entitymanager</artifactId>
        <version>5.2.12.Final</version>
    </dependency>
    <dependency>
        <groupId>mysql</groupId>
        <artifactId>mysql-connector-java</artifactId>
        <version>6.0.6</version>
    </dependency>
    <dependency>
        <groupId>com.mchange</groupId>
        <artifactId>c3p0</artifactId>
        <version>0.9.5.2</version>
    </dependency>

    <dependency>
        <groupId>org.springframework.data</groupId>
        <artifactId>spring-data-commons</artifactId>
        <version>2.0.2.RELEASE</version>
    </dependency>
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
        <version>2.5.1</version>
    </dependency>
    ```
2. 下载源代码（下载链接位于文章末尾）

3. 将各文件按下图所示放置
![](/images/85f66fbbda45880a12e59e27462136ecaae32f45.png)

4. 配置 `searchConfig.json`, 其中

    entityName为允许搜索的实体类名
    properties为允许搜索的属性名
    type的值可以是 "int|boolean|double|string|timestamp"，默认是"string"
    
    例：
    ```json
    {
      "version": 1,
      "search": [{
        "entityName": "User",
        "properties": [
          {"name": "nickname"},
          {"name": "username"},
          {"name": "group", "type": "int"},
          {"name": "enable", "type": "boolean"}
        ]
      }, {
        "entityName": "entityName",
        "properties": [
          {"name": "propertyName", "type": "int|boolean|double|string|timestamp"}
        ]
      }]
    }
    ```
5. 编写 Controller层代码如下
    ```java
    import org.springframework.beans.factory.annotation.Autowired;
    import org.springframework.data.domain.PageRequest;
    import org.springframework.web.bind.annotation.*;
    
    @RestController
    @RequestMapping("/search-service")
    public class SearchController {
    
        private final ISearchService searchService;
    
        @Autowired
        public SearchController(ISearchService searchService) {
            this.searchService = searchService;
        }
    
        @GetMapping("/{entityName}/{property}/{searchText}/")
        public Iterable<?> search(@PathVariable String entityName,
                                  @PathVariable String property,
                                  @PathVariable String searchText,
                                  @RequestParam Integer page,
                                  @RequestParam Integer size) throws Exception {
            return searchService.search(
                    entityName,
                    property,
                    searchText,
                    PageRequest.of(page, size)
            );
        }
    }
    ```

6. 尝试运行它

### 搜索规则

如果您的Controller代码与上述一致，那么搜索的请求地址为
`http://localhost:8080/search-service/{entityName}/{property}/{searchText}/?page=0&size=10`

其中entityName为实体类名（区分大小写），property为属性名，searchText为搜索的内容

searchText具体可为
1. 搜索的内容，如`apple`可匹配`apple`,`an-apple`,`apple-vinegar`等
    搜索的类型会自动转换为配置文件中配置的类型，
    int: `9`等
    boolean: `true`等
    double: `9.9`等
    timestamp: `1514764800000`等
2. 数组，如 `[,1,2]` 可匹配 `1`, `2`等
3. 范围，如 `[~1~100]` 可匹配 `1`, `2`, `99.9`等 但不匹配 `100`

### 其它

如果您无法运行该实例，可参考下载附件里的 Demo并作参考
运行Demo前，需先导入数据库并配置好用户名和密码，详细见根目录下的README.md
[原文出处](http://blog.cloudself.cn/2017/12/20/general-search-service-with-kotlin/)

### 附件
[源代码](/download/general-search-service-with-kotlin/source-code-v103.zip)
[spring-orm-jpa.xml](/download/general-search-service-with-kotlin/spring-orm-jpa.xml)
[Demo1(Maven版)](/download/general-search-service-with-kotlin/demo-search-service.zip)
[Demo2(Spring Boot版)](/download/general-search-service-with-kotlin/demo-search-service-with-spring-boot.zip)