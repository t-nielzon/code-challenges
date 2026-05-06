function siri(s) {
  let m;
  if (m = s.match(/Add (\d+) to (\d+)/)) {
    return Number(m[1]) + Number(m[2]);
  }
  if (m = s.match(/Subtract (\d+) from (\d+)/)) {
    return Number(m[2]) - Number(m[1]);
  }
  if (m = s.match(/weather at (\d+):(\d+)(am|pm)/i)) {
    let h = Number(m[1]);
    const min = Number(m[2]);
    const ap = m[3].toLowerCase();
    if (ap === 'am' && h === 12) h = 0;
    else if (ap === 'pm' && h !== 12) h += 12;
    const total = h * 60 + min;
    return (total >= 360 && total <= 1080) ? 'sunny' : 'raining';
  }
}