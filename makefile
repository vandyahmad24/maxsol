
create-migration:
	migrate create -ext sql -dir migration $(NAME)	
run-migration:
	@echo "running migration ${PORT_GOLANG}"
	migrate -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}" -path migration up
run:
	go run main.go