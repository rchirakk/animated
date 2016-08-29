#!/usr/bin/python

import datetime
import json
from ansible.module_utils.basic import AnsibleModule

def main():
  module = AnsibleModule(
        argument_spec = dict(
            state     = dict(default='present', choices=['present', 'absent']),
            name      = dict(required=False),
            enabled   = dict(required=False, type='bool'),
            something = dict(aliases=['whatever'])
        )
  )

  date = str(datetime.datetime.now())
  module.exit_json(changed=True, date=date)

if __name__ == '__main__':
    main()
