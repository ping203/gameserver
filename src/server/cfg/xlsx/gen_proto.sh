#!/bin/bash

# type config
./tabtoy \
--mode=exportorv2 \
--protover=3 \
--proto_out=../gameconf/game_type.proto  \
--combinename=DoNotUseThis \
--lan=zh_cn \
sgs_game_type_conf.xlsx

# base config
./tabtoy \
--mode=exportorv2 \
--protover=3 \
--protooutputignorefile=DoNotUseThis \
--protoimport=gameconf/game_type.proto \
--proto_out=../gameconf/game_base_config.proto \
--pbt_out=../gameconf/game_base_config.pbt \
--combinename=GameBaseConfig \
--lan=zh_cn \
sgs_game_type_conf.xlsx \
sgs_global_conf.xlsx \

