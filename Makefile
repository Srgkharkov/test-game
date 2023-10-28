build:
	go build -o ./bin/app ./cmd/main.go
	./bin/app

run:
	go run ./cmd/main.go

testresult:

	curl -X POST -H 'Content-Type:multipart/form-data' -F conftype=reels -F confname=confreels -F config=@test/confreels.json localhost:8080/addconfig
	curl -X POST -H 'Content-Type:multipart/form-data' -F conftype=lines -F confname=conflines -F config=@test/conflines.json localhost:8080/addconfig
	curl -X POST -H 'Content-Type:multipart/form-data' -F conftype=payouts -F confname=confpayouts -F config=@test/confpayouts.json localhost:8080/addconfig
	curl -X POST -H 'Content-Type:application/json' -d '{"conf_reels_name":"confreels","conf_lines_name":"conflines","conf_payouts_name":"confpayouts"}' localhost:8080/getresult
