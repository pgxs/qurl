doc:
	vuepress dev ./docs
migrate:
	@echo "Run migrations"
	@cd scripts/sql/migrations && go-bindata -pkg migrations -o ../../../pkg/migrations/bindata.go . && cd ../../../ \
	&& QURL_CONF=$(PWD)/configs/server.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go test -v ./pkg/migrations/.
lint:
	@gofmt -s -w .
test:
	@echo "Run go test"
	@export QURL_CONF=$(PWD)/configs/server.yml \
	&& export PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	&& go test -count=1 pgxs.io/qurl/pkg/util \
	&& go test -count=1 pgxs.io/qurl/pkg/repository \
	&& go test -count=1 pgxs.io/qurl/pkg/service
test-all: migrate test
run:
	@QURL_CONF=$(PWD)/configs/server.yml \
	PG_TEST_DATA_FILE=$(PWD)/scripts/sql/test/data.sql \
	go run cmd/qurl/main.go


