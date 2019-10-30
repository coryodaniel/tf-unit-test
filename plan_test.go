package main

import "testing"

func TestPlan(t *testing.T) {
  t.Run("generates a tmp file path", func(t *testing.T) {
    plan := GetPlan()

    resource := plan.ChangedResources[0]

    got := resource.ChangedAttributes.Filename.New.Value
    expected := "/tmp/index.md"

    if got != expected {
      t.Errorf("got '%s' expected '%s'", got, expected)
    }
  })
}
