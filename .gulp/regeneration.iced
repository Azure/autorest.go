
###############################################
# LEGACY
# Instead: have bunch of configuration files sitting in a well-known spot, discover them, feed them to AutoRest, done.

regenExpected = (opts,done) ->
  keys = Object.getOwnPropertyNames(opts.mappings)
  instances = keys.length

  for kkey in keys
    optsMappingsValue = opts.mappings[kkey]
    key = kkey.trim();

    swaggerFiles = (if optsMappingsValue instanceof Array then optsMappingsValue[0] else optsMappingsValue).split(";")
    args = [
      "--use=#{basefolder}"
      "--go",
      "--output-folder=#{opts.outputBaseDir}/#{key}",
      "--license-header=#{if !!opts.header then opts.header else 'MICROSOFT_MIT_NO_VERSION'}",
      "--enable-xml",
      "--package-name=#{opts.packageNameBase}/#{key}"
    ]

    for swaggerFile in swaggerFiles
      args.push("--input-file=#{if !!opts.inputBaseDir then "#{opts.inputBaseDir}/#{swaggerFile}" else swaggerFile}")

    if (opts.addCredentials)
      args.push("--go.add-credentials=true")

    if (opts.azureArm)
      args.push("--go.azure-arm=true")

    if (opts.fluent)
      args.push("--go.fluent=true")
    
    if (opts.syncMethods)
      args.push("--go.sync-methods=#{opts.syncMethods}")
    
    if (opts.flatteningThreshold)
      args.push("--go.payload-flattening-threshold=#{opts.flatteningThreshold}")

    args.push("--go.namespace=#{optsMappingsValue[1]}")

    if (opts['override-info.version'])
      args.push("--override-info.version=#{opts['override-info.version']}")
    if (opts['override-info.title'])
      args.push("--override-info.title=#{opts['override-info.title']}")
    if (opts['override-info.description'])
      args.push("--override-info.description=#{opts['override-info.description']}")

    autorest args,() =>
      instances--
      if instances is 0
        await execute "go fmt ./generated/...", { cwd: './test/src/tests' }, defer code, stderr, stdout
        return done()

goMappings = {
  'additionalproperties':['additionalProperties.json', 'additionalproperties'],
  'arraygroup':['body-array.json','arraygroup'],
  'booleangroup':['body-boolean.json', 'booleangroup'],
  'bytegroup':['body-byte.json','bytegroup'],
  'complexgroup':['body-complex.json','complexgroup'],
  'dategroup':['body-date.json','dategroup'],
  'datetimerfc1123group':['body-datetime-rfc1123.json','datetimerfc1123group'],
  'datetimegroup':['body-datetime.json','datetimegroup'],
  'dictionarygroup':['body-dictionary.json','dictionarygroup'],
  'durationgroup':['body-duration.json','durationgroup'],
  'filegroup':['body-file.json', 'filegroup'],
  'formdatagroup':['body-formdata.json', 'formdatagroup'],
  'integergroup':['body-integer.json','integergroup'],
  'numbergroup':['body-number.json','numbergroup'],
  'stringgroup':['body-string.json','stringgroup'],
  'custombaseurlgroup':['custom-baseUrl.json', 'custombaseurlgroup'],
  'headergroup':['header.json','headergroup'],
  'httpinfrastructuregroup':['httpInfrastructure.json','httpinfrastructuregroup'],
  'lrogroup':['lro.json', 'lrogroup'],
  'modelflatteninggroup':['model-flattening.json', 'modelflatteninggroup'],
  'report':['report.json','report'],
  'optionalgroup':['required-optional.json','optionalgroup'],
  'urlgroup':['url.json','urlgroup'],
  'urlmultigroup':['url-multi-collectionFormat.json','urlmultigroup'],
  'validationgroup':['validation.json', 'validationgroup'],
  'paginggroup':['paging.json', 'paginggroup'],
  'morecustombaseurigroup':['custom-baseUrl-more-options.json', 'morecustombaseurigroup'],
  'azurereport':['azure-report.json', 'azurereport']
}

swaggerDir = "node_modules/@microsoft.azure/autorest.testserver/swagger"

task 'regenerate-go', '', (done) ->
  regenExpected {
    'outputBaseDir': 'test/src/tests/generated',
    'inputBaseDir': swaggerDir,
    'mappings': goMappings,
    'packageNameBase': 'tests/generated'
  },done
  return null

task 'regenerate', "regenerate expected code for tests", ['regenerate-go'], (done) ->
  done();
