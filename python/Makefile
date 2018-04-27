docker-run:
	docker run -it -v $(HOME):/root --net=host onuryilmaz/k8s-client-example:python bash

docker-build:
	docker build -t onuryilmaz/k8s-client-example:python .

docker-push:
	docker push onuryilmaz/k8s-client-example:python

k8s-run:
	kubectl run python-client -it --image=onuryilmaz/k8s-client-example:python

k8s-logs:
	kubectl logs $(shell kubectl get pods --selector run=python-client -o jsonpath="{.items[0].metadata.name}")

k8s-stop:
	kubectl delete deployment python-client