package example

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ExecRedis() {
	var ctx = context.Background()
	// 连接 Redis 数据库
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// 测试连接 Redis
	var ping, _ = client.Ping(ctx).Result()
	fmt.Printf("测试链接 Redis：%v\n", ping)
	defer client.Close()

	// string：字符串
	// 参数 ctx 是内置包 context 创建的上下文
	// 参数 expiration 是有效期，数据类型为 time.Duration
	// var strSet, _ = client.Set(ctx, "name", "Tom", time.Hour).Result()
	// fmt.Println("设置字符串类型的数据：", strSet)
	// // 获取字符串类型的数据
	// var strGet, _ = client.Get(ctx, "name").Result()
	// fmt.Println("获取字符串类型的数据", strGet)
	// // 删除字符串类型的数据
	// var strDel, _ = client.Del(ctx, "name").Result()
	// fmt.Println("删除字符串类型的数据", strDel)

	// hash：哈希
	// values 是不固定参数，参数类型为空接口，代表哈希数值
	// var hashSet, _ = client.HSet(ctx, "Tom", "age", 10).Result()
	// fmt.Println("设置哈希类型的数据：", hashSet)
	// // 获取哈希类型的数据
	// var hashGet, _ = client.HGet(ctx, "Tom", "age").Result()
	// fmt.Println("获取哈希类型的数据", hashGet)
	// // 删除哈希类型的数据
	// var hashDel, _ = client.HDel(ctx, "Tom", "age").Result()
	// fmt.Println("删除哈希类型的数据", hashDel)

	// list：列表
	// var litRPush, _ = client.RPush(ctx, "Tom", "English", "Chinese").Result()
	// fmt.Println("设置列表的数据：", litRPush)
	// 获取列表指定范围的元素
	// 参数 start 和 stop 是列表索引，数据类型为整型
	// litRange, _ := client.LRange(ctx, "Tom", 0, 2).Result()
	// fmt.Println("获取列表指定范围内的元素：", litRange)
	// 移出并获取列表的第一个元素
	// 参数 timeout 设置超时，数据类型为 time.Duration
	// litBLPop, _ := client.BLPop(ctx, time.Second, "Tom").Result()
	// fmt.Println("移出并获取列表的第一个元素：", litBLPop)

	// set / map：集合
	// var setsadd, _ = client.SAdd(ctx, "Tim", 20, "180").Result()
	// fmt.Println("向集合添加一个或多个成员：", setsadd)
	// 获取集合中的所有成员
	// setsMembers, _ := client.SMembers(ctx, "Tim").Result()
	// fmt.Println("获取集合中的所有成员：", setsMembers)
	// // 移除并返回集合中的第一个元素
	// setsPop, _ := client.SPop(ctx, "Tim").Result()
	// fmt.Println("移除并返回集合中的第一个元素：", setsPop)

	// zset：有序集合
	// var z1 = redis.Z{Member: "190", Score: 5}
	// var z2 = redis.Z{Member: 10, Score: 10}
	// zsetzAdd, _ := client.ZAdd(ctx, "rank", &z1, &z2).Result()
	// fmt.Println("有序集合添加或更新一个或多个成员的分数：", zsetzAdd)
	// 通过索引区间返回有序集合指定区间内的成员
	// zsetzRange, _ := client.ZRange(ctx, "Tim", 0, 2).Result()
	// fmt.Println("通过索引区间返回有序集合指定区间内的成员：", zsetzRange)
	// 移除有序集合中的一个或多个成员
	// zsetzRem, _ := client.ZRem(ctx, "rank", "190").Result()
	// fmt.Println("移除有序集合中的一个或多个成员：", zsetzRem)

	// // stream：流类型数据
	// var x1 = redis.XAddArgs{
	// 	Stream: "Jesse",
	// 	Values: map[string]any{"age": 10, "height": "160CM"},
	// }
	// streXAdd, _ := client.XAdd(ctx, &x1).Result()
	// fmt.Println("新增 stream：", streXAdd)
	// // 获取 stream
	// // - 和 + 代表范围
	// streXRange, _ := client.XRange(ctx, "Jesse", "-", "+").Result()
	// fmt.Println("获取 stream 所有数据：", streXRange)
	// for _, v := range streXRange {
	// 	fmt.Printf("stream 数据的ID：%v，数据值：%v\n", v.ID, v.Values)
	// 	// 通过 stream ID 删除数据
	// 	strexDel, _ := client.XDel(ctx, "Jesse", v.ID).Result()
	// 	fmt.Printf("ID：%v 已删除，数据量：%v\n", v.ID, strexDel)
	// }

	// bit：二进制
	// 参数 values 因为是二进制只有 0 和 1
	// offset 是二进制的位数偏移量，0表示从左边第一位算起
	// client.Set(ctx, "en", "ABCDEFGHIJKLMNRPQLSTUVWXYZ", 0)
	// bitSetBit, _ := client.SetBit(ctx, "en", 0, 1).Result()
	// fmt.Println("bit 数据类型：", bitSetBit)
	// var bitGetBit, _ = client.GetBit(ctx, "en", 0).Result()
	// fmt.Println("获取 bit 数据类型的值：", bitGetBit)
	// var bitDel, _ = client.Del(ctx, "en").Result()
	// fmt.Println("删除 bit 数据类型的值：", bitDel)

}
