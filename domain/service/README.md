# 领域服务

不适合放到聚合根的业务逻辑，比如需要调用外部服务生成实体 ID ，又或者是调用 TLM 服务申请租户的资源位置。

ApplicationService 和 DomainService 是两个很不一样的概念，前者是必须有的 DDD 组件，而后者只是一种妥协的结果，程序中的 DomainService
应该越少越好。
