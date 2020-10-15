FROM scratch

COPY bin/desafio-go /desafio-go

ENTRYPOINT ["/desafio-go"]