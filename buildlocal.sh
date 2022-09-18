#!/bin/bash

if [ -z "$repo" ]; then
    repo=$PWD
fi

if [ -z "$packageVersion" ]; then
    packageVersion=0.0.0
fi

generatorUrl=$(cat ${repo}/bindgen-config.json | json generate.generatorUrl)
customGeneratorUrl=$(cat ${repo}/bindgen-config.json | json generate.customGeneratorUrl)

#echo generatorUrl = ${generatorUrl}
#echo customGeneratorUrl = ${customGeneratorUrl}

# Pull down the latest version of the generator
rm -f ${customGeneratorUrl} ${generatorUrl}
wget -q -O ${repo}/go-oapi-codegen.jar ${customGeneratorUrl}
wget -q -O ${repo}/openapi-generator-cli.jar ${generatorUrl}

cat ${repo}/openapi.json > ${repo}/openapi.json.tmp
preprocessCount=$(cat ${repo}/bindgen-config.json | json generate.preprocess.length)
for (( i=0; i<$preprocessCount; i++))
do
    current=$(cat ${repo}/bindgen-config.json | json generate.preprocess.$i)
    key=$(echo "${current}" | json key)
    type=$(echo "${current}" | json type)
    value=$(echo "${current}" | json value)
    valueIsString=$(echo "${current}" | json -e "console.log(typeof this.value === 'string')" | head -n 1)
    if [ $valueIsString = true ]; then
        value='"'${value}'"'
    fi
    if [ $type = remove ]; then
        json -I -f ${repo}/openapi.json.tmp -e 'try { this.'"${key}"'=undefined; } catch(err) { console.log("Could not remove key: '${key}'"); }'
        elif [ $type = update ]; then
        json -I -f ${repo}/openapi.json.tmp -e 'try { this.'"${key}"'='"${value}"'; } catch(err) { console.log("Could not update key: '${key}'"); }'
    else
        echo Unknown preprocessor replacement type "${type}" for key "${key}"
        exit 1
    fi
done
changedVersion=$(cat ${repo}/openapi.json | json info.version)
echo '::set-output name=change::'"${changedVersion}"
echo '::set-output name=random-ext::'"${RANDOM}"

java -cp ${repo}/go-oapi-codegen.jar:${repo}/openapi-generator-cli.jar org.openapitools.codegen.OpenAPIGenerator generate -i ${repo}/openapi.json.tmp -g go-oapi-codegen -o ${repo}/onshape --type-mappings DateTime=JSONTime --additional-properties=packageVersion=${packageVersion} --additional-properties=useOneOfDiscriminatorLookup=true -c ${repo}/openapi_config.json
go fmt ${repo}/onshape
