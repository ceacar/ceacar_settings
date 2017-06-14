FROM ubuntu
RUN apt-get update
RUN apt-get install -y python
RUN apt-get install -y python3
#make port 10000 to outside world
Expose 10000
ENV var_name var_value



#two form of CMD
#shell form like CMD ping localhost  --> /bin/sh ping localhost
#exec form like CMD ["/bin/ping","localhost"]  --> /bin/ping localhost, this carries no surprises
CMD [ "python3 -c 'print(\"python3\", \"ready\")'"]

