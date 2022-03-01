FROM --platform=${TARGETPLATFORM} alpine:3 AS executor

# copy over the binary from the first stage
COPY megabot /megabot/megabot

WORKDIR "/megabot"
ENTRYPOINT [ "/megabot/megabot" ]