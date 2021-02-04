# Container Solutions Hiring API Test

This a functional test used by [Container-Solutions](https://www.container-solutions.com/careers) to see whether you are ready to join our awesome team of Containerds.

## Terms and Conditions

TODO: Add T's & C's

## Just a Note

While our automation on this pipeline is just a quick test, your submission will be viewed by our engineers to check for best practices. So if you have anything to add to wow us, please add it.

## Steps

### Fork this repository

Please fork the repository and do the test on your fork.

There are a couple of files that you may not change, these include:

- anything in the `.github/` directory
- anything in the `.ci/` directory

### Do The API Excercise

The API Excercise is designed as a very simple way for us to understand your skill set. While we fully understand that it is difficult to have a test that encompasses the wide variety of skill sets in the technology industry, this test checks the aspects that you are most likely to find working at Container Solutions. We have some basic steps that you need to adhere to but otherwise you are free to set this project up as you see fit. For the specifications please view the [API_SPECIFICATIONS.md](./API_SPECIFICATIONS.md)

Once done there are a few files that you will need to update, these files will be run by the CI/CD pipeline in order to verify your Excercise:

- [build.sh](./build.sh): This file is used to build your application's Docker image, please update the command in here to reflect your setup. Also a note, in the CI/CD pipeline we have a local docker registry `kind-registry:5000`, so please use this image registry when tagging your image.
- [push.sh](./push.sh): This file is used to push your image to the image registry found at `kind-registry:5000`.
- [deploy.sh](./deploy.sh`): This file is used to deploy your kubernetes manifests.

_Note_ the CI/CD pipeline will run a test suite against the endpoints, please use the domain `api.awesome` in your kubernetes routing.

## CI/CD Kubernetes Environment

To run this test we use [kind](https://kind.sigs.k8s.io/) to run the setup, we have setup an NGINX ingress according to [this](https://kind.sigs.k8s.io/docs/user/ingress/#ingress-nginx). We will look at supporting another ingress in the future. Currently we have [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), [helm](https://helm.sh/) and [kustomize](https://kustomize.io/) installed in our pipeline. We are always interested in new and exciting tools to manage kubernetes resources, so please if you want to see a tool add it in an issue on this repo.

### Open A Pull Request

Once you are happy with your approach you can open a Pull Request, if the CI/CD pipeline passes and none of the directories that are specified in step **Fork this repository** have been altered, your description of your PR will be sent to the Container Solutions talent team along with your Github username.

### What's next?

If you are already in contact with Container Solutions Talent Team, then you are more than welcome to reach out to your contact and let them know you have completed the API test, if not then you can do one of two things:

- If your email address is on your Github user profile and you meet the candidate requirements specified on our [Container Solutions Career Page](https://www.container-solutions.com/careers) our talent team will reach out to you with regards to your availability.
- Alternatively you can apply to one of our positions on our [Container Solutions Career Page](https://www.container-solutions.com/careers).
