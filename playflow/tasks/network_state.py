#!/usr/bin/env python

import sys

sys.path.append("..")
from pkgx import common
from tasks.tasks import ResultsCollectorJSONCallback
from service import client
from tasks.tasks import Task

from ansible.module_utils.common.collections import ImmutableDict
from ansible.inventory.manager import InventoryManager
from ansible.vars.manager import VariableManager
from ansible.parsing.dataloader import DataLoader
from ansible.executor.task_queue_manager import TaskQueueManager
from ansible.playbook.play import Play
from ansible import context
import ansible.constants as C
import shutil


class NetworkState(Task):

  @staticmethod
  def name():
    return 'NetworkState'

  @staticmethod
  def run():
    print('NetworkState.run() -->>>')
    host_list = common.HOST_LIST_A

    context.CLIARGS = ImmutableDict(
      connection='smart', module_path=[
        '/to/mymodules', '/usr/share/ansible'], forks=10, become=None,
      become_method=None, become_user=None, check=False, diff=False,
      verbosity=0
    )

    sources = ','.join(host_list)
    if len(host_list) == 1:
      sources += ','

    loader = DataLoader()
    passwords = dict(vault_pass='secret')

    results_callback = ResultsCollectorJSONCallback()
    inventory = InventoryManager(loader=loader, sources=sources)
    variable_manager = VariableManager(loader=loader, inventory=inventory)

    tqm = TaskQueueManager(
      inventory=inventory,
      variable_manager=variable_manager,
      loader=loader,
      passwords=passwords,
      stdout_callback=results_callback,
    )

    play_source = dict(
      name="Ansible Play",
      hosts=host_list,
      gather_facts='no',
      tasks=[
        dict(action=dict(module='shell', args='ifconfig'), register='shell_out'),
      ]
    )
    print("brakpoint 1 -->>>")
    play = Play().load(play_source, variable_manager=variable_manager, loader=loader)
    try:
      print("brakpoint 2 -->>>")
      result = tqm.run(play)
    finally:
      print("brakpoint 5 -->>>")
      tqm.cleanup()
      if loader:
        loader.cleanup_all_tmp_files()

    print("brakpoint 3 -->>>")
    shutil.rmtree(C.DEFAULT_LOCAL_TMP, True)
    print("brakpoint 4 -->>>")

    pub_data = ''
    print("UP ***********")
    for host, result in results_callback.host_ok.items():
      # print('{0} >>>\n {1}'.format(host, result._result['stdout']))
      pub_data += result._result['stdout']

    print("FAILED *******")
    for host, result in results_callback.host_failed.items():
      print('{0} >>>\n {1}'.format(host, result._result['msg']))
      pub_data += result._result['msg']

    print("DOWN *********")
    for host, result in results_callback.host_unreachable.items():
      print('{0} >>>\n {1}'.format(host, result._result['msg']))
      pub_data += result._result['msg']

    # invoke broker service
    # client.Client.invoke("neon_broker", "pub", pub_data)

    return Task.check_run_status(results_callback.host_ok, results_callback.host_failed,
                                 results_callback.host_unreachable, pub_data)

