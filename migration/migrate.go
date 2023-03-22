package migration

import (
	"fmt"
	"os"
	"os/exec"
	"vandyahmad24/maxsol/app/config"
)

func Up() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}
	//commnad := fmt.Sprintf(`migrate -database "mysql://%s:%s@tcp(%s:%d)/%s" -path migration up`,
	//	cfg.Db.User,
	//	cfg.Db.Password,
	//	cfg.Db.Address,
	//	cfg.Db.Port,
	//	cfg.Db.DbName)
	//fmt.Println("command ", commnad)
	command := exec.Command("migrate",
		"-database", fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
			cfg.Db.User,
			cfg.Db.Password,
			cfg.Db.Address,
			cfg.Db.Port,
			cfg.Db.DbName),
		"-path", "migration", "up")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		return err
	}
	fmt.Println("Migration berjalan")
	return nil
}
