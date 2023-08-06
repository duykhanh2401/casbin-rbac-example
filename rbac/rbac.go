package rbac

import (
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var RBAC *casbin.Enforcer

func Setup() *casbin.Enforcer {
	dataSourceName := "rbac:rbac@(localhost:3306)/casbin"
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	// sub: Đối tượng muốn truy cập vào 1 tài nguyên
	// obj: Tài nguyên sẽ đc truy cập
	// act: Phương thức mà ng dùng thực hiện trên tài nguyên
	// request: Phần gửi lên để check
	// policy: Phần tài nguyên đc định nghĩa để kiểm tra
	// role: User - Role
	text := `
        [request_definition]
        r = sub, obj, act
        r2 = sub, obj, act
        r3 = sub, obj, act

        [policy_definition]
        p = sub, obj, act
        p2 = sub, obj, act
        p3 = sub, obj, act

        [role_definition]
        g = _, _
        g2 = _, _, _
        g3 = _, _, _

        [policy_effect]
        e = some(where (p.eft == allow))
        e2 = some(where (p.eft == allow)) && !some(where (p.eft == deny))
        e3 = some(where (p.eft == allow)) && !some(where (p.eft == deny))

        [matchers]
        m = keyMatch2(r.obj, p.obj) && r.act == p.act && g(r.sub, p.sub)
        m2 = keyMatch2(r2.obj, p2.obj) && r2.act == p2.act && g2(r2.sub, keyGet2(r2.obj, p2.obj, 'id'), p2.sub)    
        m3 = keyMatch2(r3.obj, p3.obj) && r3.act == p3.act && g3(r3.sub, keyGet2(r3.obj, p3.obj, 'id'), p3.sub)
    `
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}

	a, err := sqlxadapter.NewAdapter(db, "casbin_rule")

	r, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}

	r.LoadPolicy()
	RBAC = r

	return r

}
