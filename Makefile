
# Run templ generation in watch mode
templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

# Run air for Go hot reload 
server:
	@air 

# Watch Tailwind CSS changes
tailwind:
	@npx tailwindcss/cli -i ./assets/css/input.css -o ./assets/css/output.css --watch

# Start development server with all watchers
dev:
	@make -j3 tailwind templ server
