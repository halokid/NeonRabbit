#!/usr/bin/env python

import sys

sys.path.append("..")
from flow.flow import Flow
# from tasks.network_state import NetworkState
# from tasks.lsopt import Lsopt

if __name__ == '__main':
  flow = Flow()
  # tasks = [NetworkState, Lsopt]
  tasks = []
  flow.set_tasks(tasks)
  flow.run()

