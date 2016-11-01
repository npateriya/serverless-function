## Build and Run UI
UI to connect and invoke serverless function REST Api.


### Build  and Publish docker container

- Build
    - docker build -t npateriya/serverless-function-ui:preview001
- Publish 
	- docker push  npateriya/serverless-function-ui:preview001


### Running UI container
- docker run --name some-nginx -d -p 8080:80 npateriya/serverless-function-ui:preview001
  - This will start docker container and UI will accessible at http://<dockerhost>:8080

