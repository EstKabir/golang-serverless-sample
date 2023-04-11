APP_BINARY_FILE=build/app.out
APP_MAIN_FILE=src/main.go

run_build:
	./${APP_BINARY_FILE}
build_app:
	go build -o ${APP_BINARY_FILE} ${APP_MAIN_FILE}
run:
	go run ${APP_MAIN_FILE}
clean:
	go clean
	rm ${APP_BINARY_FILE}