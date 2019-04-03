package eidas

import (
	"encoding/hex"
	"testing"
)

var defaultCA = CompetentAuthority{
	Name: "Financial Conduct Authority",
	ID:   "GB-FCA",
}

func TestSimple(t *testing.T) {
	pspAS := "305b3013060604008e4601063009060704008e4601060330440606040081982702303a301330110607040081982701010c065053505f41530c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341"
	d, err := Serialize([]string{"PSP_AS"}, defaultCA, QWACType)
	if err != nil {
		t.Error(err)
	}
	if hex.EncodeToString(d) != pspAS {
		t.Error("Mismatch with PSP_AS")
	}
}

// TestAll tests against all the example data in "eIDAS PSD2 Certificate Signing Request Profiles", dated 12th March 2019.
func TestAll(t *testing.T) {
	type testData struct {
		Expected string
		Roles    []string
	}
	expected := []testData{testData{
		Expected: "305b3013060604008e4601063009060704008e4601060330440606040081982702303a301330110607040081982701010c065053505f41530c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS"},
	}, testData{
		Expected: "305b3013060604008e4601063009060704008e4601060330440606040081982702303a301330110607040081982701020c065053505f50490c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_PI"},
	}, testData{
		Expected: "305b3013060604008e4601063009060704008e4601060330440606040081982702303a301330110607040081982701030c065053505f41490c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AI"},
	}, testData{
		Expected: "305b3013060604008e4601063009060704008e4601060330440606040081982702303a301330110607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_IC"},
	}, testData{
		Expected: "306c3013060604008e4601063009060704008e4601060330550606040081982702304b302430220607040081982701010c065053505f41530607040081982701020c065053505f50490c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_PI"},
	}, testData{
		Expected: "306c3013060604008e4601063009060704008e4601060330550606040081982702304b302430220607040081982701010c065053505f41530607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_IC"},
	}, testData{
		Expected: "306c3013060604008e4601063009060704008e4601060330550606040081982702304b302430220607040081982701020c065053505f50490607040081982701030c065053505f41490c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_PI", "PSP_AI"},
	}, testData{
		Expected: "306c3013060604008e4601063009060704008e4601060330550606040081982702304b302430220607040081982701020c065053505f50490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_PI", "PSP_IC"},
	}, testData{
		Expected: "306c3013060604008e4601063009060704008e4601060330550606040081982702304b302430220607040081982701030c065053505f41490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AI", "PSP_IC"},
	}, testData{
		Expected: "307d3013060604008e4601063009060704008e4601060330660606040081982702305c303530330607040081982701010c065053505f41530607040081982701020c065053505f50490607040081982701030c065053505f41490c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_PI", "PSP_AI"},
	}, testData{
		Expected: "307d3013060604008e4601063009060704008e4601060330660606040081982702305c303530330607040081982701010c065053505f41530607040081982701020c065053505f50490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_PI", "PSP_IC"},
	}, testData{
		Expected: "307d3013060604008e4601063009060704008e4601060330660606040081982702305c303530330607040081982701010c065053505f41530607040081982701030c065053505f41490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_AI", "PSP_IC"},
	}, testData{
		Expected: "307d3013060604008e4601063009060704008e4601060330660606040081982702305c303530330607040081982701020c065053505f50490607040081982701030c065053505f41490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_PI", "PSP_AI", "PSP_IC"},
	}, testData{
		Expected: "30818e3013060604008e4601063009060704008e4601060330770606040081982702306d304630440607040081982701010c065053505f41530607040081982701020c065053505f50490607040081982701030c065053505f41490607040081982701040c065053505f49430c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341",
		Roles:    []string{"PSP_AS", "PSP_PI", "PSP_AI", "PSP_IC"},
	}}

	for _, e := range expected {
		_ = DumpFromHex(e.Expected)
		// Check our serialization matches theirs.
		s, err := Serialize(e.Roles, defaultCA, QWACType)
		if err != nil {
			t.Error(err)
		}
		if hex.EncodeToString(s) != e.Expected {
			t.Errorf("Mismatch with roles: %v", e.Roles)
		}

		// Check we can extract the roles, name and ID correctly.
		d, err := hex.DecodeString(e.Expected)
		if err != nil {
			t.Error(err)
		}
		roles, name, id, err := Extract(d)
		if err != nil {
			t.Error(err)
		}
		for i, r := range roles {
			if e.Roles[i] != r {
				t.Errorf("Expected role: %s but got %s", e.Roles[i], r)
			}
		}
		if name != defaultCA.Name {
			t.Errorf("Expected CA name: %s but got %s", defaultCA.Name, name)
		}
		if id != defaultCA.ID {
			t.Errorf("Expected CA id: %s but got %s", defaultCA.ID, id)
		}
	}
}

func TestQSEAL(t *testing.T) {
	pspAS := "305b3013060604008e4601063009060704008e4601060230440606040081982702303a301330110607040081982701010c065053505f41530c1b46696e616e6369616c20436f6e6475637420417574686f726974790c0647422d464341"
	d, err := Serialize([]string{"PSP_AS"}, defaultCA, QSEALType)
	if err != nil {
		t.Error(err)
	}
	if hex.EncodeToString(d) != pspAS {
		t.Error("Mismatch with PSP_AS")
	}
}
