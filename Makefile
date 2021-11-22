l/up:
	# Remove all, including the volumes.
	docker-compose down -v --remove-orphans --rmi all
	docker-compose --env-file ./.env up --detach --remove-orphans --force-recreate --build
