package alert.scc

test_scc {
	got := alert with input as data.input
	got[_].title == "Exfiltration: CloudSQL Over-Privileged Grant"
	print(got)
}
