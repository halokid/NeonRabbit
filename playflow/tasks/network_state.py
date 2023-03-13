#!/usr/bin/env python

from ..pkgx import common
from .tasks import ResultsCollectorJSONCallback
from .tasks import Task
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

    def run(self):
        host_list = common.host_list

        context.CLIARGS = ImmutableDict(
            connection='smart', module_path=[
                '/to/mymodules', '/usr/share/ansible'], forks=10, become=None,
            become_method=None, become_user=None, check=False, diff=False,
            verbosity=0
        )

        sources = ','.join(host_list)
        if len(host_list) == 1:
            sources += ','

        loader = DataLoader()  # Takes care of finding and reading yaml, json and ini files
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
                dict(action=dict(module='shell', args='ls'), register='shell_out'),
                # dict(action=dict(module='debug', args=dict(msg='{{shell_out.stdout}}'))),
                # dict(action=dict(module='command', args=dict(cmd='/usr/bin/uptime'))),
            ]
        )

        play = Play().load(play_source, variable_manager=variable_manager, loader=loader)

        # Actually run it
        try:
            result = tqm.run(play)  # most interesting data for a play is actually sent to the callback's methods
        finally:
            # we always need to cleanup child procs and the structures we use to communicate with them
            tqm.cleanup()
            if loader:
                loader.cleanup_all_tmp_files()

        # Remove ansible tmpdir
        shutil.rmtree(C.DEFAULT_LOCAL_TMP, True)

        print("UP ***********")
        for host, result in results_callback.host_ok.items():
            print('{0} >>> {1}'.format(host, result._result['stdout']))

        print("FAILED *******")
        for host, result in results_callback.host_failed.items():
            print('{0} >>> {1}'.format(host, result._result['msg']))

        print("DOWN *********")
        for host, result in results_callback.host_unreachable.items():
            print('{0} >>> {1}'.format(host, result._result['msg']))
