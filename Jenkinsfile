pipeline {
  agent any

  environment {
    GOPATH = "${env.WORKSPACE}/.gopath"
    GOBIN  = "${env.GOPATH}/bin"
    PATH   = "${env.PATH}:${env.GOBIN}"
    DIST   = "dist"
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
        sh 'go version || true'
      }
    }

    stage('Cache & Env') {
      steps {
        sh 'mkdir -p ${GOPATH} ${GOBIN} ${DIST}'
      }
    }

    stage('Build') {
      steps {
        sh 'go mod tidy'
        sh 'go build -o ${DIST}/app ./cmd/app'
      }
    }

    stage('Test') {
      steps {
        sh 'go install gotest.tools/gotestsum@latest'
        sh '${GOBIN}/gotestsum --junitfile test-report.xml -- -v ./...'
        junit 'test-report.xml'
      }
    }

    stage('Security Scan') {
      steps {
        sh 'go install github.com/securego/gosec/v2/cmd/gosec@latest'
        sh '${GOBIN}/gosec -fmt junit-xml -out gosec-report.xml ./... || true'
        junit 'gosec-report.xml'
        archiveArtifacts artifacts: 'gosec-report.xml', fingerprint: true
      }
    }

    stage('Publish Artifacts') {
      steps {
        archiveArtifacts artifacts: '${DIST}/**', fingerprint: true
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}
