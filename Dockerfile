FROM alpine
ADD user_srv /user_srv
ENTRYPOINT [ "/user_srv" ]
