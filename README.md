# Grace Hopper Celebration 2018 - Find Issues

This is the TDD in OSS workshop repo for GHC 18.

## Maintainers

1. [Genevieve L'Esperance](github.com/genevieve)
1. [Samantha Langit](github.com/samanthalangit)
1. [Angela Chin](github.com/chinangela)

## Setup

1. [Install `go`.](https://golang.org/doc/install#install)

1. [Install `ginkgo`.](https://onsi.github.io/ginkgo/#getting-ginkgo) This is our testing framework.

1. [Fork the repo in github](https://help.github.com/articles/fork-a-repo/#fork-an-example-repository) to your account.

1. Clone the repo.

    ```bash
    git clone git@github.com:YOUR-ACCOUNT/find-issues
    ```

1. Run the tests!

    ```bash
    ginkgo -r -race -parallel .
    ```

## The Exercise

1. Pick an issue from the [issues page](https://github.com/ghc-tdd/find-issues/issues).

1. Using a text editor or IDE, create a file in the `acceptance/` directory.

1. Write the acceptance test for the new feature.

1. Run the test.

    ```bash
    ginkgo -r -race -parallel acceptance/
    ```

1. Create a file in the appropriate package directory.

1. Write the unit test.

1. Run the test. It will fail to compile! That's okay!

1. Create a file to write the code. Write the least amount you think
you need to get the test to pass.

1. Run the test and fix the code until it goes green!

1. Run the acceptance test and fix the code until it goes green.

1. Refactor if need be.

1. Push the code to a branch or fork of the repo.

1. Create a pull request that references the open issue.


:raised_hands:**You're done!**:raised_hands:

## Questions?

If you have any questions during the workshop, open
a github issue on this repo. You can answer other
people's questions/issues or vote on issues that you want
to have answered or resolved.

