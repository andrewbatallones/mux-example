# Mux Tutorial
You can follow here: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

## Things I've learned
- You shouldn't import the main package. (If you're using the gofmt, it will remove it)

## Things that I've changed
- I moved a lot of the files to their own package to keep things organized as well as have them easier to test since I cannot import the main package.