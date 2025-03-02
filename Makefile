target_directory := bin

.PHONY: all backend frontend clean

all: backend frontend

init:
	touch .env
	sudo docker pull postgres
	sudo docker create -it --name mp1_db --env-file .env --publish 127.0.0.1:5432:5432 postgres
	cd backend && go mod download
	cd frontend && npm install

frontend:
	cd frontend && npm run build

backend:
	mkdir -p $(target_directory)
	cd backend && go build -o ../$(target_directory)/mp1_backend

start:
	sudo docker start mp1_db
	./$(target_directory)/mp1_backend &
	cd frontend && npm run dev &

stop:
	cd frontend && npm run stop &
	killall mp1_backend
	sudo docker stop mp1_db

clean:
	rm -r $(target_directory)
