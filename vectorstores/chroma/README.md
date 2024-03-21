# Chroma Support

<<<<<<< HEAD
You can access Chroma via the included implementation of the 
=======
You can access Chroma via the included implementation of the
>>>>>>> upstream/main
[`vectorstores.VectorStore` interface](../vectorstores.go)
by creating and using a Chroma client `Store` instance with
the [`New` function](./chroma.go) API.

## Client/Server

<<<<<<< HEAD
Until 
[an "in-memory" version](https://docs.trychroma.com/usage-guide#running-chroma-in-clientserver-mode)
is released, only client/server mode is available.

=======
Until
[an "in-memory" version](https://docs.trychroma.com/usage-guide#running-chroma-in-clientserver-mode)
is released, only client/server mode is available.

> **Note:** Additional ways to run Chroma locally can be found
> in [Chroma Cookbook](https://cookbook.chromadb.dev/running/running-chroma/)

>>>>>>> upstream/main
Use the [`WithChromaURL` API](./options.go) or the `CHROMA_URL` environment
variable to specify the URL of the Chroma server when creating the client instance.

## Using OpenAI LLM

<<<<<<< HEAD
To use the OpenAI LLM with Chroma, use either the 
[`WithOpenAiAPIKey` API](./options.go) or the `OPENAI_API_KEY` environment
=======
To use the OpenAI LLM with Chroma, use either the
[`WithOpenAIAPIKey` API](./options.go) or the `OPENAI_API_KEY` environment
>>>>>>> upstream/main
variable when creating the client.

## Running With Docker

Running a Chroma server in a local docker instance can be especially useful for testing
<<<<<<< HEAD
and development workflows.  An example invocation scenario is presented below:

### Starting the Chroma Server

As of this writing, the newest release of the Chroma docker image is 
[chroma:0.4.9](https://github.com/chroma-core/chroma/pkgs/container/chroma/125222480?tag=0.4.9).
=======
and development workflows. An example invocation scenario is presented below:

### Starting the Chroma Server

As of this writing, the newest release of the Chroma docker image is
[chroma:0.4.24](https://github.com/chroma-core/chroma/pkgs/container/chroma/184319417?tag=0.4.24).
>>>>>>> upstream/main
Running it directly while exposing its port to your local machine can be
accomplished with:

```shell
<<<<<<< HEAD
$ docker run -p 8000:8000 ghcr.io/chroma-core/chroma:0.4.9
=======
$ docker run -p 8000:8000 ghcr.io/chroma-core/chroma:0.4.24
>>>>>>> upstream/main
```

### Running an Example `langchaingo` Application

<<<<<<< HEAD
With the "Simple Docker Server" running (see above), running the included 
=======
With the "Simple Docker Server" running (see above), running the included
>>>>>>> upstream/main
example `langchaingo` app should produce the following result:

```shell
$ export CHROMA_URL=localhost:8000
$ export OPENAI_API_KEY=YourOpenApiKeyGoesHere
$ go run ./examples/chroma-vectorstore-example/chroma_vectorstore_example.go
Results:
1. case: Up to 5 Cities in Japan
    result: Tokyo, Nagoya, Kyoto, Fukuoka, Hiroshima
2. case: A City in South America
    result: Buenos Aires
3. case: Large Cities in South America
    result: Sao Paulo, Rio de Janeiro
```

## Tests

The test suite `chroma_test.go` started as a clone of the adjacent `pinecone_test.go`,
<<<<<<< HEAD
and is initially quite sparse.  Consider contributing new test cases, or adding
=======
and is initially quite sparse. Consider contributing new test cases, or adding
>>>>>>> upstream/main
coverage to accompany any changes made to the code.
