<<<<<<< HEAD
module chroma-vectorstore-example

go 1.20

// NOTE: remove the following line to use the official (rather than local development) version
replace github.com/tmc/langchaingo => ../..

require (
	github.com/amikos-tech/chroma-go v0.0.0-20230901221218-d0087270239e
	github.com/google/uuid v1.3.0
	github.com/tmc/langchaingo v0.0.0-20231016073449-5620c5b08983
)

require (
	github.com/dlclark/regexp2 v1.8.1 // indirect
	github.com/pkoukk/tiktoken-go v0.1.2 // indirect
	golang.org/x/exp v0.0.0-20230510235704-dd950f8aeaea // indirect
)
=======
module github.com/tmc/langchaingo/examples/chroma-vectorstore-example

go 1.21

toolchain go1.21.4

require (
	github.com/amikos-tech/chroma-go v0.1.2
	github.com/google/uuid v1.6.0
	github.com/tmc/langchaingo v0.1.7
)

require (
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkoukk/tiktoken-go v0.1.6 // indirect
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1 // indirect
)

replace github.com/tmc/langchaingo => ../../
>>>>>>> upstream/main
