import json

from ansible.plugins.callback import CallbackBase


# Create a callback plugin so we can capture the output
class ResultsCollectorJSONCallback(CallbackBase):
  """A sample callback plugin used for performing an action as results come in.

If you want to collect all results into a single object for processing at
the end of the execution, look into utilizing the ``json`` callback plugin
or writing your own custom callback plugin.
"""

  def __init__(self, *args, **kwargs):
    super(ResultsCollectorJSONCallback, self).__init__(*args, **kwargs)
    self.host_ok = {}
    self.host_unreachable = {}
    self.host_failed = {}

  def v2_runner_on_unreachable(self, result):
    host = result._host
    self.host_unreachable[host.get_name()] = result

  def v2_runner_on_ok(self, result, *args, **kwargs):
    """Print a json representation of the result.

Also, store the result in an instance attribute for retrieval later
"""
    host = result._host
    self.host_ok[host.get_name()] = result
    print(json.dumps({host.name: result._result}, indent=4))

  def v2_runner_on_failed(self, result, *args, **kwargs):
    host = result._host
    self.host_failed[host.get_name()] = result


# '''
class Task(object):
  name = ''

  def __init__(self, name):
    self.name = name

  @staticmethod
  def run(self):
    return False

  @staticmethod
  def check_run_status(host_ok, host_failed, host_unreachable, task_report = 'No Report'):
    if host_ok:
      return True, task_report
    if host_failed or host_unreachable:
      return False, task_report
    return False, task_report
# '''

