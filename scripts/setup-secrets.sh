cat ~/.pypi/pypi-key | gh secret set PYPI_API_TOKEN
cat ~/.pypi/test-pypi-key | gh secret set TEST_PYPI_API_TOKEN
cat ~/.docker/ghcr-token | gh secret set DOCKER_TOKEN
