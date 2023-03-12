#!/usr/bin/env python

from ..pkgx import common


def run():
  host_list = common.host_list

  context.CLIARGS = ImmutableDict(
    connection='smart', module_path=[
      '/to/mymodules', '/usr/share/ansible'], forks=10, become=None,
    become_method=None, become_user=None, check=False, diff=False,
    verbosity=0
  )
