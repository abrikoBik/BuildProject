target_directory := bin

.PHONY: all backend frontend clean

all: backend frontend

init:
	touch .env
	sudo docker pull postgres
	sudo docker create -it --name mp1_db --env-file .env --expose 5432 postgres

frontend:
	cp .env frontend/.env

backend:
	mkdir -p $(target_directory)
	cd backend && go build -o ../$(target_directory)/backend

start:
	sudo docker start mp1_db
	./$(target_directory)/backend

stop:
	sudo docker stop mp1_db

clean:
	rm -r $(target_directory)
