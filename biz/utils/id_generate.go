package utils

import "github.com/bwmarrin/snowflake"

var (
	userIDGenerator *snowflake.Node
)

func init() {
	var err error
	userIDGenerator, err = snowflake.NewNode(100)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() uint64 {
	return uint64(userIDGenerator.Generate().Int64())
}
