require './common.iced'

# ==============================================================================
# tasks required for this build 
Tasks "dotnet"  # dotnet functions
Tasks "regeneration"

# ==============================================================================
# Settings
Import
  initialized: false
  solution: "#{basefolder}/autorest.go.sln"
  sourceFolder:  "#{basefolder}/src/"

# ==============================================================================
# Tasks

task 'init', "" ,(done)->
  Fail "YOU MUST HAVE NODEJS VERSION GREATER THAN 7.10.0" if semver.lt( process.versions.node , "7.10.0" )
  done()
  
# Run language-specific tests:
task 'test', "", ['regenerate'], (done) ->
  process.env.GOPATH = "#{basefolder}/test"
  await execute "glide up",               { cwd: './test/src/tests' }, defer code, stderr, stdout
  await execute "go fmt ./generated/...", { cwd: './test/src/tests' }, defer code, stderr, stdout
  await execute "go run ./runner.go",     { cwd: './test/src/tests' }, defer code, stderr, stdout
  done();
