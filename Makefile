default: build
	gulp build

.PHONY: elm
elm:
	cd elm && elm make src/ContactForm.elm --output /home/luis/projects/boxd/src/assets/boxd/ContactForm.js #--optimize
	cd elm && elm make src/JobForm.elm --output /home/luis/projects/boxd/src/assets/boxd/JobForm.js #--optimize

build: elm
	cd pkg/contact-us && go build -o /home/luis/projects/boxd/netlify/functions
