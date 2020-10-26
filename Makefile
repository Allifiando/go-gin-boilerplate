.PHONY: run-dev
run-dev:
	bash -c "export ENV=development && nodemon --exec go run main.go --signal SIGTERM"