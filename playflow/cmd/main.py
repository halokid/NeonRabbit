#!/usr/bin/env python

import sys
import time

sys.path.append("..")
# from ..pkgx.common import TASKS
from pkgx import common
from flow import flow

if __name__ == '__main__':
  # for k, task in common.TASKS.items():
  #   print('task run -->>>', k)
  #   task.run()

  flow = flow.Flow('mysample')
  # tasks = [NetworkState, Lsopt]
  tasks = common.TASKS
  flow.set_tasks(tasks)
  flow.run()
  # time.sleep(100)
