URL ?= "http://pages.cs.wisc.edu/~ballard/bofh/excuses"


.PHONY: update
update:
	curl --location --output "templates/excuses.txt" $(URL)
