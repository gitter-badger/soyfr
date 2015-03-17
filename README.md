[![Build Status](https://travis-ci.org/manyminds/soyfr.svg?branch=master)](https://travis-ci.org/manyminds/soyfr)

# soyfr
a crowd based party drinking game

#installation instructions Mac OS X

```
brew install mongo
bower install
npm install
```

#running the application
in order to run the application you need to compile frontend 
files with grunt and after that start the go server. 

```
grunt
godep go run main.go
```

#get command line options
```
godep go run main.go -help
```

the application can now be reached via [0.0.0.0:8800](http://0.0.0.0:8800).
