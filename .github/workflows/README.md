# Project Workflow and Variables

This project leverages GitHub Actions for streamlined automation. While most variables are securely stored as GitHub Secrets for heightened security, there's an exception.

# Important Note

One crucial variable, `LINUX_HOST`, is intentionally exposed as a GitHub Secret. Despite its sensitive nature, this approach is adopted to simplify customization for developers cloning this project. The decision stems from the project's non-commercial orientation.

# Why?

The rationale behind this choice is to empower developers with a seamless experience. By making `LINUX_HOST` a secret, you can effortlessly clone this project and tailor the variable directly in your GitHub repository settings. This eliminates the need for code modifications for basic configuration adjustments.


# How it works 

every workflow was set to on `workflow_dispatch` means you have to navigate to `Action` Tabs in Github website, in this repository and press `Run` to run each workflow.

This is done because not every service need to re-deploy when push to some branch 

** Further improve by testing when push to branch 

# Secret variables to create in github secrets

1. DOCKERHUB_USERNAME
2. DOCKERHUB_TOKEN
3. LINUX_HOST
4. LINUX_USERNAME
5. LINUX_PRIVATE_KEY_BASE64
6. ELASTIC_PASSWORD
7. ELASTIC_USERNAME
8. AUTH_DB_USER
9. AUTH_DB_PASSWORD
10. JWT_KEY
11. SENDER_PASSWORD