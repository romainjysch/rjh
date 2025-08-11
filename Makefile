# Binary
b:
	go build -o rjh

mv:
	sudo mv rjh /usr/local/bin/rjh

# Vendoring
tv:
	go mod tidy && go mod vendor

# Quick tests
c:
	./rjh w c Morges

f:
	./rjh w f Morges

n:
	./rjh n p -c 2 1.1.1.1

t:
	./rjh t ls
