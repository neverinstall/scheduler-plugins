#!/bin/sh

all_or_one=$1

if [[ $all_or_one = 'ukdo' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to uk do"
    python3 $HOME/Python/utils/setKubeContext.py ukdo
    kubectl rollout restart -f ./scheduler-config-do.yaml
fi

if [[ $all_or_one = 'indiaaz' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to india azure"
    python3 $HOME/Python/utils/setKubeContext.py india
    kubectl rollout restart -f ./scheduler-config-azure.yaml
fi

if [[ $all_or_one = 'seasiaaz' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to southeast asia azure"
    python3 $HOME/Python/utils/setKubeContext.py seasia
    kubectl rollout restart -f ./scheduler-config-azure.yaml
fi

if [[ $all_or_one = 'usaz' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to westus azure"
    python3 $HOME/Python/utils/setKubeContext.py westus
    kubectl rollout restart -f ./scheduler-config-azure.yaml
fi

if [[ $all_or_one = 'indiado' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to do india"
    python3 $HOME/Python/utils/setKubeContext.py indiado
    kubectl rollout restart -f ./scheduler-config-azure.yaml # the core is named like it is in azure
fi

if [[ $all_or_one = 'sfdo' || $all_or_one = 'all' ]]; then
    printf "\ndeploying to do san fransisco"
    python3 $HOME/Python/utils/setKubeContext.py sfdo
    kubectl rollout restart -f ./scheduler-config-azure.yaml # the core is named like it is in azure
fi