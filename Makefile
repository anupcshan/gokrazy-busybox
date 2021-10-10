bins/busybox.x64: Dockerfile busybox.config
	docker build --rm -t serial .
	docker run --rm -v $$(pwd)/bins:/tmp/bins serial cp -r /bins/ /tmp/
