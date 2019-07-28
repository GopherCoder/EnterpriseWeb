# Lottery

> 小程序：抽奖助手


## 抽象实体

- 用户账号
- 抽奖
    - 发起抽奖
        - 不同类型
        - 奖项
            - 名称
            - 份额
            - 级别：默认一等
        - 开奖时间（5天之内，创建时间节点之后）
        - 开奖条件
            - 类型
            - 条件
                - 按时间自动开奖
                - 按人数自动开奖
                - 即开即中
- 查询
    - 发起
    - 参与        
- 首页：推荐的抽奖

## GraphQL

> 接口查询语言


- schema: 定义对象和操作动作（查询、更改）
- type: 类型（内置、自定义对象类型）

```graphql

{
    query{
        A: user(id:1){
            email
            id
        }
    }
}

```

- 字段(id, email)
```graphql
type User {
    id: ID
    email: String!
}
```
- 参数(id:1)

```graphql

query {
    user(id: 1){
        id
    }
}

```
- 别名(A:user)

```graphql
query{
    A:user(id:1){
        id
        email
    }
}

```
- 片段(fragment)
```graphql

query{
    user(id:1){
        ...Common
    }
}

fragment Common on User{
    id
    email
}
```
- 操作名称(操作类型（query, mutation, subscription）+操作名称(自定义), 操作名称可省略)

```graphql

query FindOneUser{
    user(id:1){
        id
        email
    }
}
```
- 枚举

```graphql

enum E {
    First
    Second
    Thired
}
```

- 修饰符:(!: 非空)

```graphql

type Name {
    name: String!
}

```

- 接口

```graphql
interface A {
    id: ID!
}

type B implements A {
    id:ID!
    name: String!
}
```

- 联合

```graphql

union  Search = A | B | C 
```

## Graph 开发指南

- schema 定义
    - query
    - mutation

- 定义 type (同时定义 model)
- 开发：定义 schema、type（field, args, Resolve)
- 字段定义命名格式
    - schema type: 驼峰式`adminId` --> Fields
    - 参数：驼峰式 --> args
    - json 序列化，依旧驼峰式

 

## 问题

- 接口测试发现，如何将请求体正确的解析再传入给 graphql 服务解析，是第一步

