# Gofiler - File Sharing System

Users want to share their files these days, so I developed a file sharing system called `gofiler` and users can share their files using this system easily from command line. This system has two main parts, which are `client` and `server`. To share a file, users can upload their file to the server using the provided command line interface and the server will generate a unique code for the uploaded file, so users can share their files by sharing this code to each other.
This system provides four functionalities which are `upload`,`download`,`delete` and `info`.

### Upload file
This feature allows the user to upload and share any file in such a way that the user can choose a file and the client sends the file to the server through TCP connection, then the server will return a unique share code corresponding to uploaded file.

### Download file
This feature allows the user to download the shared file, and by entering the unique share code, the user can receive the desired file from the server.

### File Info
user can get the file information by entering the file sharing code. The file information includes: `file name`, `extension`, `size` and `checksum`.

### Delete File
This functionality allows the user to delete the file from the server in such a way that the user sends the share code to the server through the client and server will delete the file if the share code already exists.

Client and server are written in Go programming language. Client communicates to the server over TCP socket and the server is multi thread thanks to groutins and go standard library, so it can handle multiple client connection at the same time.

## Challenges and solutions
In the process of developing this system, there were various challenges, which are:

### User interface for client and server
User interaction is always challenging to meet their needs, the system provides a command line interface, so the end user can access the client and its functionalities and the system admin can configure the server using cli
We also used an open source library for cli implementation, which is:
https://github.com/urfave/cli

### Client/server communication
To transfer structured messages over `TCP` connection, we need to serialize data, so protocol buffer (protobuf) is the best solution for this challenge because it can convert messages to binary data bytes.
We also need to implement a custom protocol to transfer data bytes to answer this need, the desired protocol has been implemented as follows:
| size in bytes  | description |
| ------------- | ------------- |
| 1 byte  | message type  |
| 2048 - 1 byte  | payload binday data |

### file storage
Data storage has always been a challenge, specially in server side. In file sharing systems, a server needs to store uploaded files, and it's information, so we used the OS file system to save files also a separate package developed for this purpose which provides 4 functionalities for managing files on file system such as: save file, remove file, get file and file info. Each file has a unique share code which will be saved beside of its original name on the file system, for example: 
`AB78C9GR6M_hello.txt`

### server logging
To log incoming requests details and server activities, our system need a level based logging system so an open source logging library used for this purpose which is: https://github.com/uber-go/zap

## Testing
All functionalities in software development process have been tested manually instead of writing unit tests by integrating client to server and sending requests.
For example, to test the upload file functionality, a test file is uploaded by the client, and the correctness of file transferring will be ensured.
This method has been used to test all other functionalities.