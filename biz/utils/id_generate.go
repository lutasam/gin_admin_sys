package utils

import "github.com/bwmarrin/snowflake"

var (
	userIDGenerator      *snowflake.Node
	commodityIDGenerator *snowflake.Node
)

func init() {
	var err error
	userIDGenerator, err = snowflake.NewNode(100)
	if err != nil {
		panic(err)
	}

	commodityIDGenerator, err = snowflake.NewNode(200)
	if err != nil {
		panic(err)
	}
}

func GenerateUserID() uint64 {
	return uint64(userIDGenerator.Generate().Int64())
}

func GenerateCommodityID() uint64 {
	return uint64(commodityIDGenerator.Generate().Int64())
}
