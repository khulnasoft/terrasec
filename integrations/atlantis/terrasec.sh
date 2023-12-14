#!/bin/bash

terrasec scan -i tfplan --iac-version v1 -f ${PLANFILE}.json -l error > output
exitcode=$?

if [[ ! $exitcode -eq 0 ]]; then
    echo
    echo '- Terrasec identified IAC policy violations:'
    echo
    echo 'Scan Results:'
    cat output
    echo
    echo '```'
    echo '</details>'
    echo '<p><strong>Further atlantis details below:</strong></p>'
    echo '<details>'
    echo
    echo '```diff'
    echo
fi

exit $exitcode
