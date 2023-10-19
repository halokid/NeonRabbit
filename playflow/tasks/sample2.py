#!/usr/bin/env python

import sys
sys.path.append("..")
from tasks.tasks import Task

class Sample2(Task):

  @staticmethod
  def name():
    return 'Sample2'

  @staticmethod
  def run():
    print('Sample2.run() -->>>')
    # invoke broker service
    # client.Client.invoke("neon_broker", "pub", pub_data)
    return Task.check_run_status([1], [1], [1], 'Sample2 run result')

