
app = "qr"
url = "http://localhost:8080/ipfs/"

publish:
	@echo $(url)$(shell ipfs add -r -q . | tail -n1)/$(app)
