// Run on an agent where we want to use Go
node {
    // Install the desired Go version
    def root = tool name: 'go 1.10', type: 'go'
    checkout scm
    // Export environment variables pointing to the directory where Go was installed
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        sh 'go version'
        sh 'go build .'
    }
}