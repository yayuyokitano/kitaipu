package kitaipu

import (
	"encoding/json"
	"testing"
)

func TestJSONUnMarshal(t *testing.T) {
	var i Command
	err := json.Unmarshal([]byte(`{"token":"aW50ZXJhY3Rpb246OTQ5MDEzNTYwNzExODU2MjU5OllaY25UZHdXMGYwcDBtRW9NaTdlOFBFMFQ0MlZhaVZONTBlM3BsOVpjNlllVXJIVm9vOHFRbUgwWkJhdEllVHQyQ2VHNjRKMTUwNkdkZUx3VkhCdmdvSUFyeE0zanluNE0yODdjeXpDbUhRMmRIeUdwMVl3Q1gxdVNROUR3ZnFV","data":{"id":"948994743487062127","type":1,"name":"level","options":[{"type":1,"options":[],"name":"display"}]},"application_id":"541298511430287395","member":{"premium_since":null,"joined_at":"2020-06-07T18:23:11.395000+00:00","avatar":null,"user":{"public_flags":0,"id":"196249128286552064","avatar":"a_91417d8d7fa6a87bdcbb85c4551b40c4","username":"Themex","discriminator":"2404"},"permissions":"2199023255551","mute":false,"communication_disabled_until":null,"nick":null,"roles":["947536902112813077"],"deaf":false,"is_pending":false,"pending":false},"locale":"no","channel_id":"942872224254292018","guild_locale":"en-US","version":1,"id":"949013560711856259","guild_id":"719255152170762301","type":2}`), &i)
	if err != nil {
		t.Error(err)
	}
	if i.Token != "aW50ZXJhY3Rpb246OTQ5MDEzNTYwNzExODU2MjU5OllaY25UZHdXMGYwcDBtRW9NaTdlOFBFMFQ0MlZhaVZONTBlM3BsOVpjNlllVXJIVm9vOHFRbUgwWkJhdEllVHQyQ2VHNjRKMTUwNkdkZUx3VkhCdmdvSUFyeE0zanluNE0yODdjeXpDbUhRMmRIeUdwMVl3Q1gxdVNROUR3ZnFV" {
		t.Error("Token is not correct")
	}

}

func TestOptionGet(t *testing.T) {
	opts := Options{
		{
			Name: "test1",
			Type: OptSubCommand,
			Options: Options{
				{
					Name:  "test2",
					Type:  OptUser,
					Value: "test2val",
				},
				{
					Name:  "test3",
					Type:  OptBool,
					Value: "test3val",
				},
			},
		},
	}
	opt, ok := opts.Get("test1")
	if !ok {
		t.Error("Option not found")
	}
	if opt.Name != "test1" {
		t.Error("Name is not correct")
	}
	if opt.Type != OptSubCommand {
		t.Error("Type is not correct")
	}

	opt, ok = opt.Options.Get("test3")
	if !ok {
		t.Error("Option not found")
	}
	if opt.Name != "test3" {
		t.Error("Name is not correct")
	}
	if opt.Type != OptBool {
		t.Error("Type is not correct")
	}
	if opt.Value != "test3val" {
		t.Error("Value is not correct")
	}

	opt = Option{}

	opt, ok = opts.Get("test2")
	if ok {
		t.Error("Option found")
	}
	if opt.Name != "" {
		t.Errorf("opt should not have a name, name is %s", opt.Name)
	}

}
