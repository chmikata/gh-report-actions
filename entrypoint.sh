#!/bin/bash
set -e

while getopts "c:o:r:t:T:i:l:s:" opt; do
  case "${opt}" in
    c)
      cmd=${OPTARG}
      args=${cmd}
    ;;
    o)
      args=$(echo "${args} -o ${OPTARG}")
    ;;
    r)
      args=$(echo "${args} -r ${OPTARG}")
    ;;
    t)
      args=$(echo "${args} -t ${OPTARG}")
    ;;
    T)
      args=$(echo "${args} -T ${OPTARG}")
    ;;
    i)
      args=$(echo "${args} -i ${OPTARG}")
    ;;
    l)
      args=$(echo "${args} -l ${OPTARG}")
    ;;
    s)
      args=$(echo "${args} -s ${OPTARG}")
    ;;
  esac
done

output=$(/app/gh-report-cli ${args})
echo "result=${output}" >> "${GITHUB_OUTPUT}"
cat "${GITHUB_OUTPUT}"
