.PHONY: clean test.ruby test.go test.init test

test: test.ruby test.go clean

test.init: clean
	terraform init
	TF_VAR_name=index.md terraform plan -out tf.plan
	terraform show -json tf.plan > tf.json

test.go: test.init
	go test -v

test.ruby: test.init
	rspec --format doc

clean:
	rm -f tf.plan
	rm -f tf.json
