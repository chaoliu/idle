REGISTRY  = 007726414523.dkr.ecr.cn-northwest-1.amazonaws.com.cn
NAME = space/idle
VERSION = 1.0.0

.PHONY: build start push tag

build:
	docker build -t ${NAME}:${VERSION}  .
tag:
	docker tag ${NAME}:${VERSION} ${REGISTRY}/${NAME}:latest;
	docker tag ${NAME}:${VERSION} ${REGISTRY}/${NAME}:${VERSION};

push: build tag
	docker push ${REGISTRY}/${NAME}:latest; docker push ${REGISTRY}${NAME}:${VERSION}

start:
	docker run -it --rm -p 3000:3000 ${NAME}:${VERSION}
