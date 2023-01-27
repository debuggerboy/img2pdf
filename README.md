# img2pdf

Images to PDF is a simple program to create PDF files from a series of image files.
We can use either "jpg/jpeg" or "png" images are input.

## Build Executable:

Create a module

```
go mod init img2pdf
```

Get the dependencies:

```
go get .
```

Build `img2pdf` executable binary.

```
go build img2pdf.go
```

## Invocation

We can run the `img2pdf` apl=plication from command-line as below:

```
./img2pdf --images=image1.jpg,image2.png,image3.jpeg --output demo.pdf
```

## To-Do

- add annotations
