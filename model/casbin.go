package model

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Casbin struct {
	Self *casbin.Enforcer
}

var CB *Casbin

func initCasbin(username, password, addr, DBName string) *casbin.Enforcer {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		addr,
		DBName)

	a, err := gormadapter.NewAdapter("mysql", config, true) // Your driver and data source.
	if err != nil {
		zap.L().Error("casbin数据库加载失败!", zap.Error(err))
		return nil
	}

	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub,p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
		return nil
	}
	cb, _ := casbin.NewEnforcer(m, a)

	return cb
}

func (c *Casbin) Init() {
	CB = &Casbin{
		Self: initCasbin(viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.addr"),
			viper.GetString("db.name")),
	}
}
