#!/usr/bin/env python

import sys
import re
from typing import Optional

key_emoji = {
    'NEW': '📦',
    'IMPROVE': '👌',
    'FIX': '🐛',
    'DOCS': '📖',
    'RELEASE': '🚀',
    'MERGE': '🧲'
}

def transformHeader(header: str) -> Optional[str]:
    # Handle the automatically generated merge commit message
    merge_message_prefix = 'Merge branch'
    if header.startswith(merge_message_prefix):
        header = 'MERGE:' + header.removeprefix(merge_message_prefix)

    re_keys = '|'.join([key for key in key_emoji.keys()])

    m = re.match(rf'(?P<key>{re_keys}):\s*(?P<value>.+)', header, re.I)
    if m is None:
        return None

    key = sanitizeKey(m.group('key'))
    value = sanitizeValue(m.group('value'))
    return f'{key} {key_emoji[key]}: {value}\n'

def sanitizeKey(key: str) -> str:
    return key.upper()

def sanitizeValue(value: str) -> str:
    value = value.strip()
    # pop last char if it's a dot
    if value[-1] == '.':
        value = value[:-1]
    words = value.split(' ')
    words[0] = words[0].capitalize();
    return ' '.join(words);

def verifyMessage(path: str) -> bool:
    with open(path, 'r+') as file:
        lines = file.readlines();
        header = transformHeader(lines[0])
        if header is None:
            return False

        lines[0] = header
        file.seek(0)
        file.writelines(lines)

    return True


if __name__ == '__main__':
    try:
        assert len(sys.argv) == 2, 'This file should only be called as a commit-msg hook in Git'

        path = sys.argv[1]

        assert verifyMessage(path), 'Invalid commit message format. Check the docs if you are unsure.'
    except Exception as e:
        print(e)
        sys.exit(1)
        


    


