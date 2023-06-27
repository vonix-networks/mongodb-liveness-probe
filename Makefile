server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o ./build/mongodb-liveness-probe ./cmd

image:
	docker build --platform=linux/amd64 -t us-docker.pkg.dev/vonix-io/public/mongodb-liveness-probe:latest .

deploy: image
	docker push us-docker.pkg.dev/vonix-io/public/mongodb-liveness-probe:latest
