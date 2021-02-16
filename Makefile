dev:
	make -j web auth circle

web:
	go run handlers/web/main.go

auth:
	go run handlers/auth/main.go

circle:
	go run handlers/circle/main.go