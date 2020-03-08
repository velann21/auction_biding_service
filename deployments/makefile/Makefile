setup:
	if [ -e ./dist ]; then rm -rf ./dist; fi; mkdir ./dist; mkdir -p ./dist/darwin/bin; mkdir -p ./dist/linux/bin

build:
	@echo "******************"
	@echo "**** Building ****"
	@echo "******************"

	GO111MODULE=on GOOS=linux go build -mod=vendor ../../cmd/auction_biding_bootstrap.go;cp dist/linux/bin/auction_biding_bootstrap ../docker;mv ./auction_biding_bootstrap dist/linux/bin/auction_biding_bootstrap


docker-build-prep:
	cd docker;cp ../dist/linux/bin/UserService .

docker-build-cleanup:
	cd ../docker;rm auction_biding_bootstrap

release-tag: build

	cd ../; git tag `cat .version`
	git push --tags

release: release-tag
	cd ../docker; docker build -t "singaravelan21/auction_biding" . --no-cache;
	cd ../docker; docker tag "singaravelan21/auction_biding" "singaravelan21/auction_biding:`cat ../.version`";
	cd ../docker; docker push "singaravelan21/auction_biding:`cat ../.version`"
	make docker-build-cleanup

