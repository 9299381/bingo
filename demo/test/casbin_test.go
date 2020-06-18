package test

import (
	"fmt"
	"github.com/casbin/casbin"
	string_adapter "github.com/qiangmzsx/string-adapter"
	"testing"
)

func TestCasbin(t *testing.T) {
	model := casbin.NewModel()
	model.LoadModelFromText(GetModel1())
	policy := string_adapter.NewAdapter(GetPolicy1())
	e := casbin.NewEnforcer(model, policy)
	fmt.Println(e.GetPolicy())
	fmt.Println(e.GetPermissionsForUser("P1301"))
	//(sub, obj, act)
	if e.Enforce("P1101", "broadcast", "push") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	//(sub, obj, act)
	if e.Enforce("P1101", "broadcast", "pull") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//------------------//
	// GetAllSubjects 获取当前策略中显示的主题列表。
	fmt.Println(e.GetAllObjects())
	fmt.Println(e.GetAllNamedObjects("p"))
	// action 获取当前策略中显示的操作列表
	fmt.Println(e.GetAllActions())
	fmt.Println(e.GetAllNamedActions("p"))

	// subject 获取当前命名策略中显示的主题列表
	fmt.Println(e.GetAllSubjects())
	fmt.Println(e.GetAllNamedSubjects("p"))

	fmt.Println(e.GetPolicy())
	fmt.Println(e.GetNamedPolicy("p"))

	fmt.Println(e.HasPolicy("P1101", "broadcast", "push"))
	fmt.Println(e.HasPolicy("P1101", "broadcast", "pull"))

}

func GetModel1() string {
	model := `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act

`

	return model
}
func GetPolicy1() string {
	policy := `
	p, P1101, broadcast, pull

	p, P1301, broadcast, push
	p, P1301, broadcast, pull
	p, P1301, supply,allow

	p, P1501, broadcast, push
	p, P1501, broadcast, pull

	p, P1701, broadcast, push

`

	return policy
}

/**
#原理分析
PERM(Policy, Effect, Request, Matchers)模型很简单, 但是反映了权限的本质 – 访问控制

* Policy: 定义权限的规则
* Effect: 定义组合了多个 Policy 之后的结果, allow/deny
* Request: 访问请求, 也就是谁想操作什么
* Matcher: 判断 Request 是否满足 Policy

![](https://images.cnblogs.com/cnblogs_com/wang_yb/1345158/o_perm.png)


典型的配置格式（RABC）：
```
# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act
p2 = sub, act
//定义的每一行称为 policy rule, p, p2 是 policy rule 的名字.
//p2 定义的是 sub 所有的资源都能执行 act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))
// 上面表示有任意一条 policy rule 满足, 则最终结果为 allow

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
//定义了 request 和 policy 匹配的方式, p.eft 是 allow 还是 deny, 就是基于此来决定的

[role_definition]
g = _, _
g2 = _, _
g3 = _, _, _
//g, g2, g3 表示不同的 RBAC 体系, _, _ 表示用户和角色 _, _, _ 表示用户, 角色, 域(也就是租户)

```
sub, obj, act，分别表示，用户（或者分组），资源对象，权限
* request_definition：表示请求的格式
* policy_definition：表示权限的格式
* policy_effect：表示匹配之后的动作
* matchers：表示匹配的规则，通过运算得出bool值

有了权限定义的文件之后，还需要有一个用户或者角色分组的存储的信息：
```
p, alice, data1, read
p, bob, data2, write
```
这里的格式，表示：
* alice对资源data1可以进行read
* bob对资源data2可以进行write

当alice去读取data1的时候，就会触发matchers的规则，通过存储的角色信息进行运算，按照上面的例子，就可以通过审核。

通过对权限配置和用户角色存储的结构进行组合，然后按照matchers的规则，就可以得出用户是否有权限对资源进行访问了。


`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act
	p2 = sub, act

	[role_definition]
	g = _, _
	g2 = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`


		`
	p, superAdmin, project, read
	p, superAdmin, project, write
	p, admin, project, read
	p, admin, project, write
	p, admin, asse, read
	p, admin, asse, write
	p, zhuangjia, project, write
	p, zhuangjia, asse, write
	p, shangshang, project, read
	p, shangshang, asse, read
	p2, user, read
	g, quyuan, admin
	g, wenyin, zhuangjia
	g2,wangjh,user

`
*/
