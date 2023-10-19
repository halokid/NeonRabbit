from dapr.clients import DaprClient
# import logging

class Client(object):

  def __init__(self):
    pass

  @staticmethod
  def invoke(service, method, data):
    with DaprClient() as d:
      resp = d.invoke_method(service, method, data)
    # logging.basicConfig(level='INFO')
    # logging.info('Client invoke -->>>')
    # logging.info(resp.data)
    return resp.data
