package test

import (
	"fmt"
	"github.com/9299381/bingo/package/db"
	"github.com/9299381/bingo/package/id"
	"testing"
	"xorm.io/builder"
)

type Detail struct {
	Id     string
	CardNo string
}
type UserDetail struct {
	CommUser `xorm:"extends"`
	Detail   `xorm:"extends"`
}

type CommUser struct {
	Id         string       `xorm:"pk varchar(21) notnull unique 'id'" json:"id"`
	UserName   string       `xorm:"varchar(50) not null 'user_name'" json:"user_name"`
	LoginName  string       `xorm:"varchar(20)  'login_name'" json:"login_name"`
	Status     string       `xorm:"varchar(2)  'status'" json:"status"`
	CreateTime db.LocalTime `xorm:"datetime created 'create_time'" json:"create_time"`
	UpdateTime db.LocalTime `xorm:"datetime updated 'update_time'" json:"update_time"`
}

func (s *CommUser) TableName() string {
	return "comm_user_info"

	//return "comm_user_info_copy1"
}

func TestListJoin(t *testing.T) {

	var users []UserDetail
	results := db.Engine().
		Table("comm_user_info").
		Alias("t1").
		Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
		Join("LEFT", "user_bank as t2", "t1.user_id=t2.user_id").
		Limit(10, 0).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestOneJoin(t *testing.T) {
	user := &UserDetail{}
	result, err := db.Engine().
		Table("comm_user_info").
		Alias("t1").
		Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
		Join("LEFT", "user_bank as t2", "t1.user_id=t2.user_id").
		Where("t1.id = ?", "1189164474851006208").
		And("t1.user_name = ?", "ccc").
		Limit(1).
		Get(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(user.CardNo)
}

func TestOne(t *testing.T) {
	user := &CommUser{}
	result, _ := db.Engine().
		//Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		Where("id =?", "1189164474851006208").
		Limit(1).
		Get(user)
	fmt.Println(result)
	fmt.Println(user)
}
func TestGet(t *testing.T) {
	user := &CommUser{Id: "1189164474851006208"}
	result, _ := db.Engine().Get(user)
	fmt.Println(result)
	fmt.Println(user)
}

func TestList(t *testing.T) {
	var users []CommUser
	results := db.Engine().
		Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		OrderBy("id DESC").
		Limit(5, 10).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestPage(t *testing.T) {
	var users []CommUser
	page := 1
	pageSize := 10 //页面大小
	results := db.Engine().
		Table("comm_user_info").
		Select("id,user_name,status,create_time,update_time").
		Where("status = ?", "20").
		OrderBy("id DESC").
		Limit(pageSize*(page), (page-1)*pageSize).
		Find(&users)
	fmt.Println(results)
	fmt.Println(users)
	for key, value := range users {
		fmt.Println(key)
		fmt.Println(value.UserName)
	}
}

func TestUpdateOne(t *testing.T) {
	user := &CommUser{Id: "1306582895206334464"}
	//_, _ = db.Engine().Get(user)
	//fmt.Println(user)
	user.UserName = "ccc"
	_, _ = db.Engine().Update(user, &CommUser{Id: user.Id})
	fmt.Println(user)
}

func TestUpdateTwo(t *testing.T) {
	user := &CommUser{Id: "1306582895206334464"}
	//_, _ = db.Engine().Get(user)
	//fmt.Println(user)
	user.UserName = "ccc"
	_, _ = db.Engine().
		ID(user.Id).
		Cols("user_name").
		Update(user)
	fmt.Println(user)
}

//insert
func TestInsert(t *testing.T) {
	user := &CommUser{
		Id:        id.New(),
		UserName:  "go_test",
		Status:    "30",
		LoginName: "aaaaa",
	}
	_, _ = db.Engine().Insert(user)

}

//可以创建表
func TestSync2(t *testing.T) {
	err := db.Engine().Sync2(new(CommUser))
	if err != nil {
		t.Error(err)
	}
}

/// builder sql 方式

func TestBuilderFetchListJoin(t *testing.T) {

	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Limit(10, 0).
			ToSQL()
	var users []UserDetail
	results := db.Engine().SQL(sql, args...).Find(&users)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(users)
}

func TestBuilderFetchOneJoin(t *testing.T) {
	sql, args, _ :=
		builder.
			Select("t1.id,t1.user_name,t1.create_time,t1.update_time,t2.card_no").
			From("comm_user_info as t1").
			LeftJoin("user_bank as t2", "t1.user_id=t2.user_id").
			Where(builder.Eq{"t1.id": "1189164474851006208"}).
			And(builder.Eq{"t1.user_name": "aaaaaaaaa"}).
			ToSQL()

	user := &UserDetail{}
	results, _ := db.Engine().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(results)
	fmt.Println(user.CardNo)
}
func TestBuilderFetchOne(t *testing.T) {
	req := make(map[string]interface{})
	req["id"] = "1306582895206334464"
	cond := builder.Eq{}
	for k, v := range req {
		cond[k] = v
	}
	sql, args, _ :=
		builder.
			Select("id,user_name,status,create_time,update_time").
			From("comm_user_info").
			Where(cond).
			ToSQL()

	user := &CommUser{}
	has, _ := db.Engine().SQL(sql, args...).Get(user)
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println(has)
	fmt.Println(user)
}

func TestBuilderFetch(t *testing.T) {
	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			OrderBy("id DESC").
			Limit(5, 10).
			ToSQL()

	var users []CommUser
	err := db.Engine().SQL(sql, args...).Find(&users)
	for _, v := range users {
		fmt.Println(v.Id)
	}
	fmt.Println(users)
	fmt.Println(err)
}
func TestBuilderFage(t *testing.T) {
	page := 1
	pageSize := 10 //页面大小

	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			Where(builder.Eq{"status": "20"}).
			OrderBy("id DESC").
			Limit(pageSize*(page), (page-1)*pageSize).
			ToSQL()

	var users []CommUser
	err := db.Engine().
		SQL(sql, args...).
		Find(&users)

	fmt.Println(users)
	fmt.Println(err)
}

func TestBuilderPageList(t *testing.T) {
	page := 1
	pageSize := 10 //页面大小

	sql, args, _ :=
		builder.
			Select("*").
			From("comm_user_info").
			Where(builder.Eq{"status": "20"}).
			OrderBy("id DESC").
			ToSQL()
	var users []CommUser
	err := db.Engine().
		SQL(sql, args...).
		Limit(pageSize*(page), (page-1)*pageSize).
		Find(&users)

	fmt.Println(users)
	fmt.Println(err)
}
