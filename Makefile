deploy:
	git pull origin development
	docker image build -t stevenfrst/ppob-service .
	docker image push stevenfrst/ppob-service
	docker stack rm api
	docker stack deploy -c docker-compose.yml api