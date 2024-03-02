package utils

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

// InitSnowflake 初始化生成器对象
func InitSnowflake(datacenterId int64, workerId int64) (*snowflake.Snowflake, error) {
	s, err := snowflake.NewSnowflake(datacenterId, workerId)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetNextSnowflakeId 生成全局唯一的自增uid
func GetNextSnowflakeId(s *snowflake.Snowflake) int64 {
	return s.NextVal()
}
