.PHONY: swagger.run
swagger.run:
	@echo "===========> Generating swagger API docs"
	@swagger generate spec --scan-models  -o /root/iam/api/swagger/swagger.yaml

.PHONY: swagger.serve
swagger.serve:
	@swagger serve -F=redoc --no-open --port 36666 /root/iam/api/swagger/swagger.yaml