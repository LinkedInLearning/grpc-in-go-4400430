# gRPC in Go
This is the repository for the LinkedIn Learning course gRPC in Go. The full course is available from [LinkedIn Learning][lil-course-url].

![gRPC in Go][lil-thumbnail-url] 

A growing number of companies use gRPC, a popular framework that writes services and clients. In this course, Miki Tebeka presents an overview of gRPC, then shows you the protocol buffers serialization format, as well as how to write a gPRC definition file (.proto), write gRPC servers and clients, and use advanced gRPC features like streaming, reflection, and more. Learn about RPC in general, then go into protocol buffers, how gRPC use them, why gRPC chose HTTP/2, and how the gRPC ecosystem works. Explore writing and compiling a .proto file, using generated code, and generating JSON encoding. Dive into gRPC servers and clients, and discover how to define a streaming endpoint, send streaming data to the client, and handle streaming responses from the server. Plus, go over advanced topics like writing interceptors, sharing .proto definitions, testing your code, and using the gRPC gateway.



## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go SDK](https://go.dev/dl)
    - [Visual Studio Code](https://code.visualstudio.com/) with the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
    - [git](https://git-scm.com/)
    - [protoc](https://grpc.io/docs/protoc-installation/)
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.


### Instructor

Miki Tebeka 
                            


                            

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/miki-tebeka).

[lil-course-url]: https://www.linkedin.com/learning/grpc-in-go?dApp=59033956&leis=LAA
[lil-thumbnail-url]: https://media.licdn.com/dms/image/D560DAQENGbFJUcfkbg/learning-public-crop_288_512/0/1689793845488?e=2147483647&v=beta&t=kmn30g1oiW0MERiaon5lO1OJ9VKlajiVlyBRwTFSkwE

