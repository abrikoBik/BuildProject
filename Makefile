init:
	sudo docker pull postgres
	sudo docker create -it --name mp1_db --env-file .env --expose 5432 postgres
	touch .env

build:
	mkdir -p bin
	go build -o bin
	ln -s .env bin

start:
	sudo docker start mp1_db
	./bin/mp1

stop:
	sudo docker stop mp1_db

clean:
	rm -r bin
