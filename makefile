
.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: assets-watch
assets-watch:
	$(MAKE) tailwind-watch & $(MAKE) templ-watch

.PHONY: air
air:
	air

.PHONY: run-dev
run-dev:
	$(MAKE) tailwind-watch & $(MAKE) templ-watch & $(MAKE) air