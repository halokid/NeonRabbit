import sys

sys.path.append("..")
from tasks.network_state import NetworkState
from tasks.lsopt import Lsopt
from tasks.sample1 import Sample1
from tasks.sample2 import Sample2

# hosts need to connect
# TODO: different ansible node take care of different HOST LIST, set in the run argument
HOST_LIST_A = ['root@192.168.1.139']
HOST_LIST_B = ['root@192.168.1.139']

# tasks need to run
# TASKS = {'NetworkState': NetworkState}
TASKS = {
  # 'NetworkState': NetworkState,
  # 'Lsopt': Lsopt,
  'Sample1': Sample1,
  'Sample2': Sample2,
}
