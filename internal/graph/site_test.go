package graph

import (
	"github.com/mach-composer/mach-composer-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"testing"
)

func TestSite_Hash_NestedComponentConfigSorted(t *testing.T) {
	su := NewSite(nil, "", "", "", nil, config.SiteConfig{})
	su.NestedSiteComponentConfigs = []config.SiteComponentConfig{
		{Name: "b", Definition: &config.ComponentConfig{Name: "b"}},
		{Name: "a", Definition: &config.ComponentConfig{Name: "a"}},
	}

	unsortedHash, err := su.Hash()

	s := NewSite(nil, "", "", "", nil, config.SiteConfig{})
	s.NestedSiteComponentConfigs = []config.SiteComponentConfig{
		{Name: "a", Definition: &config.ComponentConfig{Name: "a"}},
		{Name: "b", Definition: &config.ComponentConfig{Name: "b"}},
	}
	sortedHash, err := s.Hash()

	assert.NoError(t, err)
	assert.Equal(t, unsortedHash, sortedHash, "Hashes should be equal")
}

func TestSite_HasChanges_True(t *testing.T) {
	s := NewSite(nil, "", "", "", nil, config.SiteConfig{})
	s.NestedSiteComponentConfigs = []config.SiteComponentConfig{
		{Name: "b", Definition: &config.ComponentConfig{Name: "b"}},
		{Name: "a", Definition: &config.ComponentConfig{Name: "a"}},
	}
	s.outputs = cty.ObjectVal(map[string]cty.Value{
		"hash": cty.ObjectVal(map[string]cty.Value{
			"sensitive": cty.BoolVal(false),
			"value":     cty.StringVal("different-hash"),
			"type":      cty.StringVal("some-type"),
		}),
	})

	changed, err := s.HasChanges()
	assert.NoError(t, err)
	assert.True(t, changed)
}

func TestSite_HasChanges_False(t *testing.T) {
	s := NewSite(nil, "", "", "", nil, config.SiteConfig{})
	s.NestedSiteComponentConfigs = []config.SiteComponentConfig{
		{Name: "b", Definition: &config.ComponentConfig{Name: "b"}},
		{Name: "a", Definition: &config.ComponentConfig{Name: "a"}},
	}
	s.outputs = cty.ObjectVal(map[string]cty.Value{
		"hash": cty.ObjectVal(map[string]cty.Value{
			"sensitive": cty.BoolVal(false),
			"value":     cty.StringVal("0bc0ceeb092a6d7f8a949dea9f9a264c9f86409a279b11421d688b2cd4a4367e"),
			"type":      cty.StringVal("some-type"),
		}),
	})

	changed, err := s.HasChanges()
	assert.NoError(t, err)
	assert.False(t, changed)
}