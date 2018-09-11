
#!/bin/bash
SERVICE_DIR=$1
echo "Uploading for image "+$SERVICE_DIR
echo "$DOCKERHUB_PASSWORD" | docker login -u "$DOCKERHUB_USER" --password-stdin
cd $SERVICE_DIR
export IMAGE_NAME=stmalike/$SERVICE_DIR:latest
docker build -f Dockerfile -t $IMAGE_NAME .
docker tag $IMAGE_NAME:$COMMIT $IMAGE_NAME
docker tag $IMAGE_NAME:$COMMIT $IMAGE_NAME:travis-$TRAVIS_BUILD_NUMBER
docker push $IMAGE_NAME

  