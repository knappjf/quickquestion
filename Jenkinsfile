pipeline {
  agent { docker { image 'golang' }}
  stages {
    stage('dependencies') {
      sh go mod tidy
    }
    stage('test') {
      sh go test ./...
    }
    stage('build') {
      steps{
        sh go build -o 'bin/service' 'github.com/knappjf/quickquestion' 
      }
    }
  }
}