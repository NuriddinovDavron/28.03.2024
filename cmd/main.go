package main

import (
	"api_exam/api"
	"api_exam/config"
	"api_exam/pkg/logger"
	"api_exam/services"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// casbin with postgres -----------------------------------
	// psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, "localhost", 5432, "postgres", "1234", "backend")
	// db, err := gormadapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("gormadapter error", logger.Error(err))
	// 	return
	// }
	// casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, db)
	// if err != nil {
	// 	log.Error("NewEnforcer error", logger.Error(err))
	// 	return
	// }

	// casbin with CSV -------------------------------------------
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatal("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil { // gormadapter "github.com/casbin/gorm-adapter/v3"
		log.Fatal("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	apiServer := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		CasbinEnforcer: casbinEnforcer,
	})

	if err := apiServer.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
