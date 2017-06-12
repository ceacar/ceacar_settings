FROM ubuntu
RUN apt-get update
RUN apt-get install -y python
RUN apt-get install -y python3
#make port 10000 to outside world
Expose 10000
ENV var_name var_value

CMD [ "python3 -c 'print(\"python3\", \"ready\")'"]

