#!/bin/sh

kpath=$HOME/.kube

if [[ -f $kpath/india ]]; then
    echo "Found india config file"
else
    echo "india config file not found"
    exit 1
fi;

if [[ -f $kpath/westus ]]; then
    echo "Found westus config file"
else
    echo "westus config file not found"
    exit 1
fi;

if [[ -f $kpath/seasia ]]; then
    echo "Found seasia config file"
else
    echo "seasia config file not found"
    exit 1
fi;

if [[ -f $kpath/ukdo ]]; then
    echo "Found ukdo config file"
else
    echo "ukdo config file not found"
    exit 1
fi;

echo -e "\ndeploying to uk do\n"
mv $kpath/ukdo $kpath/config
kubectl rollout restart -f ./scheduler-config-do.yaml
mv $kpath/config $kpath/ukdo


echo -e "\ndeploying to india azure\n"
mv $kpath/india $kpath/config
kubectl rollout restart -f ./scheduler-config-azure.yaml
mv $kpath/config $kpath/india


echo -e "\ndeploying to southeast asia azure\n"
mv $kpath/seasia $kpath/config
kubectl rollout restart -f ./scheduler-config-azure.yaml
mv $kpath/config $kpath/seasia

echo -e "\ndeploying to westus azure\n"
mv $kpath/westus $kpath/config
kubectl rollout restart -f ./scheduler-config-azure.yaml
mv $kpath/config $kpath/westus