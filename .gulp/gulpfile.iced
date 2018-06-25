require './common.iced'

# ==============================================================================
# tasks required for this build 
Tasks "dotnet"  # dotnet functions
Tasks "regeneration"
Tasks "publishing"

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

# Update test dependencies
task 'updatetestdep', "", [], (done) ->
  process.env.GOPATH = path.normalize("#{basefolder}/test")
  await execute "dep ensure -update",          { cwd: './test/src/tests' }, defer code, stderr, stdout
  done()

task 'install_common',"", (done) ->
  # global.verbose = true
  execute "npm install",{cwd:"#{basefolder}/autorest.common", silent:false }, done


# Run language-specific tests:
task 'test', "", ['regenerate'], (done) ->
  process.env.GOPATH = path.normalize("#{basefolder}/test")
  await execute "dep ensure",          { cwd: './test/src/tests' }, defer code, stderr, stdout
  await execute "go run ./runner.go",     { cwd: './test/src/tests' }, defer code, stderr, stdout
  done()

# CI job
task 'testci', "more", [], (done) ->
  # install latest AutoRest
  await autorest ["--latest"], defer code, stderr, stdout

  ## TEST SUITE
  global.verbose = true
  await run "test", defer _

  ## REGRESSION TEST
  global.verbose = false
  # diff ('add' first so 'diff' includes untracked files)
  await  execute "git add -A", defer code, stderr, stdout
  await  execute "git diff --staged -w", defer code, stderr, stdout
  # eval
  echo stderr
  echo stdout
  throw "Potentially unnoticed regression (see diff above)! Run `npm run regenerate`, then review and commit the changes." if stdout.length + stderr.length > 0
  done() 