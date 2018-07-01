///**
// * 简介：
// *     该任务为批量处理人员信息，因为妇女节到了，公司决定给公司的女性员工发放1000元奖金，在月
// *  末结算工资时发放，现在需要在每个女性员工的数据中加上这一个奖励，写一个程序按照要求完成它。
// *  要求：将下面的函数体补充完整，使用通道和 goroutine 完成它，而且需要可以控制 goroutine 的数量。
// */
//
package main
//
//import (
//	"fmt"
//	"time"
//	"math/rand"
//)
//
//const (
//	male = 0
//	female = 1
//
//	// 妇女节奖励的 Item 的 名字
//	bonusName = "Women's Day bonus"
//)
//
//// 定义员工数据结构
//type Person struct {
//	Name   string
//	Sex    int
//	Reward []*Item
//}
//
//// 定义地址数据结构
//type Item struct {
//	name   string
//	amount int
//}
//
//// 定义处理接口
//type PersonHandler interface {
//	Batch(origs <-chan Person) <-chan Person
//	Handle(orig *Person)
//}
//
//// 定义空结构体，为其添加方法，实现 PersonHandler 接口
//type PersonHandlerImpl struct{}
//
//// 方法 Batch 被声明为实现批量处理人员信息功能的方法，
//// 其方法声明中的两个通道分别对该方法和该方法的调用方使用它的方式进行了约束
//func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
//
//}
//
//// 具体处理操作
//func (handler PersonHandlerImpl) Handle(orig *Person) {
//
//}
//
////定义要被处理的数据并初始化
//var personTotal = 200
//
//// 未被处理的员工
//var iniPersons = make([]Person, personTotal)
//
//// 将处理过的员工存放到这里
//var desPersons = make([]Person, personTotal)
//
//var personCount int
//
//func init() {
//	// 随机生成 200 名性别不完全相同的员工
//	for i := 0; i < 200; i++ {
//		name := fmt.Sprintf("%s%d", "P", i)
//		r := rand.New(rand.NewSource(time.Now().UnixNano()))
//		p := Person{name, r.Intn(2), nil}
//		iniPersons[i] = p
//	}
//}
//
//func main() {
//	// 首先获取 handler
//	handler := getPersonHandler()
//	// 初始化 origs 通道
//	origs := make(chan Person, 100)
//
//	// 将人员信息通过 origs 通道传入 Batch 中处理，处理后的信息放入 dests 通道中，
//	// 并将 dests 通道返回，进行下一步处理。
//	dests := handler.Batch(origs)
//
//	// fecthPerson 获取人员信息放入到 origs 中
//	fecthPerson(origs)
//
//	// savePerson 从 dests 中接收处理过的信息进行保存
//	sign := savePerson(dests)
//
//	// sign 通道作用为在批处理完全执行结束之前阻塞主 Goroutine
//	<-sign
//}
//
//func getPersonHandler() PersonHandler {
//	return PersonHandlerImpl{}
//}
//
//func savePerson(dest <-chan Person) <-chan byte {
//
//}
//
//func fecthPerson(origs chan<- Person) {
//
//}
//
//// goTicket 是为了限制该程序启用的 goroutine 的数量而声明的一个缓冲通道
//// 根据传进来的 total 初始化通道，total 即表示可以启用 goroutine 数量
//// 每当启用一个 goroutine 时从该通道中接受一个值表示可用 goroutine 少了一个
//// 即每个 goroutine 要想启动必须要有 ticket。上述是在 origs 为缓冲条件下，即整个过程为异步完成情况下
//func initGoTicket(total int) chan byte {
//	var goTicket chan byte
//	if total == 0 {
//		return goTicket
//	}
//	goTicket = make(chan byte, total)
//	for i := 0; i < total; i++ {
//		goTicket <- 1
//	}
//	return goTicket
//}
//
//func fecthPerson1() (Person, bool) {
//
//}
//
//func savePerson1(p Person) bool {
//
//}