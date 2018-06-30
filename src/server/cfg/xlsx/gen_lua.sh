#!/bin/bash

# global config
./tabtoy \
--mode=exportorv2 \
--protover=3 \
--lua_out=../config/global.lua \
--combinename=GlobalConfig \
--lan=zh_cn \
sgs_game_type_conf.xlsx \
sgs_global_conf.xlsx
