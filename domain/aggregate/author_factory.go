package aggregate

import "github.com/huimingz/ddd-go/domain/entity"

// 工程函数用于创建聚合根，封装了聚合根的创建逻辑。
// 创建过程依赖外部的领域服务（例如调用外部服务生成ID等）来自于 domain/service 包。

func NewAuthor(userName string) *Author {
	return &Author{
		user: &entity.User{
			Username: userName,
		},
	}
}
