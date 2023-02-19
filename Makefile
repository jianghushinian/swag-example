.PHONY: swag-fmt
swag-fmt:
	swag fmt

.PHONY: swag-init
swag-init: swag-fmt
	swag init -g internal/api/controller/v1/docs.go --exclude internal/api/controller/v2 --instanceName v1
	swag init -g internal/api/controller/v2/docs.go --exclude internal/api/controller/v1 --instanceName v2
