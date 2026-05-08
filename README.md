# 📦 GoTools

> 一个功能强大、易于使用的 Go 语言通用工具库

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/your-repo)

<div align="center">

**简洁高效的 Go 工具集，让开发更简单**

[快速开始](#快速开始) • [功能特性](#功能特性) • [API 文档](#api-文档) • [示例代码](#示例代码) • [贡献指南](#贡献指南)

</div>

---

## 📖 项目简介

GoTools 是一个精心设计的 Go 语言通用工具库，旨在为开发者提供一套完整、高效、易用的工具集合。无论你是正在构建微服务、CLI 工具还是 Web 应用，GoTools 都能帮助你大幅提升开发效率，减少重复代码。

### ✨ 设计理念

- **零依赖**：核心模块无外部依赖，开箱即用
- **高性能**：经过优化的算法和数据结构
- **类型安全**：完整的泛型支持，编译期类型检查
- **生产就绪**：经过充分测试，稳定可靠

---

## 🚀 快速开始

### 安装

```bash
go get github.com/your-org/goTools
```

### 基础使用

```go
package main

import (
    "fmt"
    "github.com/your-org/goTools"
)

func main() {
    // 使用字符串工具
    str := "  Go Tools  "
    trimmed := goTools.StrTrim(str)
    fmt.Println(trimmed) // 输出: "Go Tools"

    // 使用集合工具
    slice := []int{1, 2, 3, 4, 5}
    filtered := goTools.SliceFilter(slice, func(v int) bool {
        return v%2 == 0
    })
    fmt.Println(filtered) // 输出: [2 4]
}
```

---

## 📚 功能特性

### 🔤 字符串工具 (str)

提供常用的字符串操作函数：

| 函数 | 说明 | 示例 |
|------|------|------|
| `StrTrim` | 去除首尾空格 | `"  hello  "` → `"hello"` |
| `StrToLower` | 转换为小写 | `"Hello"` → `"hello"` |
| `StrToUpper` | 转换为大写 | `"hello"` → `"HELLO"` |
| `StrContains` | 判断包含 | `"hello"` 包含 `"ell"` |
| `StrJoin` | 连接字符串切片 | `["a","b","c"]` → `"a,b,c"` |
| `StrSplit` | 分割字符串 | `"a,b,c"` → `["a","b","c"]` |
| `StrReplace` | 字符串替换 | `"hello"` 替换 `"l"` → `"hexxo"` |
| `StrSub` | 子字符串 | `"hello"`[1:4] → `"ell"` |
| `StrPad` | 字符串填充 | `"5"` 填充到 3 位 → `"005"` |

### 📊 集合工具 (slice)

强大的切片操作函数：

| 函数 | 说明 | 示例 |
|------|------|------|
| `SliceFilter` | 过滤元素 | 过滤奇数 |
| `SliceMap` | 映射转换 | 每个元素 ×2 |
| `SliceReduce` | 归约聚合 | 求和/求积 |
| `SliceFind` | 查找元素 | 返回第一个匹配 |
| `SliceContains` | 判断包含 | 是否存在某值 |
| `SliceUnique` | 去重 | 移除重复元素 |
| `SliceReverse` | 反转顺序 | [1,2,3] → [3,2,1] |
| `SliceChunk` | 分块 | 每 N 个一组 |
| `SliceSort` | 排序 | 升序/降序 |

### 🔢 数字工具 (num)

数值处理与转换：

| 函数 | 说明 | 示例 |
|------|------|------|
| `NumAbs` | 绝对值 | `-5` → `5` |
| `NumMax/Min` | 最大/最小值 | `Max(1,5)` → `5` |
| `NumClamp` | 限制范围 | 限制在 [0,100] |
| `NumRound` | 四舍五入 | `3.6` → `4` |
| `NumPow` | 幂运算 | `2^3` → `8` |
| `NumIsPrime` | 判断素数 | `7` → `true` |
| `NumGCD/LCM` | 最大公约数/最小公倍数 | `GCD(12,18)` → `6` |

### 📅 时间工具 (time)

日期时间操作：

| 函数 | 说明 | 示例 |
|------|------|------|
| `TimeNow` | 当前时间 | `time.Now()` |
| `TimeFormat` | 格式化 | `"2006-01-02 15:04:05"` |
| `TimeParse` | 解析时间字符串 | 解析自定义格式 |
| `TimeAdd` | 时间加减 | `+1 day`, `-2 hours` |
| `TimeDiff` | 时间差 | 相差多少天/时/分/秒 |
| `TimeStartOfDay` | 当天开始 | `00:00:00` |
| `TimeEndOfDay` | 当天结束 | `23:59:59` |
| `TimeIsLeapYear` | 判断闰年 | `2024` → `true` |

### 🔐 加密工具 (crypto)

安全相关的加密函数：

| 函数 | 说明 |
|------|------|
| `CryptoMD5` | MD5 哈希 |
| `CryptoSHA1/SHA256` | SHA 哈希 |
| `CryptoHash` | 通用哈希（支持多种算法） |
| `CryptoBase64Encode/Decode` | Base64 编解码 |
| `CryptoAESEncrypt/Decrypt` | AES 加密解密 |
| `CryptoRandomString` | 随机字符串生成 |
| `CryptoRandomInt` | 随机整数生成 |

### 📝 验证工具 (validate)

数据验证函数：

| 函数 | 说明 |
|------|------|
| `ValidateEmail` | 邮箱格式验证 |
| `ValidatePhone` | 手机号验证 |
| `ValidateURL` | URL 格式验证 |
| `ValidateIP` | IP 地址验证 |
| `ValidateChinese` | 中文字符验证 |
| `ValidateIDCard` | 身份证号验证（支持18位） |

---

## 📁 项目结构

```
goTools/
├── README.md           # 项目说明文档
├── LICENSE             # 开源许可证
├── go.mod             # Go 模块定义
├── go.sum             # 依赖校验和
│
├── str/               # 字符串工具包
│   ├── str.go         # 核心字符串函数
│   ├── str_test.go    # 单元测试
│   └── examples/      # 使用示例
│
├── slice/             # 集合工具包
│   ├── slice.go       # 核心切片函数
│   ├── slice_test.go  # 单元测试
│   └── examples/      # 使用示例
│
├── num/               # 数字工具包
│   ├── num.go         # 核心数字函数
│   ├── num_test.go    # 单元测试
│   └── examples/      # 使用示例
│
├── time/              # 时间工具包
│   ├── time.go        # 核心时间函数
│   ├── time_test.go   # 单元测试
│   └── examples/      # 使用示例
│
├── crypto/            # 加密工具包
│   ├── crypto.go      # 核心加密函数
│   ├── crypto_test.go # 单元测试
│   └── examples/      # 使用示例
│
└── validate/          # 验证工具包
    ├── validate.go    # 核心验证函数
    ├── validate_test.go # 单元测试
    └── examples/      # 使用示例
```

---

## 💻 API 文档

### 字符串工具 (str)

```go
// StrTrim 去除字符串首尾空白字符
func StrTrim(s string) string

// StrToLower 转换为小写
func StrToLower(s string) string

// StrToUpper 转换为大写
func StrToUpper(s string) string

// StrContains 判断字符串是否包含子串
func StrContains(s, substr string) bool

// StrJoin 使用分隔符连接字符串切片
func StrJoin(slice []string, sep string) string

// StrSplit 按分隔符分割字符串
func StrSplit(s, sep string) []string

// StrReplace 替换字符串中的子串
func StrReplace(s, old, new string, n int) string

// StrSub 返回子字符串 [start, end)
func StrSub(s string, start, end int) string

// StrPad 用字符填充字符串到指定长度
func StrPad(s string, length int, padChar rune) string

// StrPadLeft 左侧填充
func StrPadLeft(s string, length int, padChar rune) string

// StrPadRight 右侧填充
func StrPadRight(s string, length int, padChar rune) string
```

### 集合工具 (slice)

```go
// SliceFilter 过滤切片元素
func SliceFilter[T any](slice []T, fn func(T) bool) []T

// SliceMap 映射转换切片
func SliceMap[T, R any](slice []T, fn func(T) R) []R

// SliceReduce 归约聚合切片
func SliceReduce[T any](slice []T, initial T, fn func(T, T) T) T

// SliceFind 查找第一个匹配的元素
func SliceFind[T any](slice []T, fn func(T) bool) (T, bool)

// SliceContains 判断切片是否包含元素
func SliceContains[T comparable](slice []T, elem T) bool

// SliceUnique 去重
func SliceUnique[T comparable](slice []T) []T

// SliceReverse 反转切片顺序
func SliceReverse[T any](slice []T) []T

// SliceChunk 分块
func SliceChunk[T any](slice []T, size int) [][]T

// SliceSort 排序
func SliceSort[T any](slice []T, less func(T, T) bool) []T

// SliceShuffle 随机打乱
func SliceShuffle[T any](slice []T) []T
```

### 数字工具 (num)

```go
// NumAbs 绝对值
func NumAbs(x int) int

// NumMax 最大值
func NumMax(a, b int) int

// NumMin 最小值
func NumMin(a, b int) int

// NumClamp 限制范围
func NumClamp(x, min, max int) int

// NumRound 四舍五入
func NumRound(x float64) int

// NumCeil 向上取整
func NumCeil(x float64) int

// NumFloor 向下取整
func NumFloor(x float64) int

// NumPow 幂运算
func NumPow(base, exp int) int

// NumSqrt 平方根
func NumSqrt(x float64) float64

// NumIsPrime 判断素数
func NumIsPrime(n int) bool

// NumGCD 最大公约数
func NumGCD(a, b int) int

// NumLCM 最小公倍数
func NumLCM(a, b int) int

// NumFactorial 阶乘
func NumFactorial(n int) int

// NumFibonacci 斐波那契数列第N项
func NumFibonacci(n int) int
```

### 时间工具 (time)

```go
// TimeNow 获取当前时间
func TimeNow() time.Time

// TimeFormat 格式化时间
func TimeFormat(t time.Time, layout string) string

// TimeParse 解析时间字符串
func TimeParse(value, layout string) (time.Time, error)

// TimeAdd 时间加减
func TimeAdd(t time.Time, d time.Duration) time.Time

// TimeDiff 计算时间差
func TimeDiff(from, to time.Time) time.Duration

// TimeStartOfDay 获取当天开始时间
func TimeStartOfDay(t time.Time) time.Time

// TimeEndOfDay 获取当天结束时间
func TimeEndOfDay(t time.Time) time.Time

// TimeIsLeapYear 判断是否闰年
func TimeIsLeapYear(year int) bool

// TimeDaysInMonth 获取月份天数
func TimeDaysInMonth(year, month int) int

// TimeWeekday 获取星期几
func TimeWeekday(t time.Time) time.Weekday

// TimeUnix 转换为 Unix 时间戳
func TimeUnix(t time.Time) int64

// TimeFromUnix 从时间戳创建时间
func TimeFromUnix(ts int64) time.Time
```

### 加密工具 (crypto)

```go
// CryptoMD5 MD5 哈希
func CryptoMD5(data []byte) []byte

// CryptoSHA1 SHA1 哈希
func CryptoSHA1(data []byte) []byte

// CryptoSHA256 SHA256 哈希
func CryptoSHA256(data []byte) []byte

// CryptoHash 通用哈希函数
func CryptoHash(data []byte, algo string) ([]byte, error)

// CryptoBase64Encode Base64 编码
func CryptoBase64Encode(data []byte) string

// CryptoBase64Decode Base64 解码
func CryptoBase64Decode(s string) ([]byte, error)

// CryptoAESEncrypt AES 加密
func CryptoAESEncrypt(plaintext, key []byte) ([]byte, error)

// CryptoAESDecrypt AES 解密
func CryptoAESDecrypt(ciphertext, key []byte) ([]byte, error)

// CryptoRandomString 生成随机字符串
func CryptoRandomString(length int) (string, error)

// CryptoRandomInt 生成随机整数
func CryptoRandomInt(min, max int) int
```

### 验证工具 (validate)

```go
// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool

// ValidatePhone 验证手机号（中国大陆）
func ValidatePhone(phone string) bool

// ValidateURL 验证 URL 格式
func ValidateURL(url string) bool

// ValidateIP 验证 IP 地址
func ValidateIP(ip string) bool

// ValidateIPv4 验证 IPv4
func ValidateIPv4(ip string) bool

// ValidateIPv6 验证 IPv6
func ValidateIPv6(ip string) bool

// ValidateChinese 验证是否包含中文
func ValidateChinese(s string) bool

// ValidateIDCard 验证身份证号（18位）
func ValidateIDCard(id string) bool

// ValidatePassport 验证护照号码
func ValidatePassport(passport string) bool

// ValidateCreditCard 验证信用卡号
func ValidateCreditCard(card string) bool
```

---

## 🎯 示例代码

### 完整示例：用户数据处理

```go
package main

import (
    "fmt"
    "log"
    "github.com/your-org/goTools"
    "github.com/your-org/goTools/str"
    "github.com/your-org/goTools/slice"
    "github.com/your-org/goTools/validate"
)

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func main() {
    users := []User{
        {1, "Alice", "alice@example.com", 25},
        {2, "Bob", "bob@example.com", 17},
        {3, "Charlie", "charlie@example.com", 30},
        {4, "David", "david@example.com", 15},
        {5, "Eve", "eve@example.com", 22},
    }

    // 1. 过滤成年用户 (>= 18岁)
    adults := slice.Filter(users, func(u User) bool {
        return u.Age >= 18
    })
    fmt.Printf("成年用户: %d 人\n", len(adults))

    // 2. 映射用户姓名（转大写并去空格）
    names := slice.Map(adults, func(u User) string {
        return str.ToUpper(str.Trim(u.Name))
    })
    fmt.Println("用户姓名:", names)

    // 3. 统计总年龄
    totalAge := slice.Reduce(adults, 0, func(acc int, u User) int {
        return acc + u.Age
    })
    avgAge := float64(totalAge) / float64(len(adults))
    fmt.Printf("平均年龄: %.1f\n", avgAge)

    // 4. 验证邮箱格式
    for _, u := range adults {
        if validate.Email(u.Email) {
            fmt.Printf("✅ %s 的邮箱格式正确\n", u.Name)
        } else {
            log.Printf("❌ %s 的邮箱格式错误: %s\n", u.Name, u.Email)
        }
    }

    // 5. 去重用户名
    uniqueNames := slice.Unique(names)
    fmt.Println("唯一姓名:", uniqueNames)

    // 6. 按年龄排序
    sorted := slice.Sort(adults, func(a, b User) bool {
        return a.Age < b.Age
    })
    fmt.Println("按年龄排序:")
    for _, u := range sorted {
        fmt.Printf("  %s: %d岁\n", u.Name, u.Age)
    }
}
```

### 时间处理示例

```go
package main

import (
    "fmt"
    "time"
    "github.com/your-org/goTools/timeutil"
)

func main() {
    now := timeutil.Now()

    // 格式化
    fmt.Println("当前时间:", timeutil.Format(now, "2006-01-02 15:04:05"))

    // 时间计算
    tomorrow := timeutil.Add(now, 24*time.Hour)
    fmt.Println("明天:", timeutil.Format(tomorrow, "2006-01-02"))

    // 判断是否闰年
    year := now.Year()
    fmt.Printf("%d年是否闰年: %v\n", year, timeutil.IsLeapYear(year))

    // 获取季度
    quarter := (now.Month()-1)/3 + 1
    fmt.Printf("当前季度: Q%d\n", quarter)

    // 时间差
    past := timeutil.Add(now, -7*24*time.Hour)
    diff := timeutil.Diff(past, now)
    fmt.Printf("一周前到现在: %v\n", diff)
}
```

---

## 🧪 测试

### 运行所有测试

```bash
# 运行所有包的测试
go test ./...

# 运行特定包的测试
go test ./str/...
go test ./slice/...
go test ./num/...

# 生成测试覆盖率报告
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### 基准测试

```bash
# 运行基准测试
go test -bench=. ./...

# 查看内存分配
go test -bench=. -benchmem ./...
```

---

## 📦 依赖管理

### 核心依赖（无外部依赖）

GoTools 核心模块**零依赖**，直接 `import` 即可使用。

### 可选依赖

某些高级功能可能需要以下依赖：

```bash
# 如果需要更强大的正则表达式支持
go get github.com/dlclark/regexp2

# 如果需要国际化支持
go get golang.org/x/text

# 如果需要 JSON 增强
go get github.com/json-iterator/go
```

---

## 🤝 贡献指南

我们非常欢迎社区贡献！请遵循以下流程：

### 1. Fork 项目

```bash
# Fork 后在本地添加 upstream
git remote add upstream https://github.com/your-org/goTools.git
```

### 2. 创建特性分支

```bash
git checkout -b feature/amazing-feature
```

### 3. 提交更改

```bash
git add .
git commit -m "feat: add amazing feature"
```

**提交信息规范**（遵循 [Conventional Commits](https://www.conventionalcommits.org/)）：

- `feat:` 新功能
- `fix:` Bug 修复
- `docs:` 文档更新
- `style:` 代码格式调整
- `refactor:` 代码重构
- `test:` 测试相关
- `chore:` 构建/工具变更

### 4. 推送到分支

```bash
git push origin feature/amazing-feature
```

### 5. 提交 Pull Request

- 确保所有测试通过
- 添加必要的测试用例
- 更新相关文档
- 填写 PR 模板

### 代码规范

- 使用 `go fmt` 格式化代码
- 遵循 [Effective Go](https://golang.org/doc/effective_go) 指南
- 添加适当的注释和文档
- 确保测试覆盖率不低于 80%

---

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

---

## 🙏 致谢

感谢所有为 GoTools 做出贡献的开发者：

<a href="https://github.com/your-org/goTools/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=your-org/goTools" />
</a>

---

## 📞 联系方式

- **项目主页**: https://github.com/your-org/goTools
- **问题反馈**: https://github.com/your-org/goTools/issues
- **邮箱**: your-email@example.com

---

<div align="center">

**如果这个项目对你有帮助，请给我们一个 ⭐ Star！**

Made with ❤️ by GoTools Team

</div>
