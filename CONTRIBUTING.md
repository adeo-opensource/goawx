# How to contribute

If you would like to help us and contribute to this project, you are very welcome!
However, some rules must be respected in order to keep a maximum of consistency and stability within the repo.  

## Things to know before submitting code

### Good practices

* All code submission are done through pull requests against the main branch
* [Create an issue][open-issue] (pull request or bug report) before submitting your pull request
* [Associate issue and pull request][link-pr-with-issue]

### Pre commit hooks

When attempting to perform a git commit, a pre-commit hook is executed before the commit is allowed in your local repository.

## Submit Pull requests

Fixes and features will go through the GitHub pull request process.
Submit your pull request (PR) against the `main` branch.

Before submitting your pull request, make sure you validate the style and tests.
To do this, install the prerequisites, see [DEVELOPERS](DEVELOPERS.md) documentation.
The pre-commit and tests are played during the push.

* Write tests for new features and update/add tests for bug fixes.
* Make the smallest possible change per pull request
* Write good validation messages. [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/)

## Reporting issues

Use the GitHub issue tracker for filing bugs.
In order to save time, and help us respond to issue quickly, make sure to fill out as much of the issue template as possible.
Version information, and an accurate reproducing scenario are critical to helping us identify the problem.

### Issue states

`state:needs_triage` This issue has not been looked at by a person yet and still needs to be triaged.
This is the initial state for all new issues/pull requests.

`state:needs_info` The issue needs more information.
This could be more debug output, more specifics out the system such as version information.
Any detail that is currently preventing this issue from moving forward. This should be considered a blocked state.

`state:needs_review` The issue/pull request needs to be reviewed by other maintainers and contributors.
This is usually used when there is a question out to another maintainer or when a person is less familar with an area of the code base the issue is for.

`state:in_progress` The issue is actively being worked on and you should be in contact with who ever is assigned.
It's important if you are also working on or plan to work on a similar issue.

`state:in_testing` The issue or pull request is currently being tested.


[link-pr-with-issue]: https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword
[open-issue]: https://github.com/adeo-opensource/terraform-provider-awx/issues/new/choose
