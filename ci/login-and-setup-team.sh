#!/bin/bash
# my-cheat-sheets login-concourse-setup-team.sh

echo " "
echo "STEP 1 - LOGIN FIRST TIME TO default TEAM main"
fly --target main-ci-target \
    login \
    --team-name main \
    --concourse-url http://192.168.20.112:8080/
echo ""

echo "STEP 2 - CREATE NEW TEAM under USER test"
echo "fly --target main-ci-target set-team --team-name jeffs-ci-team --local-user test"
echo " "

echo "STEP 3 - LOGIN TO your new TEAM jeffs-ci-team"
echo "fly --target jeffs-ci-target login --team-name jeffs-ci-team --concourse-url http://192.168.20.112:8080/"
echo ""

echo " "
echo "TO CHECK YOUR TARGET"
echo "fly targets"
echo "cat ~/.flyrc"
echo " "
echo "TO SEE TEAMS"
echo "fly --target main-ci-target teams"
echo "fly --target jeffs-ci-target teams"
echo " "
echo "TO SEE USERS"
echo "fly --target main-ci-target userinfo"
echo "fly --target jeffs-ci-target userinfo"
echo " "
