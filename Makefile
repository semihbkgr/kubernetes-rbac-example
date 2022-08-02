name:=kubernetes-rbac-example
version:=1.0.0
tag:=semihbkgr/$(name):$(version)

build:
	docker build --tag $(tag) .

push:
	docker push $(tag)
