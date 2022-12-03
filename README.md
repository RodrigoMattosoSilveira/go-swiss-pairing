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
* [x] Add the react-grpc-go UI to go-swiss-pairing, replacing pingpong PingPong with the Member ping call
* [x] Implement the [Hands On React Course](https://handsonreact.com/docs/) - I had to do it to re-learn React;
* [ ] Re-factor the server to support all Member attributes
* [ ] Use the [Hands On React Course](https://handsonreact.com/docs/) framework process Members; it includes extensive UI tests;
* [ ] Migrate the Ui to use [React Material](https://mui.com/material-ui)
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
* [GOLANG Custom Errors](https://golangbot.com/custom-errors/)
* [babel](https://babeljs.io/docs/en/)
* [react hands on tutorial](https://handsonreact.com/docs/props#!)
* [Regular Expressions 101](https://regex101.com)
* [Service Worker](https://developer.mozilla.org/en-US/docs/Web/API/Service_Worker_API) -  act as proxy servers that sit between web applications, the browser, and the network (when available)
* [IntelliJ File Watchers](https://www.jetbrains.com/help/idea/using-file-watchers.html)
* User Experience
  * Libraries
    * [React Bootstrap](https://react-bootstrap.github.io)
    * [React Material](https://mui.com/material-ui)
    * [mini.css](https://minicss.us/docs.htm)
    * [Material Design for Bootstrap 5 & React 18](https://mdbootstrap.com/docs/react/#demo)
    * [React Bootstrap vs React Material](https://www.upgrad.com/blog/bootstrap-vs-material/)
  * CSS Grids
    * [An Introduction to CSS Grid Layout (with Examples)](https://www.freecodecamp.org/news/intro-to-css-grid-layout/) 
    * [Table with CSS Grid](https://stackoverflow.com/questions/68141663/table-with-css-grid)
    * [learn css grid](https://learncssgrid.com/)
    * [mincss Grid system](https://minicss.us/docs.htm#grid)
  * [React Query](https://tanstack.com/query/v4)  
  * [ARIA in HTML](https://www.w3.org/TR/html-aria/#docconformance) - Consulted it to find the `<div />` role
* Testing
  * Server
    * GINKGO
      * [Ginkgo](https://onsi.github.io/ginkgo/#running-specs)
      * **All TESTS** Execute the following command: `$ ginkgo ./...`
      * **Specific suite** Execute the following command: `ginkgo --focus "test 1" testDir`
    * [Password Validator for GO](https://libraries.io/go/github.com%2Fgo-passwd%2Fvalidator)
  * UI
    * [Testing Library](https://testing-library.com/) - a light-weight solution for testing web pages by querying and interacting with DOM nodes;
    * [Jest 14.0: React Tree Snapshot Testing](https://jestjs.io/blog/2016/07/27/jest-14#why-snapshot-testing) 
    * [Test Renderer](https://reactjs.org/docs/test-renderer.html)
    * [Functional Testing using React testing library and Jest](https://vijayt.com/post/functional-testing-using-react-testing-library-and-jest/#:~:text=We%20have%20seen%20the%20getByRole%20function.%20It%20retrieves,the%20element%20is%20not%20rendered%20in%20the%20DOM.)
    * [Common mistakes with React Testing Library](https://kentcdodds.com/blog/common-mistakes-with-react-testing-library)
    * [Mock Service Worker](https://mswjs.io/) - Mock by intercepting requests on the network level.
    * [How can I mock or simulate gRPC APIs?](https://stackoverflow.com/questions/52919769/how-can-i-mock-or-simulate-grpc-apis)
 
  

Card | Table
---|-------
MembersPage | MembersPage
MembersList | MembersList
MemberCard | MembersListRow
MemberForm | MemberFormModal