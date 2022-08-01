# https://qiita.com/koshigoe/items/054383a89bd51d099f10

import asyncio
import sys
import itertools

async def heavy():
    await asyncio.sleep(10)
    return 'done.'


async def spin():
    write, flush = sys.stdout.write, sys.stdout.flush
    for c in itertools.cycle('|/-\\'):
        write(c)
        flush()
        write('\x08')
        try:
            await asyncio.sleep(0.1)
        except asyncio.CancelledError:
            break
    write(' \x08')


async def task():
    done, pending = await asyncio.wait([spin(), heavy()], return_when=asyncio.FIRST_COMPLETED)
    for task in pending:
        task.cancel()

    return done.pop().result()


def main():
    loop = asyncio.get_event_loop()
    result = loop.run_until_complete(task())
    loop.close()
    print(f'Result: {result}')


if __name__ == "__main__":
    main()
