FROM python:3

RUN pip install kubernetes

ADD . /client
WORKDIR /client

CMD ["python", "./client.py"]

