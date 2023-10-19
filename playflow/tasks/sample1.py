#!/usr/bin/env python

import sys
sys.path.append("..")
from tasks.tasks import Task

class Sample1(Task):

  @staticmethod
  def name():
    return 'Sample1'

  @staticmethod
  def run():
    print('Sample1.run() -->>>')
    # invoke broker service
    # client.Client.invoke("neon_broker", "pub", pub_data)
    return Task.check_run_status([], [], [], 'Sample1 run fail')
    # return Task.check_run_status([1], [1], [1], 'Sample1 run success')

