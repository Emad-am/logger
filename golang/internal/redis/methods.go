package redis

import (
	"encoding/json"
	"fmt"
	"logger/internal/context"
)

func Lpush(key string, value interface{}) int64 {
	ctx := context.GetContext()

	v, e := json.Marshal(value)
	if e != nil {
		fmt.Println(e)
	}

	res, err := Rdb.LPush(*ctx, key, v).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return res
}

func Rpop(key string) map[string]string {
	ctx := context.GetContext()
	res, err := Rdb.RPop(*ctx, key).Result()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var dat map[string]string

	_ = json.Unmarshal([]byte(res), &dat)

	return dat
}
