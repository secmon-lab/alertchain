package action

run[res] {
	res := {
		"id": "run_mock",
		"uses": "mock",
	}
}

exit[res] {
	input.action.id == "run_mock"
	input.alert.attrs[k2].name == "k2"
	res := {"attrs": {
		{
			"id": input.alert.attrs[k2].id,
			"name": "k2",
			"value": "v2a",
		},
		{
			"name": "k3",
			"value": "v3",
		},
	}}
}

run[res] {
	input.called[_].id == "run_mock"
	res := {
		"id": "run2",
		"uses": "mock.after",
	}
}
