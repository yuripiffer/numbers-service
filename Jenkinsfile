pipeline {
    agent any

    environment {
        AWS_REGION = "eu-north-1"
        ECR_REGISTRY = "913704463047.dkr.ecr.${AWS_REGION}.amazonaws.com"
        IMAGE_NAME = "numbers-service"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Feature Branch') {
            when {
                not {
                    anyOf {
                        branch 'staging'
                        branch 'main'
                    }
                }
            }
            stages {
                stage('Unit Tests') {
                    steps {
                        sh 'go test ./...'
                    }
                }
                stage('Build Docker Image') {
                    steps {
                        sh 'docker build -t $IMAGE_NAME .'
                    }
                }
                stage('Push to ECR') {
                    steps {
                        script {
                            sh '''
                                aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $ECR_REGISTRY
                                docker tag $IMAGE_NAME:latest $ECR_REGISTRY/$IMAGE_NAME:latest
                                docker push $ECR_REGISTRY/$IMAGE_NAME:latest
                            '''
                        }
                    }
                }
            }
        }

        stage('Deploy to Staging') {
            when {
                branch 'staging'
            }
            steps {
                sh '''
                    helm upgrade --install numbers-service helm/numbers-service \
                      --namespace numbers-staging --create-namespace \
                      --set image.repository=$ECR_REGISTRY/$IMAGE_NAME \
                      --set image.tag=latest
                '''
            }
        }



        stage('Deploy to Production') {
             when {
                branch 'main'
            }
            steps {
                sh '''
                    helm upgrade --install numbers-service helm/numbers-service \
                      --namespace numbers-prod --create-namespace \
                      --set image.repository=$ECR_REGISTRY/$IMAGE_NAME \
                      --set image.tag=latest
                '''
            }
        }
    }
}
