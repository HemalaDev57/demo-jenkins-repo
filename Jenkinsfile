pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "your-docker-username/preprod-demo-service"
        DOCKER_TAG = "latest"
        DOCKER_REGISTRY = "docker.io"
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
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin $DOCKER_REGISTRY
                        docker build -t $DOCKER_IMAGE:$DOCKER_TAG .
                        docker push $DOCKER_IMAGE:$DOCKER_TAG
                        DIGEST=$(docker inspect --format='{{index .RepoDigests 0}}' $DOCKER_IMAGE:$DOCKER_TAG)
                        echo "DOCKER_DIGEST=$DIGEST" >> env.properties
                    '''
                }
                script {
                    def props = readProperties file: 'env.properties'
                    env.DOCKER_DIGEST = props['DOCKER_DIGEST']
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
