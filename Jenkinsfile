/*
pipeline {

    agent any

    environment {
        GOCACHE = "${WORKSPACE}/.gocache"
    }

    stages {

        stage('Checkout') {
            steps {

                deleteDir()

                git branch: 'main',
                url: 'https://github.com/maneul0498-netizen/ms-test-ci-cd'
            }
        }

        stage('Build user-service') {

            agent {
                docker {
                    image 'golang:1.26'
                    reuseNode true
                }
            }

            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd user-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Build notification-service') {

            agent {
                docker {
                    image 'golang:1.26'
                    reuseNode true
                }
            }

            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd notification-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Run Containers') {

            steps {
                sh '''
                    docker-compose down || true

                    docker-compose up -d --build
                '''
            }
        }

        stage('Integration Test') {

            steps {
                sh '''
                    sleep 5

                    curl -X POST http://localhost:8081/users \
                    -H "Content-Type: application/json" \
                    -d '{"name":"manuel"}'
                '''
            }
        }
    }

    post {

        always {

            sh '''
                docker-compose logs

                docker-compose down
            '''
        }
    }
}
*/

pipeline {

    agent any

    environment {
        GOCACHE = "${WORKSPACE}/.gocache"
    }

    stages {

        stage('Checkout') {
            steps {

                deleteDir()

                git branch: 'main',
                url: 'https://github.com/maneul0498-netizen/ms-test-ci-cd'
            }
        }

        stage('Build user-service') {
            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd user-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Build notification-service') {
            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd notification-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Run Containers') {

            steps {
                sh '''
                    docker-compose down || true

                    docker-compose up -d --build

                    docker ps
                '''
            }
        }

        stage('Wait Services') {

            steps {
                sh '''
                    sleep 10

                    docker ps
                '''
            }
        }

        stage('Integration Test') {

            steps {
                sh """
                    curl -v http://host.docker.internal:8081/users
                """
            }
        }
    }

    post {

        always {

            sh '''
                docker ps

                docker-compose logs

                docker-compose down
            '''
        }
    }
}