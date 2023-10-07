############################
# STEP 1 build executable binary
############################
FROM golang as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apt update && apt install -y sudo git ca-certificates && update-ca-certificates
# Create appuser
ENV USER=netflix
ENV UID=1000
ENV GID=1000
ENV GOFLAGS="-buildvcs=false"
RUN addgroup --gid $GID netflix && \
    adduser --uid $UID --gid $GID --disabled-password --gecos "" netflix && \
    echo 'netflix ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers
COPY . .
# Build the binary
RUN go build -trimpath -o /go/bin/netflix
############################
# STEP 2 build a small image
############################
FROM scratch
# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable
COPY --from=builder /go/bin/netflix /go/bin/netflix
# Use an unprivileged user.
USER netflix:netflix
# Run the netflix binary.
ENTRYPOINT ["/go/bin/netflix"]