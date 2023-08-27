export GOBIN=$(PWD)/bin
GO_PATH=$(go env GOPATH)
export PATH=$GOBIN:$GO_PATH:$PATH

OUTPUT_PATH_KEY="output"
OUTPUT_PACKAGE_KEY="pkgname"
TEMPLATE_KEY="templates"
ENTITIES_OUTPUT_PATH="app\/domains\/entities"
ENTITIES_OUTPUT_PACKAGE="entities"
ENTITIES_TEMPLATE='["db\/sqlboiler\/templates"]'
REPOSITORIES_OUTPUT_PATH="app\/repositories\/entities"
REPOSITORIES_OUTPUT_PACKAGE="entities"
REPOSITORIES_TEMPLATE=\
"['$(echo $GOPATH | cut -d: -f1 | sed 's/\//\\\//g')\/pkg\/mod\/github.com\/volatiletech\/sqlboiler\/v4@v4.14.2\/templates\/main',\
'$(echo $GOPATH | cut -d: -f1 | sed 's/\//\\\//g')\/pkg\/mod\/github.com\/volatiletech\/sqlboiler\/v4@v4.14.2\/templates\/test']"

SQLBOILER_TOML=$(cat db/sqlboiler/sqlboiler.toml)

# sqlboiler.toml内の値を変更(entityを生成するための設定)
MODIFIED_SQLBOILER_SETTING=$(echo "$SQLBOILER_TOML" | sed -e "s/\($OUTPUT_PATH_KEY=\)\".*\"/\1\"$ENTITIES_OUTPUT_PATH\"/")
MODIFIED_SQLBOILER_SETTING=$(echo "$MODIFIED_SQLBOILER_SETTING" | sed -e "s/\($OUTPUT_PACKAGE_KEY=\)\".*\"/\1\"$ENTITIES_OUTPUT_PACKAGE\"/")
MODIFIED_SQLBOILER_SETTING=$(echo "$MODIFIED_SQLBOILER_SETTING" | sed -e "s/\($TEMPLATE_KEY=\)\[.*\]/\1$ENTITIES_TEMPLATE/")

TMP_CONFIG=$(mktemp /tmp/sqlboiler4.toml)
echo "$MODIFIED_SQLBOILER_SETTING" > $TMP_CONFIG
$GOBIN/sqlboiler --config=$TMP_CONFIG mysql

rm $TMP_CONFIG

# sqlboiler.toml内の値を変更(repositoryを生成するための設定)
MODIFIED_SQLBOILER_SETTING=$(echo "$SQLBOILER_TOML" | sed -e "s/\($OUTPUT_PATH_KEY=\)\".*\"/\1\"$REPOSITORIES_OUTPUT_PATH\"/")
MODIFIED_SQLBOILER_SETTING=$(echo "$MODIFIED_SQLBOILER_SETTING" | sed -e "s/\($OUTPUT_PACKAGE_KEY=\)\".*\"/\1\"$REPOSITORIES_OUTPUT_PACKAGE\"/")
MODIFIED_SQLBOILER_SETTING=$(echo "$MODIFIED_SQLBOILER_SETTING" | sed -e "s/\($TEMPLATE_KEY=\)\[.*\]/\1$REPOSITORIES_TEMPLATE/")

TMP_CONFIG=$(mktemp /tmp/sqlboiler6.toml)
echo "$MODIFIED_SQLBOILER_SETTING" > $TMP_CONFIG
$GOBIN/sqlboiler --config=$TMP_CONFIG mysql

rm $TMP_CONFIG

goimports -w $(echo "$ENTITIES_OUTPUT_PATH" | sed 's/\\//g')
goimports -w $(echo "$REPOSITORIES_OUTPUT_PATH" | sed 's/\\//g')
