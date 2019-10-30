.PHONY: test.ruby test.go test.init test

test: test.init test.ruby test.go clean

test.init:
	terraform init
	TF_VAR_name=index.md terraform plan | \
		parse-terraform-plan -o plan.json

test.go:
	go test -v

test.ruby:
	rspec --format doc

clean:
	rm plan.json
