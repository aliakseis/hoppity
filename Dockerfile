FROM  golang:1.10 as build-env


ADD . /wrk
WORKDIR /wrk/src/hoppity

ENV GOPATH /wrk
EXPOSE 3000

RUN echo "=======================" \
 && echo " Building Rest API " \
 && echo "=======================" 
RUN apt-get update \
    && cd /wrk/src/hoppity && make vendor && make


RUN echo "=======================" \
 && echo " Launching unit test" \
 && echo "======================="
RUN cd /wrk/src/hoppity && make test



RUN apt-get update && apt-get install -y curl 

CMD ./service
