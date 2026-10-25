package agent

import (
	"slices"
	"testing"

	"github.com/yetone/magpie/internal/usage"
)

// LobeHub names itself in its User-Agent, `lobehub/<version>`, from both of
// the layers it talks to a vendor with. It is a client magpie knows rather
// than an agent it sets up, so the Routing and Usage views draw it with its
// logo, and the Agents page shows no row for it.
func TestLobeHubIsNamed(t *testing.T) {
	for ua, want := range map[string]string{
		"lobehub/2.2.20": "lobehub",
		"LobeHub/2.2.20": "lobehub",
		"lobehub":        "lobehub",
	} {
		if got := usage.AgentOf(ua); got != want {
			t.Errorf("%q: %q, want %q", ua, got, want)
		}
	}
	if !slices.ContainsFunc(Clients(), func(a *Agent) bool {
		return a.ID == "lobehub" && a.Name == "LobeHub" && a.Icon == "lobehub-color"
	}) {
		t.Error("LobeHub isn't a client the Usage and Routing views draw")
	}
}
