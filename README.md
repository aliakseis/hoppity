**Summary**

With creating REST API services with Go, developers are faced with the problems of documenting it. Swagger (Open API initiative) is a framework for describing your API and making it interactive.

It is possible to use only [go-swagger](https://github.com/go-swagger/go-swagger) to generate docs via golang code.

First you should install go-swagger. It is described how to do it properly in the official documents or in the README.md file. Then we will use the syntax of annotations - when swagger sees the comment section with the proper tag, the parser can understand that the tag is for swagger documentation.

go-swagger only generates specs for compiled code and code with swagger:meta tags. So, make your code compileable and DON&#39;T FORGET ABOUT swagger:meta tags. go-swagger will not generate specifications without that information.

[swagger:meta example](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/main.go#L1)

There are two ways to annotate your routes, swagger:operation and swagger:route.

Think of swagger:route as a short annotation for simple APIs.

[swagger:route example](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/health_check.go#L12)

Using swagger:operation annotations gets you access to all of the [OpenAPI specifications](https://swagger.io/specification/), letting you describe your complex endpoints. If you are interested in details, you can read the specification docs.

[swagger:operation example](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/hoppity_hop.go#L15)

**swagger:model**

A **swagger:model** annotation optionally gets a model name as extra data on the line. When this appears anywhere in a comment for a struct, then that struct becomes a schema in the definitions object of swagger.

The struct gets analyzed and all the collected models are added to the tree. The refs are tracked separately so that they can be renamed later on.

[swagger:model example](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/models/sequence_id.go#L4)

**swagger:response**

Reads a struct decorated with **swagger:response** and uses that information to fill in the headers and the schema for a response.

Swagger open source tool usage can be broken up into different use cases:

[swagger:response example](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/models/status.go#L4)

**Documenting APIs**

When described by an Open API document, Swagger open source tools may be used to interact directly with the API. It allows connecting directly to live APIs through an interactive, HTML-based user interface. Requests can be made directly from the UI and the options explored by the user of the interface.

[Swagger tool used to interact directly with the API](http://petstore.swagger.io/?url=http%3A%2F%2Flocalhost%3A3000%2Fswagger.json)

**Interacting with APIs**

Using the swagger-codegen project, end users generate client SDKs directly from the Open API document, reducing the need for human-generated client code.

**Developing APIs**

When creating APIs, Swagger tooling may be used to automatically generate an Open API document based on the code itself. This is informally called code-first or bottom-up API development.

1. Add Swagger tags and annotations as described above
2. [Update swagger.json and its dependents.](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/update_swagger_stuff.sh)
3. [Paste swagger.json contents to the Swagger Editor to verify it](https://editor.swagger.io/)
4. Fix Swagger annotations errors and return to item 2 if there are any.
5. Compile and run the application
6. [Check API functionality](http://petstore.swagger.io/?url=http%3A%2F%2Flocalhost%3A3000%2Fswagger.json)

**Notes on Swagger support in the application considering Chrome browser**

1. [Need to handle OPTIONS request for external tool to be able to invoke API](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/routes.go#L52)
2. [&quot;Access-Control-Allow-Origin&quot; header added for external tool to be able to access API](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/routes.go#L28)
3. [Expose swagger.json](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/routes.go#L77)
4. [Expose Redoc (built-in documentation viewer)](https://github.com/aliakseis/hoppity/blob/683b54f2e8be7c59e0109d5071582025a0101dff/src/hoppity/rest/routes.go#L71)
