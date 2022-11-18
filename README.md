# go-swiss-pairing
A tool for tournament directors to organize tournaments using the Swiss pairing method, written in Go

# What I want:
- [x] To use the Domain Driven Design approach and folder structure from go-swiss-pairing;
- To use the React / gRPC integration mechanism in react-grpc-go, but using the folder nexjs-ts-materio-template folder structure;
- To use the React Material components in nexjs-ts-materio-template

# The path to get what I want:
* [x] Refactor the go-swiss-pairing Makefile to use dependencies, see react-grpc-go for reference
* [x] Refactor the go-swiss-pairing club-member service name to member;
* [x] Add a ping service to MemberService;
* [x] Replace the go-swiss-pairing gRPC server (not service) with the react-grpc-go, service, including certificates, etc
* [ ] Add integration tests
    * [x] Add `server side` integration tests;
    * [ ] Add `client side` integration tests;
* [ ] Add the react-grpc-go UI to go-swiss-pairing, replacing pingpong PingPong with the Member ping call
* [ ] Refactor the UI to use the nexjs-ts-materio-template folder structure
* [ ] Replace the original simple react-grpc-go UI page with a simple nexjs-ts-materio-template component
* [ ] Design the landing page, the members page, including logic to CRUD members

# Add Members UI
* I had to run `yarn add create-react-app@latest` to then be able to run  `npx create-react-app swiss-pairing`
* [Small Go/React/Typescript gRPC-Web example](https://github.com/johanbrandhorst/grpc-web-go-react-example)- Interesting, but not based on the `improbable` `grpc-web` technology;
* [buf](https://docs.buf.build/introduction) A tool to facilitate PROTOBUF management; I'll only use the generation and dependencies features for now;
  * requires [./JQ](https://stedolan.github.io/jq/) 
* [Use gRPC with Node.js and Typescript](https://dev.to/devaddict/use-grpc-with-node-js-and-typescript-3c58) Does not show React
* [Small Go/React/TypeScript gRPC-Web example](https://github.com/johanbrandhorst/grpc-web-go-react-example)
* [buf](https://docs.buf.build/introduction) Integrated it, makes it easier to handle grpc

# Links and References
* Link to the [Recursive Wildcard Function](https://blog.jgc.org/2011/07/gnu-make-recursive-wildcard-function.html) I used to build a list of dependencies
* [Using gRPC with TLS, Golang and React (No Envoy)](https://itnext.io/using-grpc-with-tls-golang-and-react-no-envoy-92e898bf8463)
* [How to Set Up gRPC Server-Side Streaming with Go](https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/)
* [ginkgo](https://onsi.github.io/ginkgo/) - GOLANG test package that makes it similar to JASMINE and the like
* [GOLANG Custom Errors](https://golangbot.com/custom-errors/)
* [babel](https://babeljs.io/docs/en/)
* [mini.css](https://minicss.us/docs.htm)
* [react hands on tutorial](https://handsonreact.com/docs/props#!)

Card | Table
---|-------
MembersPage | MembersPage
MembersList | MembersListTable
MemberCard | MembersListTableRow
MemberForm | MemberFormModal