
app = "play\#/ipfs/QmTKZgRNwDNZwHtJSjCp6r5FYefzpULfy37JvMt9DwvXse"
url = "http://localhost:8080/ipfs/"

publish:
	@echo $(url)$(shell ipfs add -r -q . | tail -n1)/$(app)
