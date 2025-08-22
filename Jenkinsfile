pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "hddocker125/demo-jenk-service"
        DOCKER_TAG = "latest"
        DOCKER_REGISTRY = "docker.io"
    }

    tools {
        go 'Go1.21' // name configured in Jenkins
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
                withCredentials([usernamePassword(credentialsId: 'H-docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                  script {
                    sh '''
                      echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin $DOCKER_REGISTRY
                      docker build -t $DOCKER_IMAGE:$DOCKER_TAG .
                      docker push $DOCKER_IMAGE:$DOCKER_TAG
                    '''
                    env.DOCKER_DIGEST = sh(
                      script: "docker inspect --format='{{index .RepoDigests 0}}' $DOCKER_IMAGE:$DOCKER_TAG | cut -d'@' -f2",
                      returnStdout: true
                    ).trim()
                  }
                }
            }
        }

        stage('Registering build artifact') {
            steps {
                echo "Registering Docker artifact..."
                script {
                    registerBuildArtifactMetadata(
                        name: "demo-jenk-artifact",
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
            steps {
                echo 'Deploying...'
            }
        }
    }
}
