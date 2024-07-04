# questionbankgenerator
Add new config files in the configs folder for each file to generate

Tested this to generate 50 questions for each topic without failing...

I exposed a file in the configs, jsonouput, output folders to understand how things work.

The configs is your input for all subjects to be generated.
jsonoutput is what the goprogram return when you run ./run.sh

while the ouput is the final migration folder to generate docs file for all file in jsonoutput by running ./migratetodocs.sh

make sure to chmod +x for both ./run.sh and ./migratetodocs.sh to avoid permison errors