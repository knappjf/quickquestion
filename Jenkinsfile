node {
  def root = tool name: '1.13.6', type: 'go'

  sh 'go mod tidy'
  sh 'go test github.com/knappjf/quickquestion' 
}