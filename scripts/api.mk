PROJECT?=github.com/donskova1ex/application_aggregator
API_NAME?=application_aggregator
API_VERSION?=0.0.1
API_CONTAINER_NAME?=docker.io/donskova1ex/${API_NAME}


clean_api:
	rm -rf bin/application_aggregator

api_docker_build:
	docker build --no-cache -t ${API_CONTAINER_NAME}:${API_VERSION} -t ${API_CONTAINER_NAME}:latest -f Dockerfile.api .