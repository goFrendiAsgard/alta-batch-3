# Run sonar-qube server

```bash
docker run --rm \
    --name sonar-qube-server \
    -p 9000:9000 \
    sonarqube:community
```

Login to sonarqube server (http://localhost:9000)

* Default user: `admin`
* Default password: `admin`

Create auth token `User > My Account > Security`, click `Generate`.

After doing that, you should have an authentication token. Something like: `03600b63092ac8e8e65a4e5469886e61b0cf7871`

# Create sonar-project.properties

Create a file named `sonar-project.properties`.

```
sonar.projectKey=my:project
sonar.go.coverage.reportPaths=cover.out
```

# Run sonar-scanner

## Native

```bash
cd $HOME
# install
curl https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.7.0.2747-linux.zip
unzip sonar-scanner-cli-4.7.0.2747-linux.zip
mv sonar-scanner-cli-4.7.0.2747 sonar-scanner

# put this on ~/.bashrc or ~/.zshrc, or just run it
export PATH=$PATH:$HOME/sonar-scanner/bin

# on windows:
# $Env:PATH=($Env:PATH;$HOME/sonar-scanner/bin)
```

```bash
# run
cd sample
go test -coverprofile=cover.out ./...
sonar-scanner -Dsonar.host.url=http://localhost:9000 -Dsonar.login=03600b63092ac8e8e65a4e5469886e61b0cf7871
```

## Docker

```bash
export SONAR_LOGIN=03600b63092ac8e8e65a4e5469886e61b0cf7871
export SONARQUBE_URL=host.docker.internal:9000
export REPO=$(realpath sample)
docker run \
    --rm \
    -e SONAR_HOST_URL="http://${SONARQUBE_URL}" \
    -e SONAR_LOGIN="${SONAR_LOGIN}" \
    -v "${REPO}:/usr/src" \
    sonarsource/sonar-scanner-cli
```