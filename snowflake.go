package goTools

import (
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

// ==================== 内部实现 ====================
var (
	globalNode *snowflake.Node
	once       sync.Once
	mu         sync.RWMutex
	customNode map[int64]*snowflake.Node // 支持多节点
)

func init() {
	customNode = make(map[int64]*snowflake.Node)
}

// ==================== 简单用户API ====================
// 只需一行代码初始化，后续直接调用 ID() 即可

// Setup 初始化（简单模式）
// nodeNum: 节点编号（0-1023）
// 使用示例: snowflake.Setup(1)
func Setup(nodeNum int64) {
	once.Do(func() {
		var err error
		globalNode, err = snowflake.NewNode(nodeNum)
		if err != nil {
			panic(fmt.Sprintf("雪花算法初始化失败: %v", err))
		}
	})
}

// ID 生成唯一ID（最简单用法）
// 使用示例: id := snowflake.ID()
func ID() int64 {
	if globalNode == nil {
		panic("请先调用 snowflake.Setup(nodeNum)")
	}
	return globalNode.Generate().Int64()
}

// IDStr 生成字符串格式的ID
func IDStr() string {
	return snowflake.ID(ID()).String()
}

// ==================== 高级用户API ====================
// 支持自定义配置、多节点、ID解析等

// AdvancedConfig 高级配置
type AdvancedConfig struct {
	Epoch    time.Time // 自定义起始时间
	NodeBits uint8     // 节点ID位数（默认10）
	StepBits uint8     // 序列号位数（默认12）
}

// DefaultConfig 默认配置（Twitter标准）
var DefaultConfig = AdvancedConfig{
	Epoch:    time.Date(2010, 11, 4, 1, 42, 54, 0, time.UTC),
	NodeBits: 10,
	StepBits: 12,
}

// InitAdvanced 初始化高级模式
// config: 配置参数，传nil则使用默认配置
// 使用示例: snowflake.InitAdvanced(nil)
func InitAdvanced(config *AdvancedConfig) error {
	if config == nil {
		config = &DefaultConfig
	}

	// 应用配置
	snowflake.Epoch = config.Epoch.UnixNano() / 1e6
	snowflake.NodeBits = config.NodeBits
	snowflake.StepBits = config.StepBits

	// 验证配置
	maxNode := int64(1<<config.NodeBits - 1)
	if maxNode < 0 {
		return fmt.Errorf("NodeBits配置无效: %d", config.NodeBits)
	}

	return nil
}

// NewNode 创建自定义节点（支持多节点）
// nodeNum: 节点编号（范围取决于NodeBits配置）
// 使用示例: node, err := snowflake.NewNode(1)
func NewNode(nodeNum int64) (*snowflake.Node, error) {
	return snowflake.NewNode(nodeNum)
}

// RegisterNode 注册并缓存节点（方便重复使用）
func RegisterNode(nodeNum int64) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := customNode[nodeNum]; exists {
		return nil // 已存在，直接返回
	}

	node, err := snowflake.NewNode(nodeNum)
	if err != nil {
		return err
	}

	customNode[nodeNum] = node
	return nil
}

// GenerateFromNode 从指定节点生成ID
func GenerateFromNode(nodeNum int64) (int64, error) {
	mu.RLock()
	node, exists := customNode[nodeNum]
	mu.RUnlock()

	if !exists {
		return 0, fmt.Errorf("节点 %d 未注册，请先调用 RegisterNode", nodeNum)
	}

	return node.Generate().Int64(), nil
}

// ParseID 解析ID，获取详细信息
// 返回: 时间戳(毫秒), 节点ID, 序列号, 生成时间
func ParseID(id int64) (timestamp int64, nodeID int64, step int64, generateTime time.Time, err error) {
	snowID := snowflake.ID(id)

	if snowID == 0 {
		err = fmt.Errorf("无效的ID: %d", id)
		return
	}

	timestamp = snowID.Time()
	nodeID = snowID.Node()
	step = snowID.Step()

	// 计算实际时间（需要考虑自定义epoch）
	epoch := snowflake.Epoch
	generateTime = time.Unix(epoch/1000, (epoch%1000)*1e6).Add(time.Duration(timestamp) * time.Millisecond)

	return
}

// ==================== 辅助函数 ====================
// BatchID 批量生成ID（简单模式）
func BatchID(count int) []int64 {
	if count <= 0 || count > 10000 {
		panic("批量数量必须在1-10000之间")
	}

	ids := make([]int64, count)
	for i := 0; i < count; i++ {
		ids[i] = ID()
	}
	return ids
}

// Benchmark 性能测试辅助函数
func Benchmark(n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		ID()
	}
	return time.Since(start)
}
