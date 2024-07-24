* Create newProject1
	1. go mod init newProject1
	2. touch main.go

* Create newProject2
	1. go mod init newProject2
	2. touch main.go

* Create go work space
	1. go work init newProject1 newProject2

* At the root project:
	Run: go run ./newProject2