include .env
all: rs qs

rs:
	set -a && . ./.env && set +a && go run resource_server/cmd/launch_server/main.go &
qs:
	set -a && . ./.env && set +a && go run queuing_server/cmd/launch_server/main.go &
clean:
	killall main
