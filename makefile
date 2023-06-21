start:
	docker compose up -d
	CompileDaemon -command="./ds-delivery-go"
