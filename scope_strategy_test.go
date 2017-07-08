package fosite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHierarchicScopeStrategy(t *testing.T) {
	var strategy ScopeStrategy = HierarchicScopeStrategy
	var scopes = []string{}

	assert.False(t, strategy(scopes, "foo.bar.baz"))
	assert.False(t, strategy(scopes, "foo.bar"))
	assert.False(t, strategy(scopes, "foo"))

	scopes = []string{"foo.bar", "bar.baz", "baz.baz.1", "baz.baz.2", "baz.baz.3", "baz.baz.baz"}
	assert.True(t, strategy(scopes, "foo.bar.baz"))
	assert.True(t, strategy(scopes, "baz.baz.baz"))
	assert.True(t, strategy(scopes, "foo.bar"))
	assert.False(t, strategy(scopes, "foo"))

	assert.True(t, strategy(scopes, "bar.baz"))
	assert.True(t, strategy(scopes, "bar.baz.zad"))
	assert.False(t, strategy(scopes, "bar"))
	assert.False(t, strategy(scopes, "baz"))

	scopes = []string{"fosite.keys.create", "fosite.keys.get", "fosite.keys.delete", "fosite.keys.update"}
	assert.True(t, strategy(scopes, "fosite.keys.delete"))
	assert.True(t, strategy(scopes, "fosite.keys.get"))
	assert.True(t, strategy(scopes, "fosite.keys.get"))
	assert.True(t, strategy(scopes, "fosite.keys.update"))

	scopes = []string{"hydra", "openid", "offline"}
	assert.False(t, strategy(scopes, "foo.bar"))
	assert.False(t, strategy(scopes, "foo"))
	assert.True(t, strategy(scopes, "hydra"))
	assert.True(t, strategy(scopes, "hydra.bar"))
	assert.True(t, strategy(scopes, "openid"))
	assert.True(t, strategy(scopes, "openid.baz.bar"))
	assert.True(t, strategy(scopes, "offline"))
	assert.True(t, strategy(scopes, "offline.baz.bar.baz"))
}

func TestWildcardScopeStrategy(t *testing.T) {
	var strategy ScopeStrategy = WildcardScopeStrategy
	var scopes = []string{}

	assert.False(t, strategy(scopes, "foo.bar.baz"))
	assert.False(t, strategy(scopes, "foo.*.bar"))
	assert.False(t, strategy(scopes, "foo.*"))
	assert.False(t, strategy(scopes, "*"))

	scopes = []string{""}
	assert.False(t, strategy(scopes, "*"))

	scopes = []string{"foo"}
	assert.True(t, strategy(scopes, "*"))
	assert.False(t, strategy(scopes, "foo.*"))
	assert.False(t, strategy(scopes, "fo*"))

	scopes = []string{"foo.bar"}
	assert.True(t, strategy(scopes, "foo.*"))

	scopes = []string{"foo.baz"}
	assert.True(t, strategy(scopes, "foo.*"))
	assert.False(t, strategy(scopes, "foo.*.foo"))
	assert.False(t, strategy(scopes, "foo.*."))
	assert.False(t, strategy(scopes, "foo.foo.*."))
	assert.False(t, strategy(scopes, "foo.foo.*"))

	scopes = []string{"foo.baz.bar"}
	assert.False(t, strategy(scopes, "foo.*"))
	assert.True(t, strategy(scopes, "foo.*.*"))
	assert.True(t, strategy(scopes, "foo.*.bar"))
	assert.True(t, strategy(scopes, "foo.baz.*"))
	assert.True(t, strategy(scopes, "foo.baz.bar"))
	assert.False(t, strategy(scopes, "foo.b*.bar"))

	scopes = []string{"foo.bar", "foo.baz.bar"}
	assert.True(t, strategy(scopes, "foo.*"))
	assert.True(t, strategy(scopes, "foo.*.*"))
	assert.True(t, strategy(scopes, "foo.*.bar"))
	assert.False(t, strategy(scopes, "foo.bar.*"))
	assert.True(t, strategy(scopes, "foo.baz.*"))

	scopes = []string{"foo..bar"}
	assert.False(t, strategy(scopes, "foo.*.bar"))
}
