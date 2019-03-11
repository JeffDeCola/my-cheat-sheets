REM my-cheat-sheets set-pipeline.cmd
@ECHO OFF
ECHO "START set-pipeline.cmd script - From Windows"
START "concourse" "C:\Program Files (x86)\Concourse\fly-4.0.0.exe" -t ci set-pipeline -p my-cheat-sheets -c pipeline.yml --load-vars-from .credentials.yml
ECHO "END set-pipeline.cmd script"