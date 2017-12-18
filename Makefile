.PHONY: dev


dev:
	docker build . --file docker/dev -t mesops-dev
