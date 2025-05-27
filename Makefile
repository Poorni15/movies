run:
	@go run cmd/main.go

migrate-update:
	liquibase --changeLogFile=./migration/changelog.xml --url=jdbc:postgresql://localhost:5432/postgres --username=postgres --password=poorni1512 update