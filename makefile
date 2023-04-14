gen_api:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package generated -generate types -old-config-style static/api/swagger.yaml > ./internal/server/generated/types.gen.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package generated -generate chi-server -old-config-style static/api/swagger.yaml > ./internal/server/generated/server.gen.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -package generated -generate spec -old-config-style static/api/swagger.yaml > ./internal/server/generated/spec.gen.go

gen_orm: ## Generate repository
	@go run entgo.io/ent/cmd/ent generate --target internal/storage/ent ./internal/storage/schema
