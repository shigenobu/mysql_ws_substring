#!/bin/sh
name="ws_substring"
lib_name="${name}.so"

go build -buildmode=c-shared -o ${lib_name} ${name}.go

mysql_plugin_dir=`mysql_config --plugindir`
sudo mv ${lib_name} ${mysql_plugin_dir}/
sudo chown root:root ${mysql_plugin_dir}/${lib_name}
sudo chmod 755 ${mysql_plugin_dir}/${lib_name}

sudo mysql -e "drop function if exists ${name}"
sudo mysql -e "create function ${name} returns string soname '"${lib_name}"'"
sudo mysql -e "select * from mysql.func"
