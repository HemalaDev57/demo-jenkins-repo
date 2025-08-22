pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "hddocker125/demo-cbci-test"
        DOCKER_TAG = "latest"
        DOCKER_REGISTRY = "docker.io"
    }

    stages {
        stage('Build & Docker Push') {
            agent {
                docker {
                    image 'docker:24.0-dind'
                    args '--privileged -v $WORKSPACE:/workspace -w /workspace'
                }
            }
            steps {
                echo 'Building Go Application and Docker image...'

                // Install Go inside DinD container
                sh '''
                    apk add --no-cache go git bash
                    go version
                '''

                // Build Go app
                sh '''
                    go mod tidy
                    go build -o app .
                '''

                // Docker login, build and push
                withCredentials([usernamePassword(credentialsId: 'docker-cred', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin $DOCKER_REGISTRY
                        docker build -t $DOCKER_IMAGE:$DOCKER_TAG .
                        docker push $DOCKER_IMAGE:$DOCKER_TAG
                    '''
                }

                // Capture Docker digest
                script {
                    env.DOCKER_DIGEST = sh(
                        script: "docker inspect --format='{{index .RepoDigests 0}}' $DOCKER_IMAGE:$DOCKER_TAG | cut -d'@' -f2",
                        returnStdout: true
                    ).trim()
                    echo "Docker Digest: ${env.DOCKER_DIGEST}"
                }
            }
        }

        stage('Register Build Artifact') {
            agent any
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
                docker {
                    image 'golang:1.25'
                    args '-v $WORKSPACE:/workspace -w /workspace'
                }
            }
            steps {
                echo 'Runnin
