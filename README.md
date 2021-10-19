# go_learning
Practice and learning of goLang journey. Code snippets included. 


## Important Commands


go run <cwd> / <filePath>
  
  
`go build` compiles the packages, along with their dependencies, but it doesn't install the results.
  
  
`go install` compiles and installs the packages.
  
  
`go list -f '{{.Target}}'` to discover install path of go/bin
  
  
Add the Go install directory to your system's shell path.
    `export PATH=$PATH:/path/to/your/install/directory`
  
  
  As an alternative, if you already have a directory like $HOME/bin in your shell path and you'd like to install your Go programs there, you can change the install target by setting the GOBIN variable using the go env command:

 `go env -w GOBIN=/path/to/your/bin`
  
  
  `go mod init <packageName>`.  creates a go.mod file to track your code's dependencies.
  
  `go mod tidy`. Cleans and updates the changes made in package structure, dependencies.
  
  `go mod edit --replace <actualnameOfPackage>=<absolutePathOfPackage>`. This is helpful for local execution. Not recommended for prod. 
