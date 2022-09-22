# dtm show

`dtm show` 用于展示插件的配置文件模板或状态。

## 1 展示 插件/工具链 配置模板

`dtm show config` 展示插件的配置文件模板或示例工具链的配置文件模板。

命令行参数：

| 短  | 长         | 默认值 | 描述                                            |
|-----|------------|-------|------------------------------------------------|
| -p  | --plugin   | `""`  | 插件名字                                        |
| -t  | --template | `""`  | 示例工具链名字，目前支持 "quickstart" 和 "gitopts" |

## 2 展示插件状态

`dtm show status` 展示插件实例的状态。

命令行参数：

| 短  | 长            | 默认值                    | 描述                     |
|-----|---------------|--------------------------|-------------------------|
| -p  | --plugin      | `""`                     | 插件名字                 |
| -i  | --id          | `""`                     | 插件实例 id              |
| -a  | --all         | `false`                  | 展示所有插件的所有实例的状态 |
| -d  | --plugin-dir  | `"~/.devstream/plugins"` | 插件目录                 |
| -f  | --config-file | `"config.yaml"`          | 配置文件                 |

_说明：如果 `-p` 和 `-i` 都为空，`dtm` 将展示所有插件的所有实例的状态。_