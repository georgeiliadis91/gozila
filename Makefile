run:
	@templ generate
	
run-both:
	@go run main.go & rollup -c -w

