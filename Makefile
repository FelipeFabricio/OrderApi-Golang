criar_migration_inicial:
	 migrate  create -ext=sql -dir=sql/migrations -seq migration_inicial 

migrate_up:
	migrate -path=sql/migrations -database "mysql://admin:root123!@tcp(localhost:3306)/wonderfood" -verbose up 

migrate_down:
	migrate -path=sql/migrations -database "mysql://admin:root123!@tcp(localhost:3306)/wonderfood" -verbose down

.PHONY: migrate_up migrate_down criar_migration_inicial