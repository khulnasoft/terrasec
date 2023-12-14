package policy

import (
	"reflect"
	"sort"
	"testing"
)

func TestPolicyTypeAllExpandedCorrectly(t *testing.T) {
	t.Run("policy type all gets right policy names", func(t *testing.T) {

		want := SupportedPolicyTypes(false)
		got := supportedCloudProvider["all"].policyNames()

		sort.Strings(want)
		sort.Strings(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})

	t.Run("policy type all gets right policy paths", func(t *testing.T) {

		want := GetDefaultPolicyPaths(SupportedPolicyTypes(false))
		got := GetDefaultPolicyPaths([]string{"all"})

		sort.Strings(want)
		sort.Strings(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})
}
