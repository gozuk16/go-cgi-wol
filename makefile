build:
	go build -ldflags "-s -w" -o wol.cgi main.go
install:
	chmod 755 wol.cgi
	cp -p wol.cgi /usr/local/var/h2o/cgi-bin/.
