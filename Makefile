name:=k8s-rbas-example
version:=1.0.2
tag:=semihbkgr/$(name):$(version)

build:
	docker build --tag $(tag) .

push:
	docker push $(tag)
