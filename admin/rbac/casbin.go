package rbac

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	Once        sync.Once
	NewEnforcer *casbin.Enforcer
)

type CasbinSubRule struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}

type CasbinRules struct {
	Role        string          `json:"role" form:"role"`
	CasbinInfos []CasbinSubRule `json:"casbininfos" form:"casbininfos"`
}

type CasbinRole struct {
	User string `json:"user" form:"user"`
	Role string `json:"role" form:"role"`
}

func NewCasbin(db *gorm.DB) *casbin.Enforcer {
	Once.Do(func() {
		modelText := `
			[request_definition]
			r = sub, obj, act

			[policy_definition]
			p = sub, obj, act

			[role_definition]
			g = _, _

			[policy_effect]
			e = some(where (p.eft == allow))

			[matchers]
			m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
		`
		model, _ := model.NewModelFromString(modelText)
		adapter, _ := gormadapter.NewAdapterByDB(db)
		NewEnforcer, _ = casbin.NewEnforcer(model, adapter)
		NewEnforcer.LoadPolicy()
	})
	return NewEnforcer
}
