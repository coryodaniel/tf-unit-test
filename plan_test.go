package main

import "testing"

func TestPlan(t *testing.T) {
	t.Run("generates a tmp file path", func(t *testing.T) {
		plan := GetPlan()

		// Weaksauce, should be a method for getting the resource by 'address'
		// like the ruby example
		resource := plan.ResourceChanges[0]

		got := resource.Change.After.Filename
		expected := "/tmp/index.md"

		if got != expected {
			t.Errorf("got '%s' expected '%s'", got, expected)
		}
	})
}
