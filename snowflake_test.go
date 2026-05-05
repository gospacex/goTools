package goTools

import (
	"fmt"
	"testing"
)

/**
* 小白用法
* 测试获取Snowflake ID
* @param t 测试框架的测试对象
 */
func TestGetSnowflakeId(t *testing.T) {
	// 第1步：初始化（只需一次）
	Setup(1) // 使用节点1

	// 第2步：愉快地生成ID
	id := ID()
	fmt.Printf("ID: %d\n", id)

	// 字符串格式
	idStr := IDStr()
	fmt.Printf("ID String: %s\n", idStr)

	// 批量生成
	ids := BatchID(10)
	fmt.Printf("10个ID: %v\n", ids)

	// 性能测试
	elapsed := Benchmark(100000)
	fmt.Printf("生成10万ID耗时: %v\n", elapsed)
}
