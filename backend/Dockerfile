FROM golang:1.22-alpine

# RUN apk add --no-cache \
#     poppler-utils \
#     tesseract-ocr \
#     tesseract-ocr-data-eng \
#     tesseract-ocr-data-fra \
#     tesseract-ocr-data-deu \
#     tesseract-ocr-data-spa \
#     ocrmypdf
    
# Set environment variables for Tesseract and OCRmyPDF
# ENV TESSDATA_PREFIX=/usr/share/tessdata
# ENV LC_ALL=C

# Verify Tesseract installation and available languages
# RUN tesseract --list-langs

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go *.mod *.sum .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /pollis-go-app

EXPOSE 8080

CMD ["/pollis-go-app"]