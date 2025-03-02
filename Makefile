target_directory := bin

.PHONY: all backend frontend clean

all: backend frontend

init:
	touch .env
	sudo docker pull postgres
	sudo docker create -it --name mp1_db --env-file .env --publish 127.0.0.1:5432:5432 postgres

frontend:
	cp .env frontend/.env
	cd frontend && npm run build

backend:
	mkdir -p $(target_directory)
	cd backend && go build -o ../$(target_directory)/mp1_backend

start:
	sudo docker start mp1_db
	./$(target_directory)/mp1_backend &
	cd frontend && npm run dev &

stop:
	sudo docker stop mp1_db
	cd frontend && npm run stop

clean:
	rm -r $(target_directory)
