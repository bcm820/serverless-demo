Key points:
    cloud agnostic
    portable
    cheap
    customize runtime (usefull: DEMO? to pull npm packages in nexus? for highside relevance)

Suported Runtime varriants:
- NodeJS
- Python
- Ruby
- Golang
- Java
- .NET Core (C#)
- Ballerina
- Custom runetime (wow): benefit of this is we can modify existing runtimes too.
        to add an npm package to the runtime add the following files in `docker/runtime`:
            ```
            RUN apt-get update && apt-get install git
            RUN npm install -g lodash
            ```
        Then use the Makefile to generate the base image
        tag the image and push it
        add this newly tagged docker image to the kubeless config
        Finally deploy the function with the new runtime to restart kubeless controller pod

Risks: 

SWOT analysis:

    Strengths:



    Weaknesses:



    Opportunities:




    Threats:
        function warmup? does it have this like lambda



