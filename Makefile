.PHONY: dev smoke unit

vagrant := /usr/local/bin/vagrant

# Ensure vagrant is installed, but do not install it because the user should have
# that right
${vagrant}:
	@which vagrant

dev: ${vagrant}
	vagrant up --provision
	@echo "You can now get to mesos at http://192.168.33.10:5050 and http://192.168.33.10:5050/api/v1"

smoke:
	@go test -v github.com/miroswan/mesops/test/smoke

unit:
	@go test -v -cover github.com/miroswan/mesops/pkg/v1
