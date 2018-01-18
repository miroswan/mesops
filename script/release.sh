#!/usr/bin/env bash

set -eo pipefail

declare TARGET_VERSION
declare WEAK_SEMANTIC_MATCHER="[[:digit:]]+\.[[:digit:]]+\.[[:digit:]]+"


show(){
  printf "==> $@\n"
}

error(){
  printf "ERROR: $@\n"
}

validate_args(){
  # if no arguments are passed to the script, BASH_ARGC is not 0, but nothing,
  # so default to 0
  if ((${BASH_ARGC:=0} != 1)); then
    error "You must pass the target release version as the first and only argument"
    return 1
  fi
}

validate_version(){
  # good-enough validation. We expect you to know what a version is. This is
  # simply to avoid typos
  echo $@ | egrep --quiet "$WEAK_SEMANTIC_MATCHER" || {
    error "$@ is not a validate version"
    return 1
  }
}

set_target_version(){
  TARGET_VERSION="${BASH_ARGV}"
}

update_files(){
  [ -z $TARGET_VERSION ] && {
    error "TARGET_VERSION must be set"
    return 1
  }
  if [[ "$OSTYPE" == *darwin* ]]; then
    regex_flags=(-E)
  else
    regex_flags=(--regexp-extended)
  fi
  local files="$@"
  for current_file in $files; do
    local tmp=$(mktemp)
    sed ${regex_flags[@]} "s/$WEAK_SEMANTIC_MATCHER/$TARGET_VERSION/g" "$current_file" > "$tmp" || {
      error "failed to update target version in $current_file"
      return 1
    }
    mv "$tmp" "$current_file" || {
      error "failed to move $tmp to $current_file"
      return 1
    }
  done
}


main(){
  validate_args || return $?
  set_target_version || return $?
  validate_version "$TARGET_VERSION"

  show "checking out to dev"
  git checkout dev || return $?

  show "sync remote to dev"
  git pull origin dev || return $?

  show "checking out to new release branch: release/$TARGET_VERSION"
  git checkout -b "release/$TARGET_VERSION" || return $?

  show "updating files pkg/mesops.go and README.md with $TARGET_VERSION"
  update_files pkg/mesops.go README.md || return $?

  show "committing version bump"
  git commit -am "release v$TARGET_VERSION" || return $?

  show "tagging at v$TARGET_VERSION"
  git tag "v$TARGET_VERSION" || return $?

  show "pushing to release/$TARGET_VERSION"
  git push --tags origin "release/$TARGET_VERSION"

  show "updating CHANGELOG.md"
  github_changelog_generator || return $?

  show "commiting CHANGELOG.md"
  git commit -am "updating CHANGELOG.md" || return $?

  show "pushing CHANGELOG.md updates to release/$TARGET_VERSION"
  git push origin "release/$TARGET_VERSION" || return $?

  show "checking out to dev"
  git checkout dev || return $?

  show "merging release/$TARGET_VERSION into dev"
  git merge "release/$TARGET_VERSION" || return $?

  show "pushing updates to dev"
  git push origin dev || return $?

  show "checkout to master"
  git checkout master || return $?

  show "merging release/$TARGET_VERSION into master"
  git merge "release/$TARGET_VERSION" || return $?

  show "pushing updates to master"
  git push origin master || return $?

  show "checking out back to dev"
  git checkout dev || return $?
}


################################################################################
# Execute

if [[ $BASH_SOURCE == $0 ]]; then
  main $@
fi
