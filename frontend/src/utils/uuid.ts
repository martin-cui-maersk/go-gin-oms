const hexList: string[] = [];
for (let i = 0; i <= 15; i++) {
  hexList[i] = i.toString(16);
}

export function buildUUID(): string {
  let uuid = '';
  for (let i = 1; i <= 36; i++) {
    if (i === 9 || i === 14 || i === 19 || i === 24) {
      uuid += '-';
    } else if (i === 15) {
      uuid += 4;
    } else if (i === 20) {
      // uuid += hexList[(Math.random() * 4) | 8];
      uuid += hexList[getRandomIntInRange(8, 11)];
    } else {
      // uuid += hexList[(Math.random() * 16) | 0];
      uuid += hexList[getRandomIntInRange(0, 15)];
    }
  }
  return uuid.replace(/-/g, '');
}

let unique = 0;
export function buildShortUUID(prefix = ''): string {
  const time = Date.now();
  const random = Math.floor(Math.random() * 1000000000);
  unique++;
  return prefix + '_' + random + unique + String(time);
}

function getRandomIntInRange(min, max) {
  // 确保min和max是整数，并且min <= max
  if (typeof min !== 'number' || typeof max !== 'number' || min > max) {
    throw new Error('Invalid range');
  }

  // 计算范围大小
  const range = max - min + 1;

  // 生成一个Uint8Array，长度为1，因为我们只需要一个字节的随机数
  const array = new Uint8Array(1);
  window.crypto.getRandomValues(array);

  // 获取一个0到255之间的随机数
  const randomNumber = array[0];

  // 使用模运算来确保结果在0到range-1之间
  const index = randomNumber % range;

  // 将索引映射到min到max的范围内
  return min + index;
}
