buildpi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/linksys_pi

clean:
	rm bin/*
