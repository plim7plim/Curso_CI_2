FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV HOST=localhost PORT=5432
<<<<<<< HEAD

=======
>>>>>>> 3ec0fcc45b1353a0c45a67e4249a857d83244eb4
ENV USER=root PASSWORD=root DBNAME=root

COPY ./main main

<<<<<<< HEAD
CMD [ "./main" ]
=======
CMD ["./main"]
>>>>>>> 3ec0fcc45b1353a0c45a67e4249a857d83244eb4
