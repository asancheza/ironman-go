.DEFAULT_GOAL := provision

.PHONY: image ensure-repos destroy
.PHONY: run stop attach provision ports

DOCKER = $(shell which docker)
DOCKER_COMPOSE = $(shell which docker-compose)

ifeq ($(DOCKER),)
	$(error "Docker not available on this system")
endif

ifeq ($(DOCKER_COMPOSE),)
	$(error "Docker Compose not available on this system")
endif

WEB_IMG_EXISTS = $(shell docker images -qa mob/web)
DB_IMG_EXISTS = $(shell docker images -qa mob/database)
VARNISH_IMG_EXISTS = $(shell docker images -qa mob/database)

image: 
	@$(DOCKER) build --rm=true -t ironman/backend-ui -f Dockerfile.backend-ui .
	@$(DOCKER) build --rm=true -t ironman/backend-storage -f Dockerfile.backend-storage .
	#@$(DOCKER) build --rm=true -t ironman/backend-proxy -f Dockerfile.backend-proxy .

swarm: 
	@$(DOCKER) swarm init --advertise-addr 192.168.1.99

run: image
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d 

start:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d

stop:
	$(DOCKER_COMPOSE) -f docker-compose.yml stop 

provision: run

attach-ui: 
	@$(DOCKER) exec -ti backend-ui /bin/bash -li || true

attach-storage:
	@$(DOCKER) exec -ti backend-storage /bin/bash -li || true

#attach-proxy:
#	@$(DOCKER) exec -ti backend-proxy /bin/bash -li || true

ports:
	@if [ -z "${DOCKER_HOST}" ]; then \
		export DOCKER_IP="127.0.0.1"; \
	else \
		export DOCKER_IP=`echo $${DOCKER_HOST} | cut -d: -f 2 | tr -d /`; \
	fi; \
	export WEB_PORT=`docker port web 80 | sed -e "s/.*:/$${DOCKER_IP}:/g" `; \
	echo "Web: http://$${WEB_PORT}"; \

destroy: 
	@$(DOCKER) rmi -f backend-ui
	@$(DOCKER) rmi -f backend-storage
#	@$(DOCKER) rmi -f ironman/backend-proxy

destroy-all:
	docker stop $(shell docker ps -a -q)
	docker rm $(shell docker ps -a -q)
