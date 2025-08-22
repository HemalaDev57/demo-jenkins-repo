pipeline {
    agent any  // No global agent, each stage has its own

    environment {
        DOCKER_IMAGE = "hddocker125/demo-cbci-test"
        DOCKER_TAG = "latest"
        DOCKER_REGISTRY = "docker.io"
    }

    stages {
        stage('Build') {
            agent {
                docker { image 'golang:1.25' } // Go container
            }
            steps {
                echo 'Building Go Application...'
                sh 'go mod tidy'
                sh 'go build -o app .'
            }
        }

        stage('Docker Build & Push') {
            agent {
                docker {
                    image 'docker:24.0-dind'
                    args '--privileged'
                }
            }
            steps {
                echo 'Building and pushing Docker image...'
                withCredentials([usernamePassword(credentialsId: 'docker-cred', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
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
            agent any  // Any Jenkins node/container
            steps {
                echo "Registering Docker artifact..."
                script {
                    registerBuildArtifactMetadata(
                        name: "demo-cbci-test",
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
            agent {
                docker { image 'golang:1.25' }
            }
            steps {
                echo 'Running Unit Tests...'
                sh 'go test ./...'
            }
        }

        stage('Deploy') {
            agent any
            steps {
                echo 'Deploying...'
            }
        }
    }
}
