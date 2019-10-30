# tf unit test

Just an idea I was toying w/ once, dont use this.

## Setup

`yarn add global terraform-plan-parser`

## Run a unit test

```shell
TF_VAR_name=index.md terraform plan > plan.stdout
parse-terraform-plan -i ./plan.stdout -o plan.json

rspec --format doc
rm plan.json
```

