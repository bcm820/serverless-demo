isprime:
	kubeless function deploy isprime --runtime go1.10 --handler func.isprime --from-file basics/hello.go --dependencies Gopkg.toml
hello-decipher:
	kubeless function deploy hello-decipher --runtime go1.10 --handler func.Handler --from-file basics/hello.go --dependencies Gopkg.toml
