pipeline {
    agent any 

    environment {
        DOCKER_IMAGE = "hddocker125/demo-jenk-service"
        DOCKER_TAG = "latest"
        DOCKER_REGISTRY = "docker.io"
        DOCKER_DIGEST = "sha256:a449d2e0b0b2ed78f175d96f41650485ce597400a0db6fb8fd1aa18d5ee282b1"
    }
    tools {
        go 'Go1.25' // name configured in Jenkins
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building Go Application...'
                sh 'go mod tidy'
                sh 'go build -o app .'
            }
        }

        stage('Docker Build & Push') {
            steps {
                echo 'Building and pushing Docker image...'
                sleep 15
            }
        }

        stage('Registering build artifact') {
            agent any  
            steps {
                echo "Registering Docker artifact..."
                script {
                    registerBuildArtifactMetadata(
                        name: "demo-jenk-service",
                        version: "1.0.0",
                        type: "docker",
                        url: "${env.DOCKER_REGISTRY}/${env.DOCKER_IMAGE}:${env.DOCKER_TAG}",
                        digest: "${env.DOCKER_DIGEST}",
                        label: "qa, prod"
                    )
                }
            }
        }

        stage('Test') {
            steps {
                echo 'Running Unit Tests...'
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            agent any
            steps {
                echo 'Deploying...'
                sleep 10
            }
        }
    }
}
