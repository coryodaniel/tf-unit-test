package main

import (
	"errors"
	"testing"
)

func getChangedResourceByName(plan Plan, address string) (ChangedResource, error) {
	var changedResource ChangedResource
	for _, resource := range plan.ResourceChanges {
		if resource.Address == address {
			return resource, nil
		}
	}

	return changedResource, errors.New("Not found")
}

func isDestroyed(resource ChangedResource) bool {
	for _, a := range resource.Change.Actions {
		if a == "delete" {
			return true
		}
	}
	return false
}

func TestPlan(t *testing.T) {
	t.Run("generates a tmp file path", func(t *testing.T) {
		plan := GetPlan()

		resource, _ := getChangedResourceByName(plan, "module.mymod.local_file.foo")

		got := resource.Change.After.Filename
		expected := "/tmp/index.md"

		if got != expected {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})

	t.Run("does not destroy foo", func(t *testing.T) {
		plan := GetPlan()

		resourceName := "module.mymod.local_file.foo"
		resource, _ := getChangedResourceByName(plan, resourceName)

		if isDestroyed(resource) {
			t.Errorf("expected %s to not be destroyed", resourceName)
		}
	})
}
