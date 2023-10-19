import sys
sys.path.append("..")
from service import client
import json

def gen_svc_req(data):
  code = 0
  message = 'success'
  for task_data in data:
    if task_data['state'] != 'success':
      code = 1
      message = 'fail'
      break
  return {
    'code': code,
    'message': message,
    'data': data
  }

class Flow(object):

  def __init__(self, name):
    self.tasks = {}
    # self.all_report = ''
    self.data = []
    self.name = name

  def get_name(self):
    return self.name

  def set_tasks(self, tasks):
    self.tasks = tasks

  def run(self):
    print('-->>> run flow {}, includ tasks {}'.format(self.name, self.tasks))
    if not len(self.tasks):
      print('-->>> No tasks in flow {}'.format(self.name))
    tmp_data = []
    tmp_data_tasks = []
    for task_name, task in self.tasks.items():
      print('----- run Task {} -----'.format(task_name))
      tmp_data_item = {'task_name': task_name}
      task_state, task_report = task.run()
      # self.all_report += task_report
      tmp_data_item['result'] = task_report
      tmp_data_tasks.append(task_name)
      if not task_state:
        print("-->>> Task {} fail!".format(task_name))
        tmp_data_item['state'] = 'fail'
        tmp_data.append(tmp_data_item)
        break
      else:
        print("-->>> Task {} success!".format(task_name))
        tmp_data_item['state'] = 'success'
        tmp_data.append(tmp_data_item)

    print('tmp_data_tasks -->>>', tmp_data_tasks)
    print('tmp_data -->>>', tmp_data)
    # TODO: compare Flow data and tmp_data
    report = self.report(tmp_data_tasks, tmp_data)
    print("Flow report -->>>", report)
    # invoke broker service
    # svc_req = gen_svc_req(data = report)    # send a dict
    svc_req = json.dumps(gen_svc_req(data = report))    # prefer send a string
    print("svc_req -->>>", svc_req)
    resp = client.Client.invoke("neon_broker", "pub", svc_req)
    print("Invoke broker rsp -->>>", resp)

  def report(self, tmp_data_tasks, tmp_data):
    if len(self.tasks) == len(tmp_data_tasks):
      return tmp_data
    task_names = list(self.tasks.keys())
    # tasks_names has but tmp_data_tasks not
    difference_set_tasks = list(set(task_names).difference(set(tmp_data_tasks)))
    for item in difference_set_tasks:
      tmp_data_item = dict()
      tmp_data_item['task_name'] = item
      tmp_data_item['state'] = 'noStart'
      tmp_data_item['result'] = ''
      tmp_data.append(tmp_data_item)
    return tmp_data

'''
if __name__ == '__main':
  flow = Flow(object)
  tasks = [Sample1, Sample2]
  # tasks = []
  flow.set_tasks(tasks)
  flow_report = flow.run()
  print('flow report -->>> ', flow_report)
'''


