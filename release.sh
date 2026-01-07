set -eo pipefail

CURRENT_YEAR=$(date +"%Y")
NEXT_YEAR=$((CURRENT_YEAR))
# NEXT_YEAR=$((CURRENT_YEAR+1))

_configurations=(
  2 "cfg/base.yaml,cfg/rm2.base.yaml,cfg/template_months_on_side.yaml,cfg/rm2.mos.default.yaml,cfg/rm2.mos.default.dailycal.yaml"           "rm2.mos.default.dailycal"
)

_configurations_len=${#_configurations[@]}

function createPDFs() {
  for _year in $CURRENT_YEAR $NEXT_YEAR; do
    for _idx in $(seq 0 3 $((_configurations_len-1))); do
      _passes=${_configurations[_idx]}
      _cfg=${_configurations[_idx+1]}
      _name=${_configurations[_idx+2]}

      PLANNER_YEAR="${_year}" PASSES="${_passes}" CFG="${_cfg}" NAME="${_name}.${_year}" ./single.sh
    done
  done
}

function mvDefaultTo() {
  for filename in ./*pdf; do
    _newname=$(echo "$filename" | perl -pe "s/default/$1/g")
    mv "$filename" "$_newname"
  done
}

function _restore() {
  git restore cfg/base.yaml
}

_combinations=(
  ""                        "dotted.default.ampm.sun"
)

_combinations_len=${#_combinations[@]}

for _idx in $(seq 0 2 $((_combinations_len-1))); do
  _cmds=${_combinations[_idx]}
  _mvTo=${_combinations[_idx+1]}

  for _cmd in ${_cmds}; do
    ${_cmd}
  done

  createPDFs
  mvDefaultTo "${_mvTo}"
  mv ./*pdf result

  _restore
done
