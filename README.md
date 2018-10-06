# Library for interacting and rendering reMarkable tablet notes

**In progress package development**

Tools and interfaces in Golang for the reMarkable tablet

The `rm` package helps parsing a reMarkable zip file in common golang structures.

Structure of a reMarkable zip file.

    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.content
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.lines
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.pagedata
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.pdf
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.thumbnails/0.jpg
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.thumbnails/1.jpg
    bdfe0e22-16dd-4f69-a7f3-86d8e2b47627.thumbnails/2.jpg

Then the `file` package can easily draw the reMarkable note previously parsed and rendered using one of the following formats.

 - PNG
 - PDF
 - SVG

An example of usage can be found in this [example file](file/draw_test.go).

### References

The description of each format (`content`, `lines`, `pagedata`) is explained really clearly on this [webpage](https://plasma.ninja/blog/devices/remarkable/binary/format/2017/12/26/reMarkable-lines-file-format.html).
