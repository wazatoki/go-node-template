ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: execute
execute: serve
	docker exec tmp_project_name /workspace/tmp_project_name_linux_amd64.bin

.PHONY: stop
stop:
	cd ../tools && \
	docker-compose down -v

.PHONY: serve
serve:
	cd ../tools && \
	docker-compose up -d

.PHONY: clean
clean:
	rm -rf build

.PHONY: test
test:
	cd src && \
	go test ./... && \
	cd $(ROOT_DIR)

.PHONY: buildAll
buildAll: clean build
	cd ../frontend && \
	npm run build && \
	cp -r ./dist ../backend/build/resources/frontend

.PHONY: build
build: clean
	mkdir build && \
	cd src && \
	export GOOS=linux && \
	go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o ../build/tmp_project_name_linux_amd64.bin . && \
	export GOOS=windows && \
	go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o ../build/tmp_project_name_windows_amd64.exe . && \
	cd $(ROOT_DIR) && \
	cp -r resources build/

.PHONY: migrateUP
migrateUP:
	migrate -path resources/migrations -database postgres://tmp_project_name:tmp_project_name@tmp_project_name_postgres_db:5432/tmp_project_namedb?sslmode=disable up

.PHONY: migrateDOWN
migrateDOWN:
	migrate -path resources/migrations -database postgres://tmp_project_name:tmp_project_name@tmp_project_name_postgres_db:5432/tmp_project_namedb?sslmode=disable down

.PHONY: migrateCreate
migrateCreate:
	ARG=""
	migrate create -ext sql -dir migrations -seq ${ARG}
