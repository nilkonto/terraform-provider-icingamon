TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=terraform-provider-icingamon
BINARY=terraform-provider-icingamon
VERSION=2.0.1

default: build

# build: fmtcheck
# 	go install

docker_start:
	docker run -d --name icinga2 -p 8080:80 -p 8443:443 -p 5665:5665 -it jordan/icinga2:2.11.4
	sleep 20

docker_get_root_password:
	$(eval password:=$(shell docker exec icinga2 bash -c 'grep password /etc/icinga2/conf.d/api-users.conf' | awk -F'"' '{ print $$2}' ))
	echo $(password)

docker_clean:
	docker stop icinga2
	docker rm icinga2


build: fmtcheck
	GOOS=linux GOARCH=amd64 go build -o linux/terraform-provider-icingamon .
	GOOS=darwin GOARCH=amd64 go build -o darwin/terraform-provider-icingamon .

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}-darwin_amd64-${VERSION}
	cd ./bin && tar cvf ${BINARY}-darwin_amd64-${VERSION}.tar ${BINARY}-darwin_amd64-${VERSION} && gzip ${BINARY}-darwin_amd64-${VERSION}.tar
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}-linux_amd64-${VERSION}
	cd ./bin && tar cvf ${BINARY}-linux_amd64-${VERSION}.tar ${BINARY}-linux_amd64-${VERSION} && gzip ${BINARY}-linux_amd64-${VERSION}.tar
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}-linux_arm-${VERSION}
	cd ./bin && tar cvf ${BINARY}-linux_arm-${VERSION}.tar ${BINARY}-linux_arm-${VERSION} && gzip ${BINARY}-linux_arm-${VERSION}.tar
	

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

# testacc: fmtcheck
# 	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

testacc: fmtcheck
	$(eval password:=$(shell docker exec icinga2 bash -c 'grep password /etc/icinga2/conf.d/api-users.conf' | awk -F'"' '{ print $$2}' ))
	ICINGA2_API_PASSWORD="$(password)" ICINGA2_API_URL="https://127.0.0.1:5665/v1" ICINGA2_API_USER=root ICINGA2_INSECURE_SKIP_TLS_VERIFY=true TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile website website-test

