Karver: Filesystem migrations
=============================

What is Karver?
---------------

Inspired on ActiveRecord, Karver allows you to follow the same migrations pattern but against the filesystem.

How it works?
-------------

Karver needs a list of migrations and a target path. Each migration will receive the target as its first parameter.

For each migration, if the execution ends successfully, a mark (.karver) will be written on the target path. That mark will be used in the future in order to determine which migrations need to be run, and which ones have already been run.

If the mark doesn't exist, all migrations will be executed.

What is it trying to solve?
---------------------------

Karver tries to solve exactly this situation:
You have 3 instances of a filesystem, each one in a different version (imagine a VM with 3 different versions of your software):
i.e.:
* v1 creates /foo
* v2 moves /foo to /bar/foo
* v3 creates a symbolic link: /bar/foo -> /new/path/foo

How do you create a single package able to upgrade from v1 and from v2, to v3? No manual intervention is allowed.

Configuration management is not good enough. Introducing the needed logic for this situation can create a real spaghetti monster.
Instead, Karver allows you to create an executable for each step (a simple bash script?), and run them only when needed based on the starting point.
If executed against v1, it would run migrations for v1 -> v2 and v2 -> v3. If executed against v2, only v2 -> v3 migration.

This way, you don't need to worry about the version of your target filesystem.

Example
-------
```
$ mkdir -p /opt/karver/target/foo

$ mkdir -p /opt/karver/migrations

$ karver --migrations /opt/karver/migrations create 'first migration'
New migration: /opt/karver/migrations/20140604205700_first_migration.sh

$ cat > /opt/karver/migrations/20140604205700_first_migration.sh << EOF
#!/bin/bash
TARGET=\$1
touch \$TARGET/new-file-from-migration
if [ -d \$TARGET/foo ]; then mv \$TARGET/foo \$TARGET/bar; fi
EOF

$ karver --migrations /opt/karver/migrations/ --target /opt/karver/target/ run
2014/06/04 21:06:45 Karving /opt/karver/target/...
2014/06/04 21:06:45 Running migration 20140604205700_first_migration.sh...
2014/06/04 21:06:45 /opt/karver/target/ has been karved. :D

$ ls -a /opt/karver/target/
.karver                 bar                     new-file-from-migration
```

Have you met docker-karver?
---------------------------
You have a data container and want to upgrade it without manual intervention?
https://github.com/karver/docker-karver
