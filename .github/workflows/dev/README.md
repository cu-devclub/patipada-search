# Important Notes

This directory serve github action files for only development using `act`, some job or variable may be outdated due to the change of the production

-- The main different is how to ssh to server as the problem of rsa private key do not work with local environment (work only in github secret environment somehow) so in local we need to use base64 private key instead

** Please verify the code structure with workflow in `.github/workflows` as the jobs in that directory will be the one which currently updated.